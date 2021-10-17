package interactor

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"go.uber.org/zap"
)

const (
	gitployEvent = "gitploy:event"
)

func (i *Interactor) runPublishingEvents(stop <-chan struct{}) {
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
			es, err := i.ListEventsGreaterThanTime(ctx, t.Add(-period).UTC())
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

func (i *Interactor) SubscribeEvent(fn func(e *ent.Event)) error {
	return i.events.SubscribeAsync(gitployEvent, fn, false)
}

func (i *Interactor) UnsubscribeEvent(fn func(e *ent.Event)) error {
	return i.events.Unsubscribe(gitployEvent, fn)
}
