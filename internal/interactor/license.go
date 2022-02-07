// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss

package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/pkg/license"
)

func (i *LicenseInteractor) GetLicense(ctx context.Context) (*extent.License, error) {
	var (
		memberCnt     int
		deploymentCnt int
		d             *extent.SigningData
		err           error
	)

	if memberCnt, err = i.store.CountUsers(ctx); err != nil {
		return nil, err
	}

	if deploymentCnt, err = i.store.CountDeployments(ctx); err != nil {
		return nil, err
	}

	if i.LicenseKey == "" {
		lic := extent.NewTrialLicense(memberCnt, deploymentCnt)
		return lic, nil
	}

	if d, err = license.Decode(i.LicenseKey); err != nil {
		return nil, e.NewError(
			e.ErrorCodeLicenseDecode,
			err,
		)
	}

	lic := extent.NewStandardLicense(memberCnt, d)
	return lic, nil
}
