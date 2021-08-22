package license

import (
	"context"

	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
