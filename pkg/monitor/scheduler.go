package monitor

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.uber.org/atomic"

	"github.com/dennis-tra/nebula-crawler/pkg/config"
	"github.com/dennis-tra/nebula-crawler/pkg/db"
	"github.com/dennis-tra/nebula-crawler/pkg/metrics"
	"github.com/dennis-tra/nebula-crawler/pkg/models"
	"github.com/dennis-tra/nebula-crawler/pkg/queue"
	"github.com/dennis-tra/nebula-crawler/pkg/utils"
)

// The Scheduler handles the scheduling and managing of
//   a) dialers - They consume a queue of peer address information, visit them and publish their results
//                on a separate results queue. This results queue is consumed by this scheduler and further
//                processed
type Scheduler struct {
	// The libp2p node that's used to crawl the network. This one is also passed to all dialers.
	host host.Host

	// The database handle
	dbc *db.Client

	// The configuration of timeouts etc.
	config *config.Config

	// The queue of peer.AddrInfo's that need to be dialed to.
	dialQueue *queue.FIFO

	// A map from peer.ID to peer.AddrInfo to indicate if a peer was put in the queue, so
	// we don't put it there again.
	inDialQueue sync.Map

	// The number of peers in the ping queue.
	inDialQueueCount atomic.Uint32

	// The queue that the dialers publish their dial results on
	resultsQueue *queue.FIFO
}

// NewScheduler initializes a new libp2p host and scheduler instance.
func NewScheduler(ctx context.Context, conf *config.Config, dbc *db.Client) (*Scheduler, error) {
	// Set the timeout for dialing peers
	ctx = network.WithDialPeerTimeout(ctx, conf.DialTimeout)

	// Force direct dials will prevent swarm to run into dial backoff errors. It also prevents proxied connections.
	ctx = network.WithForceDirectDial(ctx, "prevent backoff")

	// Initialize a single libp2p node that's shared between all dialers.
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		return nil, err
	}

	h, err := libp2p.New(libp2p.Identity(priv), libp2p.NoListenAddrs, libp2p.UserAgent("nebula-crawler/"+conf.Version))
	if err != nil {
		return nil, err
	}

	s := &Scheduler{
		host:         h,
		dbc:          dbc,
		config:       conf,
		inDialQueue:  sync.Map{},
		dialQueue:    queue.NewFIFO(),
		resultsQueue: queue.NewFIFO(),
	}

	return s, nil
}

// StartMonitoring starts the configured amount of dialers and fills
// the dial queue with peers that are due to be dialed.
func (s *Scheduler) StartMonitoring(ctx context.Context) error {
	start := time.Now()

	// Start all dialers
	var dialers []*Dialer
	for i := 0; i < s.config.MonitorWorkerCount; i++ {
		d, err := NewDialer(s.host, s.config)
		if err != nil {
			return errors.Wrap(err, "new dialer")
		}

		dialers = append(dialers, d)
		go d.StartDialing(ctx, s.dialQueue, s.resultsQueue)
	}

	// Async handle the results from dialers
	go s.readResultsQueue(ctx)

	// Monitor the database and schedule dial jobs
	s.monitorDatabase(ctx)

	for _, d := range dialers {
		log.WithField("dialerId", d.id).Debugln("Waiting for dialer to stop")
		<-d.done
	}

	log.WithFields(log.Fields{
		"inDialQueue":     s.inDialQueueCount.Load(),
		"monitorDuration": time.Since(start),
	}).Infoln("Finished monitoring")

	return nil
}

// readResultsQueue listens for dial results on the resultsQueue and handles any
// entries in handleResult. If the scheduler is shut down it schedules a cleanup of resources.
func (s *Scheduler) readResultsQueue(ctx context.Context) {
	for {
		// Give the shutdown signal precedence
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case elem := <-s.resultsQueue.Consume():
			s.handleResult(ctx, elem.(Result))
		case <-ctx.Done():
			return
		}
	}
}

