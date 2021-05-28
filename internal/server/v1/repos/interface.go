package repos

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Store interface {
		ListRepos(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Repo, error)
		FindRepo(ctx context.Context, u *ent.User, id string) (*ent.Repo, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
	}

	SCM interface{}
)
