package license

import (
	"context"

	"github.com/gitploy-io/gitploy/extent"
)

type (
	Interactor interface {
		GetLicense(ctx context.Context) (*extent.License, error)
	}
)
