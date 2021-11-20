package store

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/review"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *Store) SearchReviews(ctx context.Context, u *ent.User) ([]*ent.Review, error) {
	rvs, err := s.c.Review.
		Query().
		Where(
			review.And(
				review.UserID(u.ID),
				review.HasDeploymentWith(deployment.StatusEQ(deployment.StatusWaiting)),
				review.StatusEQ(review.StatusPending),
			),
		).
		WithUser().
		WithDeployment(func(dq *ent.DeploymentQuery) {
			dq.
				WithRepo().
				WithUser()
		}).
		Order(ent.Desc(review.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return rvs, nil
}

func (s *Store) ListReviews(ctx context.Context, d *ent.Deployment) ([]*ent.Review, error) {
	rvs, err := s.c.Review.
		Query().
		Where(
			review.DeploymentIDEQ(d.ID),
		).
		WithUser().
		WithDeployment().
		All(ctx)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return rvs, nil
}

func (s *Store) FindReviewByID(ctx context.Context, id int) (*ent.Review, error) {
	rv, err := s.c.Review.
		Query().
		Where(
			review.IDEQ(id),
		).
		WithUser().
		WithDeployment().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The review is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return rv, nil
}

func (s *Store) FindReviewOfUser(ctx context.Context, u *ent.User, d *ent.Deployment) (*ent.Review, error) {
	rv, err := s.c.Review.
		Query().
		Where(
			review.DeploymentIDEQ(d.ID),
			review.UserIDEQ(u.ID),
		).
		WithUser().
		WithDeployment().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The review is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return rv, nil
}

func (s *Store) CreateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error) {
	rv, err := s.c.Review.
		Create().
		SetComment(rv.Comment).
		SetDeploymentID(rv.DeploymentID).
		SetUserID(rv.UserID).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("Failed to create a review. The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err,
		)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return s.FindReviewByID(ctx, rv.ID)
}

func (s *Store) UpdateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error) {
	rv, err := s.c.Review.
		UpdateOne(rv).
		SetStatus(rv.Status).
		SetComment(rv.Comment).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("Failed to update the review. The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err,
		)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return s.FindReviewByID(ctx, rv.ID)
}
