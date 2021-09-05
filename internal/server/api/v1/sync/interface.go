//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package sync

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	Interactor interface {
		ListRemoteRepos(ctx context.Context, u *ent.User) ([]*vo.RemoteRepo, error)
		IsEntryOrg(ctx context.Context, namespace string) bool
		SyncRemoteRepo(ctx context.Context, u *ent.User, re *vo.RemoteRepo) error
		DeletePermsOfUserLessThanUpdatedAt(ctx context.Context, u *ent.User, t time.Time) (int, error)
	}
)
