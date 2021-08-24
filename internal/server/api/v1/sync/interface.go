package sync

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		ListRemoteRepos(ctx context.Context, u *ent.User) ([]*vo.RemoteRepo, error)
		SyncRemoteRepo(ctx context.Context, u *ent.User, re *vo.RemoteRepo) error
		DeletePermsOfUserLessThanUpdatedAt(ctx context.Context, u *ent.User, t time.Time) error
	}
)
