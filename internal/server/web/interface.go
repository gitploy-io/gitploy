package web

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)
		IsAdminUser(ctx context.Context, login string) bool
		IsEntryMember(ctx context.Context, login string) bool
		IsOrgMember(ctx context.Context, orgs []string) bool
		CreateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		// Fetch the user information from SCM.
		// It has the id, login, avatar and so on.
		GetRemoteUserByToken(ctx context.Context, token string) (*extent.RemoteUser, error)
		ListRemoteOrgsByToken(ctx context.Context, token string) ([]string, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		GetConfigRedirectURL(ctx context.Context, u *ent.User, r *ent.Repo) (string, error)
		GetNewFileRedirectURL(ctx context.Context, u *ent.User, r *ent.Repo) (string, error)
		GetLicense(ctx context.Context) (*extent.License, error)
	}
)
