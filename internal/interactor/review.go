package interactor

import (
	"context"

	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/event"
)

type (
	ReviewInteractor service

	// ReviewStore defines operations for working with reviews.
	ReviewStore interface {
		SearchReviews(ctx context.Context, u *ent.User) ([]*ent.Review, error)
		ListReviews(ctx context.Context, d *ent.Deployment) ([]*ent.Review, error)
		FindReviewOfUser(ctx context.Context, u *ent.User, d *ent.Deployment) (*ent.Review, error)
		FindReviewByID(ctx context.Context, id int) (*ent.Review, error)
		// CreateReview creates a review of which status is pending.
		CreateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error)
		// UpdateReview update the status and comment of the review.
		UpdateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error)
	}
)

// RespondReview update the status of review.
func (i *ReviewInteractor) RespondReview(ctx context.Context, rv *ent.Review) (*ent.Review, error) {
	rv, err := i.store.UpdateReview(ctx, rv)
	if err != nil {
		return nil, err
	}

	if _, err := i.store.CreateEvent(ctx, &ent.Event{
		Kind:     event.KindReview,
		Type:     event.TypeCreated,
		ReviewID: rv.ID,
	}); err != nil {
		i.log.Error("Failed to create a review event.", zap.Error(err))
	}

	return rv, nil
}
