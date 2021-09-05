package license

import (
	"context"

	"github.com/gitploy-io/gitploy/vo"
)

type (
	Interactor interface {
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
