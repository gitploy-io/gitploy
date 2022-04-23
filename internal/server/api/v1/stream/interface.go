package stream

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	Interactor interface {
		SubscribeEvent(fn func(e *ent.Event)) error
		UnsubscribeEvent(fn func(e *ent.Event)) error
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindDeploymentStatusByID(ctx context.Context, id int) (*ent.DeploymentStatus, error)
		FindReviewByID(ctx context.Context, id int) (*ent.Review, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)
	}
)
