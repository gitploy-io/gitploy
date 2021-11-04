package search

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
)

type (
	Interactor interface {
		SearchDeployments(ctx context.Context, u *ent.User, s []deployment.Status, owned bool, from time.Time, to time.Time, page, perPage int) ([]*ent.Deployment, error)
		SearchReviews(ctx context.Context, u *ent.User) ([]*ent.Review, error)
	}
)
