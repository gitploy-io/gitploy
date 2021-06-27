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
	eventStream = "gitploy:stream"
	eventChat   = "gitploy:chat"
)

func (i *Interactor) polling(stop <-chan struct{}) {
	ctx := context.Background()
	log := i.log.Named("polling")

	// Subscribe events to notify by Chat.
	i.events.SubscribeAsync(eventChat, func(cu *ent.ChatUser, n *ent.Notification) {
		if err := i.notifyByChat(ctx, cu, n); err != nil {
			i.log.Error("failed to notify by chat.", zap.Error(err))
		}
	}, false)

	// polling with the random period to escape the conflict; 3s - 4s
	ticker := time.NewTicker(time.Millisecond * 100 * time.Duration(randint(30, 40)))

L:
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				ticker.Stop()
				break L
			}
		case t := <-ticker.C:
			ns, err := i.ListNotificationsWithEdgesFromTime(ctx, t.Add(-time.Second*4))
			if err != nil {
				log.Error("failed to read notifications.", zap.Error(err))
				continue
			}

			for _, n := range ns {
				if n.Notified {
					continue
				}

				i.publish(ctx, n)
				i.log.Debug("publish the notification event.", zap.Int("notification_id", n.ID))
			}
		}
	}
}

// publish notification to Chat event if it is connected,
// and it updates notified field true,
// whereas if not connected, it publishs to stream without update.
func (i *Interactor) publish(ctx context.Context, n *ent.Notification) error {
	u, err := i.FindUserWithChatUserByID(ctx, n.UserID)
	if err != nil {
		return err
	}

	if cu := u.Edges.ChatUser; cu != nil {
		i.events.Publish(eventChat, cu, n)
		n, _ = i.SetNotificationNotified(ctx, n)
	}

	i.events.Publish(eventStream, u, n)
	i.SetNotificationNotified(ctx, n)
	return nil
}

func (i *Interactor) notifyByChat(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	switch n.Type {
	case notification.TypeDeployment:
		return i.notifyDeploymentByChat(ctx, cu, n)
	default:
		return nil
	}
}

func (i *Interactor) notifyDeploymentByChat(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	d, err := i.FindDeploymentWithEdgesByID(ctx, n.DeploymentID)
	if err != nil {
		return err
	}

	return i.NotifyDeployment(ctx, cu, d)
}

// Notify enqueues a new notification for resource.
func (i *Interactor) Publish(ctx context.Context, iface interface{}) error {
	switch r := iface.(type) {
	case *ent.Deployment:
		return i.createDeploymentNotification(ctx, r)
	}

	return fmt.Errorf("failed to notify")
}

// createDeploymentNotification enqueues notifications for the deployer.
func (i *Interactor) createDeploymentNotification(ctx context.Context, d *ent.Deployment) error {
	_, err := i.CreateNotification(ctx, &ent.Notification{
		Type:         notification.TypeDeployment,
		UserID:       d.UserID,
		DeploymentID: d.ID,
	})
	return err
}

func (i *Interactor) Subscribe(fn func(u *ent.User, n *ent.Notification)) error {
	return i.events.SubscribeAsync(eventStream, fn, false)
}

func (i *Interactor) Unsubscribe(fn func(u *ent.User, n *ent.Notification)) error {
	return i.events.Unsubscribe(eventStream, fn)
}

func randint(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}
