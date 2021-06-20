package interactor

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/notification"
	"go.uber.org/zap"
)

const (
	eventStream       = "gitploy:stream"
	eventNotification = "gitploy:notification"
)

func (i *Interactor) polling(stop <-chan struct{}) {
	ctx := context.Background()

	i.events.SubscribeAsync(eventNotification, func(n *ent.Notification) {
		i.store.SetNotificationDone(ctx, n)

		if err := i.notifyByChat(ctx, n); err != nil {
			i.log.Error("failed to notify.", zap.Error(err))
		}
	}, false)

	// polling with the random period to escape the conflict; 2s - 4s
	ticker := time.NewTicker(time.Millisecond * 100 * time.Duration(randint(20, 40)))

L:
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				ticker.Stop()
				break L
			}
		case t := <-ticker.C:
			ns, err := i.store.ListNotificationsFromTime(ctx, t.Add(-time.Second*4))
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

// messageByChat notify by Chat if it is connected with Gitploy (e.g. Slack, Microsoft Teams).
func (i *Interactor) notifyByChat(ctx context.Context, n *ent.Notification) error {
	switch n.Type {
	case notification.TypeDeployment:
		return i.notifyDeploymentByChat(ctx, n)
	default:
		return nil
	}
}

func (i *Interactor) notifyDeploymentByChat(ctx context.Context, n *ent.Notification) error {
	u, err := i.store.FindUserWithChatUserByID(ctx, n.UserID)
	if err != nil {
		return err
	}
	if u.Edges.ChatUser == nil {
		return nil
	}

	d, err := i.store.FindDeploymentWithEdgesByID(ctx, n.ID)
	if err != nil {
		return err
	}

	return i.chat.NotifyDeployment(ctx, u.Edges.ChatUser, d)
}

func (i *Interactor) publish(n *ent.Notification) {
	i.events.Publish(eventStream, n)
	i.events.Publish(eventNotification, n)
}

func (i *Interactor) Subscribe(fn func(*ent.Notification)) {
	i.events.SubscribeAsync(eventStream, fn, false)
}

// Notify enqueues a new notification for resource.
func (i *Interactor) Notify(ctx context.Context, iface interface{}) (*ent.Notification, error) {
	switch r := iface.(type) {
	case *ent.Deployment:
		return i.createDeploymentNotification(ctx, r)
	}

	return nil, fmt.Errorf("failed to notify")
}

func (i *Interactor) createDeploymentNotification(ctx context.Context, d *ent.Deployment) (*ent.Notification, error) {
	return i.store.CreateNotification(ctx, &ent.Notification{
		Type:       notification.TypeDeployment,
		ResourceID: d.ID,
		Notified:   false,
		UserID:     d.UserID,
	})
}

func randint(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}
