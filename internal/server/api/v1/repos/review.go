// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss
// +build !oss

package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/model/ent/review"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	ReviewService service

	reviewPatchPayload struct {
		Status  string  `json:"status"`
		Comment *string `json:"comment"`
	}
)

func (r *Repo) UpdateUserReview(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		r.log.Warn("The number of deployment must be number.")
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The number of deployment must be number.", err),
		)
		return
	}

	p := &reviewPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Warn("Failed to bind the payload.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", nil),
		)
		return
	}
	if err := review.StatusValidator(review.Status(p.Status)); err != nil {
		r.log.Warn("The status is invalid.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The status is invalid.", nil),
		)
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployment.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	rv, err := r.i.FindReviewOfUser(ctx, u, d)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the user's review.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	rv.Status = review.Status(p.Status)

	if p.Comment != nil {
		rv.Comment = *p.Comment
	}

	if rv, err = r.i.UpdateReview(ctx, rv); err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to update the review.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if _, err := r.i.CreateEvent(ctx, &ent.Event{
		Kind:     event.KindReview,
		Type:     event.TypeUpdated,
		ReviewID: rv.ID,
	}); err != nil {
		r.log.Error("Failed to create the event.", zap.Error(err))
	}

	gb.Response(c, http.StatusOK, rv)
}
