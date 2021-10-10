//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package shared

import (
	"context"

	"github.com/gitploy-io/gitploy/vo"
)

type (
	Interactor interface {
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