// handleResult takes the result of dialing a peer, logs general information and inserts this visit into the database.
func (s *Scheduler) handleResult(ctx context.Context, dr Result) {
	logEntry := log.WithFields(log.Fields{
		"dialerID": dr.DialerID,
		"remoteID": utils.FmtPeerID(dr.Peer.ID),
		"alive":    dr.Error == nil,
	})
	if dr.Error != nil {
		if dr.DialError == models.DialErrorUnknown {
			logEntry = logEntry.WithError(dr.Error)
		} else {
			logEntry = logEntry.WithField("error", dr.DialError)
		}
	}
	start := time.Now()
	if err := s.insertRawVisit(ctx, dr); err != nil {
		logEntry.WithError(err).Warnln("Could not persist dial result")
	}

	// Update maps
	s.inDialQueue.Delete(dr.Peer.ID)
	stats.Record(ctx, metrics.PeersToDialCount.M(float64(s.inDialQueueCount.Dec())))

	// Track dial errors for prometheus
	if dr.Error != nil {
		if ctx, err := tag.New(ctx, tag.Upsert(metrics.KeyError, dr.DialError)); err == nil {
			stats.Record(ctx, metrics.PeersToDialErrorsCount.M(1))
		}
	}

	logEntry.
		WithField("dialDur", dr.DialDuration()).
		WithField("persistDur", time.Since(start)).
		Infoln("Handled dial result from dialer", dr.DialerID)
}

// monitorDatabase checks every 10 seconds if there are peer sessions that are due to be renewed.
func (s *Scheduler) monitorDatabase(ctx context.Context) {
	for {
		log.Infof("Looking for sessions to check...")
		sessions, err := s.dbc.FetchDueSessions(ctx)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.WithError(err).Warnln("Could not fetch sessions")
			goto TICK
		}

		// For every session schedule that it gets pushed into the dialQueue
		for _, session := range sessions {
			if err = s.scheduleDial(ctx, session); err != nil {
				log.WithError(err).Warnln("Could not schedule dial")
			}
		}
		log.Infof("In dial queue %d peers", s.inDialQueueCount.Load())

	TICK:
		select {
		case <-time.Tick(10 * time.Second):
			continue
		case <-ctx.Done():
			return
		}
	}
}

// scheduleDial takes a session entity from the database constructs a peer.AddrInfo struct and feeds
// it into the queue of peers-to-dial to be picked up by one of the dialers.
func (s *Scheduler) scheduleDial(ctx context.Context, session *models.Session) error {
	// Parse peer ID from database
	peerID, err := peer.Decode(session.R.Peer.MultiHash)
	if err != nil {
		return errors.Wrap(err, "decode peer ID")
	}
	logEntry := log.WithField("peerID", utils.FmtPeerID(peerID))

	// Parse multi addresses from database
	pi := peer.AddrInfo{ID: peerID}
	for _, maddrStr := range session.R.Peer.R.MultiAddresses {
		maddr, err := ma.NewMultiaddr(maddrStr.Maddr)
		if err != nil {
			logEntry.WithError(err).Warnln("Could not parse multi address")
			continue
		}
		pi.Addrs = append(pi.Addrs, maddr)
	}

	// Check if peer is already in dial queue
	if _, inPingQueue := s.inDialQueue.LoadOrStore(peerID, pi); inPingQueue {
		return nil
	}
	stats.Record(ctx, metrics.PeersToDialCount.M(float64(s.inDialQueueCount.Inc())))

	// Schedule dial for peer
	s.dialQueue.Push(pi)

	return nil
}

// insertRawVisit builds up a raw_visit database entry.
func (s *Scheduler) insertRawVisit(ctx context.Context, cr Result) error {
	rv := &models.RawVisit{
		VisitStartedAt: cr.DialStartTime,
		VisitEndedAt:   cr.DialEndTime,
		DialDuration:   null.StringFrom(fmt.Sprintf("%f seconds", cr.DialDuration().Seconds())),
		Type:           models.VisitTypeDial,
		PeerMultiHash:  cr.Peer.ID.Pretty(),
		MultiAddresses: utils.MaddrsToAddrs(cr.Peer.Addrs),
	}

	if cr.Error != nil {
		rv.Error = null.StringFrom(cr.DialError)
		rv.ErrorMessage = null.StringFrom(cr.Error.Error())
	}

	return s.dbc.InsertRawVisit(ctx, rv)
}
