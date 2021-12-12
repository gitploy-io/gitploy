// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/pkg/license"
)

func (i *Interactor) GetLicense(ctx context.Context) (*extent.License, error) {
	var (
		memberCnt     int
		deploymentCnt int
		d             *extent.SigningData
		err           error
	)

	if memberCnt, err = i.Store.CountUsers(ctx); err != nil {
		return nil, err
	}

	if deploymentCnt, err = i.Store.CountDeployments(ctx); err != nil {
		return nil, err
	}

	if i.licenseKey == "" {
		lic := extent.NewTrialLicense(memberCnt, deploymentCnt)
		return lic, nil
	}

	if d, err = license.Decode(i.licenseKey); err != nil {
		return nil, e.NewError(
			e.ErrorCodeLicenseDecode,
			err,
		)
	}

	lic := extent.NewStandardLicense(memberCnt, d)
	return lic, nil
}
