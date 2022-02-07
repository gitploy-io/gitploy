package interactor

import (
	"context"
	"time"

	evbus "github.com/asaskevich/EventBus"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/model/ent"
)

const (
	gitployEvent = "gitploy:event"
)

type (
	EventsInteractor struct {
		*service

		events evbus.Bus
	}
)

func (i *EventsInteractor) SubscribeEvent(fn func(e *ent.Event)) error {
	return i.events.SubscribeAsync(gitployEvent, fn, false)
}

func (i *EventsInteractor) UnsubscribeEvent(fn func(e *ent.Event)) error {
	return i.events.Unsubscribe(gitployEvent, fn)
}

func (i *EventsInteractor) runPublishingEvents(stop <-chan struct{}) {
	ctx := context.Background()

	// Read events periodically and publish to subscribers.
	period := time.Second * 2
	ticker := time.NewTicker(period)

L:
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				ticker.Stop()
				break L
			}

		case t := <-ticker.C:
			es, err := i.store.ListEventsGreaterThanTime(ctx, t.Add(-period).UTC())
			if err != nil {
				i.log.Error("It has failed to read events.", zap.Error(err))
				continue
			}

			for _, e := range es {
				i.events.Publish(gitployEvent, e)
				i.log.Debug("Publish the event.", zap.Int("event_id", e.ID))
			}
		}
	}
}
