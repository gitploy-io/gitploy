package stream

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
)

type (
	Interactor interface {
		SubscribeEvent(fn func(e *ent.Event)) error
		UnsubscribeEvent(fn func(e *ent.Event)) error
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindReviewByID(ctx context.Context, id int) (*ent.Review, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)
	}
)
