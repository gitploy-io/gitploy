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

		ListRepos(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error)
		ListSortedRepos(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error)
		FindRepo(ctx context.Context, u *ent.User, id string) (*ent.Repo, error)
		FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		Activate(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		Deactivate(ctx context.Context, r *ent.Repo) (*ent.Repo, error)

		FindPerm(ctx context.Context, u *ent.User, repoID string) (*ent.Perm, error)
		SyncPerm(ctx context.Context, user *ent.User, perm *ent.Perm, sync time.Time) error
	}

	SCM interface {
		GetUser(ctx context.Context, token string) (*ent.User, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)
		GetAllPermsWithRepo(ctx context.Context, token string) ([]*ent.Perm, error)
		CreateWebhook(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (int64, error)
		DeleteWebhook(ctx context.Context, u *ent.User, r *ent.Repo, id int64) error
	}
)
