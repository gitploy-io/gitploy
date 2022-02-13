package search

import (
	"context"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	Interactor interface {
		SearchDeploymentsOfUser(ctx context.Context, u *ent.User, opt *i.SearchDeploymentsOfUserOptions) ([]*ent.Deployment, error)
		SearchReviews(ctx context.Context, u *ent.User) ([]*ent.Review, error)
	}
)
