// +build oss

package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/vo"
)

func (i *Interactor) GetLicense(ctx context.Context) (*vo.License, error) {
	return vo.NewOSSLicense(), nil
}
