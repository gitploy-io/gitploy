package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/hanjunlee/gitploy/ent"
)

func (i *Interactor) FindUser() (*ent.User, error) {
	return i.store.FindUser()
}

func (i *Interactor) FindUserByHash(ctx context.Context, hash string) (*ent.User, error) {
	return i.store.FindUserByHash(ctx, hash)
}

func (i *Interactor) Sync(ctx context.Context, u *ent.User) error {
	perms, err := i.scm.GetAllPermsWithRepo(ctx, u.Token)
	if err != nil {
		return fmt.Errorf("failed to get all permissions: %w", err)
	}
	i.log.Debug("get all permissions.")

	sync := time.Now()

	for _, perm := range perms {
		re := perm.Edges.Repo
		if err := i.store.SyncPerm(ctx, u, perm, sync); err != nil {
			return fmt.Errorf("failed to sync with the \"%s\" repo: %w", re.Name, err)
		}
	}

	return nil
}
