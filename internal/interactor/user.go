package interactor

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/perm"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	// UserInteractor provides application logic for interacting with users.
	UserInteractor struct {
		*service

		admins        []string
		orgEntries    []string
		memberEntries []string
	}

	// UserStore defines operations for working with users.
	UserStore interface {
		CountUsers(context.Context) (int, error)
		SearchUsers(ctx context.Context, opts *SearchUsersOptions) ([]*ent.User, error)
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		FindUserByLogin(ctx context.Context, login string) (*ent.User, error)
		CreateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		DeleteUser(ctx context.Context, u *ent.User) error
	}

	// SearchUsersOptions specifies the optional parameters that
	// search users.
	SearchUsersOptions struct {
		ListOptions

		// Query search the 'login' contains the query.
		Query string
	}

	// UserSCM defines operations for working with remote users.
	UserSCM interface {
		GetRemoteUserByToken(ctx context.Context, token string) (*extent.RemoteUser, error)
		ListRemoteOrgsByToken(ctx context.Context, token string) ([]string, error)
	}
)

// IsAdminUser verifies that the login is an admin or not.
func (i *UserInteractor) IsAdminUser(ctx context.Context, login string) bool {
	for _, admin := range i.admins {
		if login == admin {
			return true
		}
	}

	return false
}

// IsEntryMember verifies that the login is a member or not.
func (i *UserInteractor) IsEntryMember(ctx context.Context, login string) bool {
	if i.memberEntries == nil {
		return true
	}

	for _, m := range i.memberEntries {
		if login == m {
			return true
		}
	}

	return false
}

// IsOrgMember verifies that the user's organizations is in the member entries or not.
func (i *UserInteractor) IsOrgMember(ctx context.Context, orgs []string) bool {
	for _, o := range orgs {
		for _, entry := range i.memberEntries {
			if o == entry {
				return true
			}
		}
	}

	return false
}

// IsEntryOrg verifies that the organization is in the organization entries or not.
func (i *UserInteractor) IsEntryOrg(ctx context.Context, namespace string) bool {
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

// SyncRemoteRepo synchronizes with the remote repository.
func (i *UserInteractor) SyncRemoteRepo(ctx context.Context, u *ent.User, re *extent.RemoteRepo, t time.Time) error {
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
