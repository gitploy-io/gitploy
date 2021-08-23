package web

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id string) (*ent.User, error)
		CreateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		// Fetch the user information from SCM.
		// It has the id, login, avatar and so on.
		GetRemoteUserByToken(ctx context.Context, token string) (*vo.RemoteUser, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
