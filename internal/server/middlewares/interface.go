//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package middlewares

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
