package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/event"
	"go.uber.org/zap"
)

func (i *Interactor) UpdateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error) {
	rv, err := i.Store.UpdateReview(ctx, rv)
	if err != nil {
		return nil, err
	}

	if _, err := i.Store.CreateEvent(ctx, &ent.Event{
		Kind:     event.KindReview,
		Type:     event.TypeUpdated,
		ReviewID: rv.ID,
	}); err != nil {
		i.log.Error("Failed to create the event.", zap.Error(err))
	}

	return rv, nil
}

func (i *Interactor) requestReviewByLogin(ctx context.Context, d *ent.Deployment, login string) (*ent.Review, error) {
	u, err := i.Store.FindUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	rv, err := i.Store.CreateReview(ctx, &ent.Review{
		DeploymentID: d.ID,
		UserID:       u.ID,
	})
	if err != nil {
		return nil, err
	}

	if _, err := i.Store.CreateEvent(ctx, &ent.Event{
		Kind:     event.KindReview,
		Type:     event.TypeCreated,
		ReviewID: rv.ID,
	}); err != nil {
		i.log.Error("Failed to create the event.", zap.Error(err))
	}

	return rv, nil
}
