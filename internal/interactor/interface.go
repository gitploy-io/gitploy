package interactor

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Store interface {
		FindUser() (*ent.User, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		FindUserByID(ctx context.Context, id string) (*ent.User, error)
		CreateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		FindRepo(ctx context.Context, u *ent.User, id string) (*ent.Repo, error)
		FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		FindPerm(ctx context.Context, u *ent.User, repoID string) (*ent.Perm, error)
		SyncPerm(ctx context.Context, user *ent.User, perm *ent.Perm, sync time.Time) error
	}

	SCM interface {
		GetUser(ctx context.Context, token string) (*ent.User, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)
		GetAllPermsWithRepo(ctx context.Context, token string) ([]*ent.Perm, error)
	}
)
