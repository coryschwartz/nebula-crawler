package provide

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	u "github.com/ipfs/go-ipfs-util"
	"github.com/libp2p/go-libp2p-core/peer"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	log "github.com/sirupsen/logrus"
)

// Measurement keeps track of the starting conditions and results of an experiment.
type Measurement struct {
	// The libp2p host peer identity of the provider
	providerID peer.ID

	// The libp2p host peer identity of the requester
	requesterID peer.ID

	// The random data that was provided
	content *Content

	// When did the provider start searching for peers for the provider record
	startTime time.Time

	// When did the provider end searching for peers for the provider record
	endTime time.Time

	// All events that occurred during the whole process
	events []Event

	// Keeps track of peers that were involved in the provide process. Since
	// events are dispatched for all dials regardless of whether they were
	// necessary for the provide process or e.g., the routing table refresh.
	// Therefore, the provider registers for query events that happened during the
	// provide process and keep track of all involved peers. This list is
	// ultimately used to remove all events from the events field that targeted
	// peers not relevant for the provide process.
	involved sync.Map

	// monitored represents the list of peers that the requester periodically
	// asked for provider records.
	monitored []peer.ID
}

// filterEvents removes all events that are related to remote peers that were not involved in the Provide process.
func (m *Measurement) filterEvents() {
	// Also mark the monitored peers as involved.
	// This is necessary if the requester found peers
	// that the provider didn't add provider records to.
	for _, mon := range m.monitored {
		m.involved.Store(mon, struct{}{})
	}

	var filtered []Event
	for _, event := range m.events {
		if _, isInvolved := m.involved.Load(event.RemoteID()); isInvolved {
			filtered = append(filtered, event)
		}
	}
	m.events = filtered
}

// checkIntegrity makes sure that for every start event there is a corresponding end event.
func (m *Measurement) checkIntegrity() bool {
	states := map[peer.ID]map[SpanType]int{}
	for _, event := range m.events {

		if _, found := states[event.RemoteID()]; !found {
			states[event.RemoteID()] = map[SpanType]int{}
		}

		if event.IsStart() {
			states[event.RemoteID()][event.Span()] += 1
		} else {
			states[event.RemoteID()][event.Span()] -= 1
		}
	}

	for p, scopes := range states {
		for scope, count := range scopes {
			if count != 0 {
				log.Warnln(p, scope)
				return false
			}
		}
	}

	return true
}

// detectSpans loops through all events and tries to detect corresponding start and end events
// to construct a span entity. E.g. there will be multiple simultaneous dials where only one
// will succeed. The other dials end with an error, yet the dial attempt is successful.
func (m *Measurement) detectSpans() []Span {
	// S
	type SpanState struct {
		Start time.Time
		Count int
	}
	var spans []Span
	spanStates := map[peer.ID]map[SpanType]*SpanState{}
	// providerSpanStates := map[peer.ID]map[SpanType]*SpanState{}
	// requesterSpanStates := map[peer.ID]map[SpanType]*SpanState{}

	// Loop through all events
	for _, event := range m.events {

		// Check if we are already tracking this peer - if not create a map for it.
		if _, found := spanStates[event.RemoteID()]; !found {
			spanStates[event.RemoteID()] = map[SpanType]*SpanState{}
		}

		// Check if it's the start of a span or not
		if event.IsStart() {
			// If we don't have an "open" start event in our state we create the SpanState
			// and set the count to 1 (number of start events)
			// If we already have an open start event in our state we just increment the
			// counter to keep track how many open events we came across
			if _, found := spanStates[event.RemoteID()][event.Span()]; !found {
				spanStates[event.RemoteID()][event.Span()] = &SpanState{
					Start: event.TimeStamp(),
					Count: 1,
				}
			} else {
				spanStates[event.RemoteID()][event.Span()].Count += 1
			}
		} else {
			// If received an end event while there is no open span state we just do nothing
			if _, found := spanStates[event.RemoteID()][event.Span()]; !found {
				continue
			}

			// Decrement the count
			spanStates[event.RemoteID()][event.Span()].Count -= 1

			// If the end event contains an error an this was not the last event for this
			// open span we do nothing and wait for the last or a successful end event
			if event.Error() != nil && spanStates[event.RemoteID()][event.Span()].Count != 0 {
				continue
			}

			// Create span
			spans = append(spans, Span{
				RelStart:  spanStates[event.RemoteID()][event.Span()].Start.Sub(m.startTime).Seconds(),
				DurationS: event.TimeStamp().Sub(spanStates[event.RemoteID()][event.Span()].Start).Seconds(),
				Start:     spanStates[event.RemoteID()][event.Span()].Start,
				End:       event.TimeStamp(),
				PeerID:    event.RemoteID(),
				Type:      event.Span(),
				Error:     event.Error().Error(),
			})

			// Delete span state so that subsequent events of this span type can be tracked again.
			delete(spanStates[event.RemoteID()], event.Span())
		}
	}

	return spans
}

func (m *Measurement) saveSpans(spans []Span) error {
	spanMap := map[string][]Span{}

	for _, span := range spans {
		if _, found := spanMap[span.PeerID.Pretty()]; !found {
			spanMap[span.PeerID.Pretty()] = []Span{}
		}
		spanMap[span.PeerID.Pretty()] = append(spanMap[span.PeerID.Pretty()], span)
	}

	data, err := json.MarshalIndent(spanMap, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal spans")
	}

	f, err := os.Create(m.prefix() + "_spans.json")
	if err != nil {
		return errors.Wrap(err, "creating spans file")
	}

	_, err = f.Write(data)
	if err != nil {
		return errors.Wrap(err, "writing spans file")
	}

	return f.Close()
}

func (m *Measurement) saveMeasurementInfo() error {
	ei := MeasurementInfo{
		StartedAt:     m.startTime,
		EndedAt:       m.endTime,
		ContentID:     m.content.cid.String(),
		ProviderID:    m.providerID.Pretty(),
		ProviderDist:  hex.EncodeToString(u.XOR(kbucket.ConvertPeerID(m.providerID), kbucket.ConvertKey(string(m.content.mhash)))),
		RequesterID:   m.requesterID.Pretty(),
		RequesterDist: hex.EncodeToString(u.XOR(kbucket.ConvertPeerID(m.requesterID), kbucket.ConvertKey(string(m.content.mhash)))),
	}

	data, err := json.MarshalIndent(ei, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal experiment info")
	}

	f, err := os.Create(m.prefix() + "_measurement_info.json")
	if err != nil {
		return errors.Wrap(err, "creating experiment info file")
	}

	_, err = f.Write(data)
	if err != nil {
		return errors.Wrap(err, "writing experiment info file")
	}

	return f.Close()
}

type Span struct {
	RelStart  float64
	DurationS float64
	Start     time.Time
	End       time.Time
	PeerID    peer.ID
	Type      SpanType
	Error     string
}

type PeerInfo struct {
	ID             peer.ID
	AgentVersion   string
	XORDistance    float64
	DiscoveredAt   time.Time
	DiscoveredFrom peer.ID
}

type MeasurementInfo struct {
	StartedAt     time.Time
	EndedAt       time.Time
	ContentID     string
	ProviderID    string
	ProviderDist  string
	RequesterID   string
	RequesterDist string
	// DialCount     int
	// Content           *Content
	// RoutingTableStart int
	// Hops       int
	// HydraCount int
}

func (m *Measurement) prefix() string {
	t := m.startTime
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
