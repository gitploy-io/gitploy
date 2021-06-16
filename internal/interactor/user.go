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

func (i *Interactor) SaveSCMUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	_, err := i.store.FindUserByID(ctx, u.ID)
	if ent.IsNotFound(err) {
		u, _ = i.store.CreateUser(ctx, u)
	} else if err != nil {
		return nil, err
	}

	return i.store.UpdateUser(ctx, u)
}

func (i *Interactor) GetSCMUserByToken(ctx context.Context, token string) (*ent.User, error) {
	return i.scm.GetUser(ctx, token)
}
