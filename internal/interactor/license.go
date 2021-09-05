package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/pkg/license"
	"github.com/gitploy-io/gitploy/vo"
)

func (i *Interactor) GetLicense(ctx context.Context) (*vo.License, error) {
	var (
		cnt int
		d   *vo.SigningData
		err error
	)

	if cnt, err = i.Store.CountUsers(ctx); err != nil {
		return nil, err
	}

	if i.licenseKey == "" {
		lic := vo.NewTrialLicense(cnt)
		return lic, nil
	}

	if d, err = license.Decode(i.licenseKey); err != nil {
		return nil, err
	}

	lic := vo.NewStandardLicense(cnt, d)
	return lic, nil
}
