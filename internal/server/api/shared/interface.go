//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package shared

import (
	"context"

	"github.com/gitploy-io/gitploy/extent"
)

type (
	Interactor interface {
		GetLicense(ctx context.Context) (*extent.License, error)
	}
)
