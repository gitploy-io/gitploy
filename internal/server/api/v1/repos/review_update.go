// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss

package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/review"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	ReviewPatchPayload struct {
		Status  string  `json:"status"`
		Comment *string `json:"comment"`
	}
)

func (s *ReviewAPI) UpdateMine(c *gin.Context) {
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

	p := &ReviewPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		s.log.Warn("Failed to bind the payload.", zap.Error(err))
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", nil))
		return
	}
	if err := review.StatusValidator(review.Status(p.Status)); err != nil {
		s.log.Warn("The status is invalid.", zap.Error(err))
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The status is invalid.", nil))
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

	rv.Status = review.Status(p.Status)

	if p.Comment != nil {
		rv.Comment = *p.Comment
	}

	if rv, err = s.i.RespondReview(ctx, rv); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to update the review.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	s.log.Info("Respond the review successfully.", zap.Int("review_id", rv.ID))
	gb.Response(c, http.StatusOK, rv)
}
