// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/pkg/license"
	"github.com/gitploy-io/gitploy/vo"
)

func (i *Interactor) GetLicense(ctx context.Context) (*vo.License, error) {
	var (
		memberCnt     int
		deploymentCnt int
		d             *vo.SigningData
		err           error
	)

	if cnt, err = i.Store.CountUsers(ctx); err != nil {
		return nil, e.NewError(
			e.ErrorCodeInternalError,
			err,
		)

	if deploymentCnt, err = i.Store.CountDeployments(ctx); err != nil {
		return nil, err
	}

	if i.licenseKey == "" {
		lic := vo.NewTrialLicense(memberCnt, deploymentCnt)
		return lic, nil
	}

	if d, err = license.Decode(i.licenseKey); err != nil {
		return nil, e.NewError(
			e.ErrorCodeLicenseDecode,
			err,
		)
	}

	lic := vo.NewStandardLicense(memberCnt, d)
	return lic, nil
}
