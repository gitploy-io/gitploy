package web

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/vo"
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
		GetRemoteUserByToken(ctx context.Context, token string) (*vo.RemoteUser, error)
		ListRemoteOrgsByToken(ctx context.Context, token string) ([]string, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
