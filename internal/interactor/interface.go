package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Store interface {
		FindUser() (*ent.User, error)
		FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
	}

	SCM interface {
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)
	}
)
