package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/stats"
	"go.uber.org/atomic"

	"github.com/dennis-tra/nebula-crawler/pkg/config"
	"github.com/dennis-tra/nebula-crawler/pkg/db"
	"github.com/dennis-tra/nebula-crawler/pkg/metrics"
	"github.com/dennis-tra/nebula-crawler/pkg/models"
	"github.com/dennis-tra/nebula-crawler/pkg/queue"
	"github.com/dennis-tra/nebula-crawler/pkg/utils"
)

var dialerID = atomic.NewInt32(0)

// Dialer encapsulates a libp2p host that dials peers.
type Dialer struct {
	id          string
	host        host.Host
	config      *config.Config
	dialedPeers int
	done        chan struct{}
}

// NewDialer initializes a new dialer based on the given configuration.
func NewDialer(h host.Host, conf *config.Config) (*Dialer, error) {
	c := &Dialer{
		id:     fmt.Sprintf("dialer-%02d", dialerID.Load()),
		host:   h,
		config: conf,
		done:   make(chan struct{}),
	}
	dialerID.Inc()

	return c, nil
}

// StartDialing enters an endless loop and consumes dial jobs from the dial queue
// and publishes its result on the results queue until it is told to stop or the
// dial queue was closed.
func (d *Dialer) StartDialing(ctx context.Context, dialQueue *queue.FIFO, resultsQueue *queue.FIFO) {
	defer close(d.done)
	for {
		// Give the shutdown signal precedence
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case <-ctx.Done():
			return
		case elem, ok := <-dialQueue.Consume():
			if !ok {
				// The crawl queue was closed
				return
			}
			result := d.handleDialJob(ctx, elem.(peer.AddrInfo))
			resultsQueue.Push(result)
		}
	}
}

// handleCrawlJob takes a crawl result, aggregates crawl information and publishes the result
// to the persist queue, so that the persisters can persist the information in the database.
// It also looks into the result and publishes new crawl jobs based on whether the found peers
// weren't crawled before or are not already in the queue.
func (d *Dialer) handleDialJob(ctx context.Context, pi peer.AddrInfo) Result {
	// Creating log entry
	logEntry := log.WithFields(log.Fields{
		"dialerID":  d.id,
		"remoteID":  utils.FmtPeerID(pi.ID),
		"dialCount": d.dialedPeers,
	})
	logEntry.Debugln("Dialing peer")
	defer logEntry.Debugln("Dialed peer")

	// Initialize dial result
	dr := Result{
		DialerID:      d.id,
		Peer:          pi,
		DialStartTime: time.Now(),
	}

	// Try to dial the peer 3 times
retryLoop:
	for retry := 0; retry < 3; retry++ {

		// Update log entry
		logEntry = logEntry.WithField("retry", retry)

		// Add peer information to peer store so that DialPeer can pick it up from there
		// Do this in every retry due to the TTL of one minute
		d.host.Peerstore().AddAddrs(pi.ID, pi.Addrs, time.Minute)

		// Actually dial the peer
		if err := d.dial(ctx, pi.ID); err != nil {
			dr.Error = err
			dr.DialError = db.DialError(dr.Error)

			if errors.Is(err, context.Canceled) {
				break retryLoop
			}

			sleepDur := time.Duration(float64(retry+1) * float64(10*time.Second))
			errMsg := fmt.Sprintf("Dial failed, sleeping %s", sleepDur)

			switch dr.DialError {
			case models.DialErrorPeerIDMismatch:
				logEntry.WithError(err).Debugln("Dial failed due to peer ID mismatch - stopping retry")
				// TODO: properly connect to new peer and see if it is part of the DHT.
				break retryLoop
			case models.DialErrorNoPublicIP, models.DialErrorNoGoodAddresses:
				logEntry.WithError(err).Debugln("Dial failed due to no public ip - stopping retry")
				break retryLoop
			case models.DialErrorMaxDialAttemptsExceeded:
				sleepDur = 70 * time.Second
				errMsg = fmt.Sprintf("Max dial attempts exceeded, sleeping longer %s", sleepDur)
			case models.DialErrorConnectionRefused:
				// The monitoring task receives a lot of "connection refused" messages. I guess there is
				// a limit somewhere of how often a peer can connect. I could imagine that this rate limiting
				// is set to one minute. As the scheduler fetches all sessions that are due in the next 10
				// seconds I'll add that and another one just to be sure ¯\_(ツ)_/¯
				if retry >= 1 {
					logEntry.WithError(err).Debugf("Received 'connection refused' the second time - stopping retry")
					break retryLoop
				}
				sleepDur = 70 * time.Second
				errMsg = fmt.Sprintf("Connection refused, sleeping longer %s", sleepDur)
			default:
			}
			logEntry.WithError(err).Debugf(errMsg)
			select {
			case <-time.After(sleepDur):
			case <-ctx.Done():
				break retryLoop
			}
			continue retryLoop
		}

		// Dial was successful - reset error
		dr.Error = nil
		dr.DialError = ""

		break retryLoop
	}

	// Close established connection to prevent running out of FDs?
	if err := d.host.Network().ClosePeer(pi.ID); err != nil {
		logEntry.WithError(err).Warnln("Could not close connection to peer")
	}

	dr.DialEndTime = time.Now()
	return dr
}

func (d *Dialer) dial(ctx context.Context, peerID peer.ID) error {
	stats.Record(ctx, metrics.MonitorDialCount.M(1))

	if _, err := d.host.Network().DialPeer(ctx, peerID); err != nil {
		stats.Record(ctx, metrics.MonitorDialErrorsCount.M(1))
		return err
	}

	return nil
}
