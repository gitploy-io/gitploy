package stream

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor interface {
		SubscribeEvent(fn func(e *ent.Event)) error
		UnsubscribeEvent(fn func(e *ent.Event)) error
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindApprovalByID(ctx context.Context, id int) (*ent.Approval, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)
	}
)
