package interactor

import (
	"context"
	"math/rand"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"go.uber.org/zap"
)

const (
	eventStream       = "gitploy:stream"
	eventNotification = "gitploy:notification"
)

func (i *Interactor) polling(stop <-chan struct{}) {
	const (
		max = 4
	)

	ctx := context.Background()

	// Notify if user connect Gitploy with Chat (e.g. Slack, Microsoft Teams).
	i.events.SubscribeAsync(eventNotification, func(n *ent.Notification) {
		i.store.SetNotificationDone(ctx, n)

		if err := i.Notify(n); err != nil {
			i.log.Error("failed to notify.", zap.Error(err))
		}
	}, false)

	// polling
	ticker := time.NewTicker(time.Millisecond * 1000 * time.Duration(randint(2, max)))

L:
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				ticker.Stop()
				break L
			}
		case t := <-ticker.C:
			ns, err := i.store.ListNotificationsFromTime(ctx, t.Add(-time.Second*max))
			if err != nil {
				i.log.Named("polling").Error("failed to read notifications.", zap.Error(err))
				continue
			}

			for _, n := range ns {
				if n.Notified {
					continue
				}

				i.publish(n)
				i.log.Named("polling").Debug("publish the notification.", zap.Int("id", n.ID))
			}
		}
	}
}

func (i *Interactor) publish(n *ent.Notification) {
	i.events.Publish(eventStream, n)
	i.events.Publish(eventNotification, n)
}

func (i *Interactor) Subscribe(fn func(*ent.Notification)) {
	i.events.SubscribeAsync(eventStream, fn, false)
}

func (i *Interactor) Notify(n *ent.Notification) error {
	return nil
}

func randint(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}
