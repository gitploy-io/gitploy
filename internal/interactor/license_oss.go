// +build oss

package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/model/extent"
)

func (i *Interactor) GetLicense(ctx context.Context) (*extent.License, error) {
	return extent.NewOSSLicense(), nil
}
