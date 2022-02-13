package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
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
)

func (i *UserInteractor) IsAdminUser(ctx context.Context, login string) bool {
	for _, admin := range i.admins {
		if login == admin {
			return true
		}
	}

	return false
}

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
