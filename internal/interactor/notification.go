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
	eventNotification = "gitploy:notification"
)

func (i *Interactor) polling(stop <-chan struct{}) {
	ctx := context.Background()
	log := i.log.Named("polling")

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
			ns, err := i.ListNotificationsFromTime(ctx, t.Add(-time.Second*4))
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
	u, err := i.FindUserByID(ctx, n.UserID)
	if err != nil {
		return err
	}

	i.events.Publish(eventNotification, u, n)
	i.setNotificationNotified(ctx, n)
	return nil
}

func (i *Interactor) setNotificationNotified(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	n.Notified = true
	return i.UpdateNotification(ctx, n)
}

func (i *Interactor) Publish(ctx context.Context, typ notification.Type, r *ent.Repo, d *ent.Deployment, a *ent.Approval) error {
	switch typ {
	case notification.TypeDeploymentCreated:
		return i.publishDeploymentCreated(ctx, r, d)

	case notification.TypeDeploymentUpdated:
		return i.publishDeploymentUpdated(ctx, r, d)

	case notification.TypeApprovalRequested:
		return i.publishApprovalRequested(ctx, r, d, a)

	case notification.TypeApprovalResponded:
		return i.publishApprovalResponded(ctx, r, d, a)

	default:
		return nil
	}
}

// publishDeploymentCreated publish the notification to the deployer.
func (i *Interactor) publishDeploymentCreated(ctx context.Context, r *ent.Repo, d *ent.Deployment) error {
	u, err := i.FindUserByID(ctx, d.UserID)
	if err != nil {
		return fmt.Errorf("the deployer is not found.")
	}

	if _, err := i.CreateNotification(ctx, &ent.Notification{
		Type:             notification.TypeDeploymentCreated,
		RepoNamespace:    r.Namespace,
		RepoName:         r.Name,
		DeploymentNumber: d.Number,
		DeploymentType:   string(d.Type),
		DeploymentRef:    d.Ref,
		DeploymentEnv:    d.Env,
		DeploymentStatus: string(d.Status),
		DeploymentLogin:  u.Login,
		UserID:           u.ID,
	}); err != nil {
		return err
	}

	return nil
}

// publishDeploymentUpdated publish a notification to the deployer.
func (i *Interactor) publishDeploymentUpdated(ctx context.Context, r *ent.Repo, d *ent.Deployment) error {
	u, err := i.FindUserByID(ctx, d.UserID)
	if err != nil {
		return fmt.Errorf("the deployer is not found.")
	}

	if _, err := i.CreateNotification(ctx, &ent.Notification{
		Type:             notification.TypeDeploymentUpdated,
		RepoNamespace:    r.Namespace,
		RepoName:         r.Name,
		DeploymentNumber: d.Number,
		DeploymentType:   string(d.Type),
		DeploymentRef:    d.Ref,
		DeploymentEnv:    d.Env,
		DeploymentStatus: string(d.Status),
		DeploymentLogin:  u.Login,
		UserID:           u.ID,
	}); err != nil {
		return err
	}

	return nil
}

// publishApprovalRequested publish notifications to who receives a request of approval.
func (i *Interactor) publishApprovalRequested(ctx context.Context, r *ent.Repo, d *ent.Deployment, a *ent.Approval) error {
	du, err := i.FindUserByID(ctx, d.UserID)
	if err != nil {
		return fmt.Errorf("the deployer is not found.")
	}

	au, err := i.FindUserByID(ctx, a.UserID)
	if err != nil {
		return fmt.Errorf("the deployer is not found.")
	}

	if _, err := i.CreateNotification(ctx, &ent.Notification{
		Type:             notification.TypeApprovalRequested,
		RepoNamespace:    r.Namespace,
		RepoName:         r.Name,
		DeploymentNumber: d.Number,
		DeploymentType:   string(d.Type),
		DeploymentRef:    d.Ref,
		DeploymentEnv:    d.Env,
		DeploymentStatus: string(d.Status),
		DeploymentLogin:  du.Login,
		ApprovalStatus:   string(a.Status),
		ApprovalLogin:    au.Login,
		UserID:           a.UserID,
	}); err != nil {
		return err
	}

	return nil
}

// publishApprovalResponded publish notifications to who has deployed.
func (i *Interactor) publishApprovalResponded(ctx context.Context, r *ent.Repo, d *ent.Deployment, a *ent.Approval) error {
	du, err := i.FindUserByID(ctx, d.UserID)
	if err != nil {
		return fmt.Errorf("the deployer is not found.")
	}

	au, err := i.FindUserByID(ctx, a.UserID)
	if err != nil {
		return fmt.Errorf("the deployer is not found.")
	}

	if _, err := i.CreateNotification(ctx, &ent.Notification{
		Type:             notification.TypeApprovalResponded,
		RepoNamespace:    r.Namespace,
		RepoName:         r.Name,
		DeploymentNumber: d.Number,
		DeploymentType:   string(d.Type),
		DeploymentRef:    d.Ref,
		DeploymentEnv:    d.Env,
		DeploymentStatus: string(d.Status),
		DeploymentLogin:  du.Login,
		ApprovalStatus:   string(a.Status),
		ApprovalLogin:    au.Login,
		UserID:           d.UserID,
	}); err != nil {
		return err
	}

	return nil
}

func (i *Interactor) Subscribe(fn func(u *ent.User, n *ent.Notification)) error {
	return i.events.SubscribeAsync(eventNotification, fn, false)
}

func (i *Interactor) Unsubscribe(fn func(u *ent.User, n *ent.Notification)) error {
	return i.events.Unsubscribe(eventNotification, fn)
}

func randint(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}
