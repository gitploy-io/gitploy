package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/event"
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
