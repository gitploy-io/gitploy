package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/event"
	"github.com/hanjunlee/gitploy/vo"
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
			es, err := i.ListEventsGreaterThanTime(ctx, t.Add(-period))
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

func (i *Interactor) ListUsersOfEvent(ctx context.Context, e *ent.Event) ([]*ent.User, error) {
	if e.Kind == event.KindDeployment {
		d, err := i.Store.FindDeploymentByID(ctx, e.DeploymentID)
		if err != nil {
			return nil, fmt.Errorf("It has failed to find the deployment: %w", err)
		}

		return []*ent.User{d.Edges.User}, nil
	}

	// Notify to who has the request of approval.
	if e.Kind == event.KindApproval && e.Type == event.TypeCreated {
		a, err := i.Store.FindApprovalByID(ctx, e.ApprovalID)
		if err != nil {
			return nil, fmt.Errorf("It has failed to find the approval: %w", err)
		}

		return []*ent.User{a.Edges.User}, nil
	}

	// Notify to who has requested the approval.
	if e.Kind == event.KindApproval && e.Type == event.TypeUpdated {
		a, err := i.Store.FindApprovalByID(ctx, e.ApprovalID)
		if err != nil {
			return nil, fmt.Errorf("It has failed to find the approval: %w", err)
		}

		d, err := i.Store.FindDeploymentByID(ctx, a.DeploymentID)
		if err != nil {
			return nil, fmt.Errorf("It has failed to find the deployment of the approval: %w", err)
		}

		return []*ent.User{d.Edges.User}, nil
	}

	return nil, fmt.Errorf("It is out of use-cases.")
}

func (i *Interactor) ConvertEventToNotification(ctx context.Context, e *ent.Event) (*vo.Notification, error) {
	if e.Kind == event.KindDeployment {
		d, err := i.Store.FindDeploymentByID(ctx, e.DeploymentID)
		if err != nil {
			return nil, fmt.Errorf("It has failed to find the deployment: %w", err)
		}

		return &vo.Notification{
			Kind:       vo.NotificationKind(e.Kind),
			Type:       vo.NotificationType(e.Type),
			Repo:       mapRepoToRepoData(d.Edges.Repo),
			Deployment: mapDeploymentToDeploymentData(d),
		}, nil
	}

	if e.Kind == event.KindApproval {
		a, err := i.Store.FindApprovalByID(ctx, e.ApprovalID)
		if err != nil {
			return nil, fmt.Errorf("It has failed to find the approval: %w", err)
		}

		d, err := i.Store.FindDeploymentByID(ctx, a.DeploymentID)
		if err != nil {
			return nil, fmt.Errorf("It has failed to find the deployment of the approval: %w", err)
		}

		return &vo.Notification{
			Kind:       vo.NotificationKind(e.Kind),
			Type:       vo.NotificationType(e.Type),
			Repo:       mapRepoToRepoData(d.Edges.Repo),
			Deployment: mapDeploymentToDeploymentData(d),
			Approval:   mapApprovalToApprovalData(a),
		}, nil
	}

	return nil, fmt.Errorf("It is out of use-cases.")
}

func mapRepoToRepoData(r *ent.Repo) *vo.RepoData {
	return &vo.RepoData{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func mapDeploymentToDeploymentData(d *ent.Deployment) *vo.DeploymentData {
	return &vo.DeploymentData{
		Number: d.Number,
		Type:   d.Type.String(),
		Ref:    d.Ref,
		Env:    d.Env,
		Status: d.Status.String(),
		Login:  d.Edges.User.Login,
	}
}

func mapApprovalToApprovalData(a *ent.Approval) *vo.ApprovalData {
	return &vo.ApprovalData{
		Status: a.Status.String(),
		Login:  a.Edges.User.Login,
	}
}
