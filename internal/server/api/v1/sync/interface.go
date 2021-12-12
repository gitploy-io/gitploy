//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package sync

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	Interactor interface {
		ListRemoteRepos(ctx context.Context, u *ent.User) ([]*extent.RemoteRepo, error)
		IsEntryOrg(ctx context.Context, namespace string) bool
		SyncRemoteRepo(ctx context.Context, u *ent.User, re *extent.RemoteRepo, t time.Time) error
		DeletePermsOfUserLessThanSyncedAt(ctx context.Context, u *ent.User, t time.Time) (int, error)
	}
)
