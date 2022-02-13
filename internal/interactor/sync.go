package interactor

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/perm"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	SyncInteractor struct {
		*service

		orgEntries []string
	}

	SyncSCM interface {
		GetRemoteUserByToken(ctx context.Context, token string) (*extent.RemoteUser, error)
		ListRemoteOrgsByToken(ctx context.Context, token string) ([]string, error)
	}
)

func (i *SyncInteractor) IsEntryOrg(ctx context.Context, namespace string) bool {
	if i.orgEntries == nil {
		return true
	}

	for _, r := range i.orgEntries {
		if namespace == r {
			return true
		}
	}

	return false
}

func (i *SyncInteractor) SyncRemoteRepo(ctx context.Context, u *ent.User, re *extent.RemoteRepo, t time.Time) error {
	var (
		r   *ent.Repo
		p   *ent.Perm
		err error
	)

	if r, err = i.store.FindRepoByID(ctx, re.ID); ent.IsNotFound(err) {
		if r, err = i.store.SyncRepo(ctx, re); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if p, err = i.store.FindPermOfRepo(ctx, r, u); ent.IsNotFound(err) {
		if _, err = i.store.CreatePerm(ctx, &ent.Perm{
			RepoPerm: perm.RepoPerm(re.Perm),
			UserID:   u.ID,
			RepoID:   r.ID,
			SyncedAt: t,
		}); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		p.RepoPerm = perm.RepoPerm(re.Perm)
		p.SyncedAt = t

		if _, err = i.store.UpdatePerm(ctx, p); err != nil {
			return err
		}
	}

	return nil
}
