package sync

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	SCMHandler interface {
		GetAllPermsWithRepo(ctx context.Context, token string) ([]*ent.Perm, error)
	}

	// StoreHandler store remote repositories into local.
	StoreHandler interface {
		GetUserByHash(ctx context.Context, hash string) (*ent.User, error)
		SyncPerm(ctx context.Context, user *ent.User, perm *ent.Perm, sync time.Time) error
	}
)
