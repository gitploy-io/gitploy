package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) Sync(ctx context.Context, u *ent.User) error {
	perms, err := i.GetAllPermsWithRepo(ctx, u.Token)
	if err != nil {
		return fmt.Errorf("failed to get all permissions: %w", err)
	}
	i.log.Debug("get all permissions.")

	sync := time.Now()

	for _, perm := range perms {
		re := perm.Edges.Repo
		if err := i.SyncPerm(ctx, u, perm, sync); err != nil {
			return fmt.Errorf("failed to sync with the \"%s\" repo: %w", re.Name, err)
		}
	}

	return nil
}

func (i *Interactor) SyncRemoteRepo(ctx context.Context, u *ent.User, re *vo.RemoteRepo) error {
	var (
		r   *ent.Repo
		p   *ent.Perm
		err error
	)

	if r, err = i.Store.FindRepoByID(ctx, re.ID); ent.IsNotFound(err) {
		if r, err = i.Store.SyncRepo(ctx, re); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if p, err = i.Store.FindPermOfRepo(ctx, r, u); ent.IsNotFound(err) {
		if _, err = i.Store.CreatePerm(ctx, &ent.Perm{
			RepoPerm: perm.RepoPerm(re.Perm),
			UserID:   u.ID,
			RepoID:   r.ID,
		}); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		p.RepoPerm = perm.RepoPerm(re.Perm)

		if _, err = i.Store.UpdatePerm(ctx, p); err != nil {
			return err
		}
	}

	return nil
}
