// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss

package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *ReviewsAPI) GetMine(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		s.log.Warn("Invalid parameter: number must be integer.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := s.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployment.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	rv, err := s.i.FindReviewOfUser(ctx, u, d)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to find the user's review.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, rv)
}
