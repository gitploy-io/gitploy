package repos

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Store interface {
		ListRepos(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Repo, error)
		FindRepo(ctx context.Context, u *ent.User, id string) (*ent.Repo, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
	}

	SCM interface {
		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error)
	}
)
