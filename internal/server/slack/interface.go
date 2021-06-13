package slack

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		FindUser() (*ent.User, error)
		FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)
	}
)
