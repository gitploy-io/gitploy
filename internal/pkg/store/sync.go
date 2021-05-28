package store

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/user"
)

func (s *Store) FindUserByHash(ctx context.Context, hash string) (*ent.User, error) {
	return s.c.User.
		Query().
		Where(
			user.HashEQ(hash),
		).
		Only(ctx)
}

func (s *Store) SyncPerm(ctx context.Context, u *ent.User, rp *ent.Perm, sync time.Time) error {
	return s.WithTx(ctx, func(tx *ent.Tx) error {
		var (
			remote *ent.Repo
			local  *ent.Repo
			err    error
		)

		// Synchronize remote repositories.
		remote = rp.Edges.Repo
		local, err = tx.Repo.Get(ctx, remote.ID)
		if ent.IsNotFound(err) {
			local, _ = tx.Repo.Create().
				SetID(remote.ID).
				SetNamespace(remote.Namespace).
				SetName(remote.Name).
				SetDescription(remote.Description).
				SetSyncedAt(sync).
				Save(ctx)
		} else {
			local, _ = tx.Repo.UpdateOne(remote).
				SetNamespace(remote.Namespace).
				SetName(remote.Name).
				SetDescription(remote.Description).
				SetSyncedAt(sync).
				Save(ctx)
		}

		// Sync perm,
		// insert a new perm if not exist,
		// update permission if exist.
		// After to synchronize it deletes perms which are not synchronized.
		lp, err := tx.Perm.Query().
			Where(
				perm.And(
					perm.HasUserWith(user.IDEQ(u.ID)),
					perm.HasRepoWith(repo.IDEQ(local.ID)),
				),
			).
			Only(ctx)
		if ent.IsNotFound(err) {
			tx.Perm.Create().
				SetUserID(u.ID).
				SetRepoID(local.ID).
				SetRepoPerm(rp.RepoPerm).
				SetSyncedAt(sync).
				Save(ctx)
		} else {
			tx.Perm.UpdateOne(lp).
				SetRepoPerm(rp.RepoPerm).
				SetSyncedAt(sync).
				Save(ctx)
		}

		_, err = tx.Perm.
			Delete().
			Where(
				perm.HasUserWith(user.IDEQ(u.ID)),
				perm.SyncedAtNEQ(sync),
			).
			Exec(ctx)
		return err
	})
}
