package external

import (
	"context"
	"sort"
	"sync"

	"github.com/nao1215/honeycomb/app/service"
	"github.com/nbd-wtf/go-nostr"
)

var _ service.EventsLister = (*EventsLister)(nil)

// EventsLister is the external service for listing events.
type EventsLister struct{}

// NewEventsLister creates a new EventsLister.
func NewEventsLister() *EventsLister {
	return &EventsLister{}
}

// ListEvents is the external service for listing events.
func (p *EventsLister) ListEvents(ctx context.Context, input *service.EventsListerInput) (*service.EventsListerOutput, error) {
	filter := input.Filter
	relay := input.Relay
	var m sync.Map

	// TODO: Use multiple relays (goroutine).
	events, err := relay.QuerySync(ctx, filter)
	if err != nil {
		return nil, err
	}
	for _, ev := range events {
		if _, ok := m.Load(ev.ID); !ok {
			if ev.Kind == nostr.KindEncryptedDirectMessage || ev.Kind == nostr.KindCategorizedBookmarksList {
				// TODO: Decrypt the event.
				continue
			}
			m.LoadOrStore(ev.ID, ev)
			if len(filter.IDs) == 1 {
				ctx.Done()
				break
			}
		}
	}

	keys := make([]string, 0)
	m.Range(func(k, _ any) bool {
		keyString, ok := k.(string)
		if !ok {
			return false
		}
		keys = append(keys, keyString)
		return true
	})

	sort.Slice(keys, func(i, j int) bool {
		left, ok := m.Load(keys[i])
		if !ok {
			return false
		}

		right, ok := m.Load(keys[j])
		if !ok {
			return false
		}

		le, ok := left.(*nostr.Event)
		if !ok {
			return false
		}

		re, ok := right.(*nostr.Event)
		if !ok {
			return false
		}

		return le.CreatedAt.Time().Before(re.CreatedAt.Time())
	})

	nostrEvents := make([]*nostr.Event, 0, len(keys))
	for _, key := range keys {
		vv, ok := m.Load(key)
		if !ok {
			continue
		}
		e, ok := vv.(*nostr.Event)
		if !ok {
			continue
		}
		nostrEvents = append(nostrEvents, e)
	}

	return &service.EventsListerOutput{
		Events: nostrEvents,
	}, nil
}
