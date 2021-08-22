package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/pkg/license"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) GetLicense(ctx context.Context) (*vo.License, error) {
	if i.license != nil {
		return i.license, nil
	}

	var (
		cnt int
		d   *vo.SigningData
		err error
	)

	if cnt, err = i.Store.CountUsers(ctx); err != nil {
		return nil, err
	}

	if i.licenseKey == "" {
		i.license = vo.NewTrialLicense(cnt)
		return i.license, nil
	}

	if d, err = license.Decode(i.licenseKey); err != nil {
		return nil, err
	}

	i.license = vo.NewStandardLicense(cnt, d)
	return i.license, nil
}
