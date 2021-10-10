//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package middlewares

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	Interactor interface {
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
