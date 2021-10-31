// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/approval"
	"github.com/gitploy-io/gitploy/ent/event"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	approvalPostPayload struct {
		UserID int64 `json:"user_id"`
	}

	approvalPatchPayload struct {
		Status string `json:"status"`
	}
)

func (r *Repo) ListApprovals(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployment.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	as, err := r.i.ListApprovals(ctx, d)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to list approvals.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, as)
}

func (r *Repo) GetApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		aid = c.Param("aid")
	)

	ap, err := r.i.FindApprovalByID(ctx, atoi(aid))
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the approval.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, ap)
}

func (r *Repo) GetMyApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "The number must be number.", nil),
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

	a, err := r.i.FindApprovalOfUser(ctx, d, u)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the user's approval.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, a)
}

func (r *Repo) CreateApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "The number must be number.", nil),
		)
		return
	}

	p := &approvalPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Warn("failed to bind the payload.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "It has failed to bind the playload.", nil),
		)
		return
	}

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployment.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	// TODO: Migrate the business logic into the interactor.
	user, err := r.i.FindUserByID(ctx, p.UserID)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the user.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	_, err = r.i.FindPermOfRepo(ctx, re, user)
	if e.HasErrorCode(err, e.ErrorCodeNotFound) {
		r.log.Warn("The approver has no permission for the repository.", zap.Error(err))
		// Override the HTTP status.
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	} else if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the perm of approver.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if d.Edges.User != nil && user.ID == d.Edges.User.ID {
		r.log.Warn("Failed to create a new approval.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeUnprocessableEntity, "The deployer can not be the approver.", nil),
		)
		return
	}

	ap, err := r.i.CreateApproval(ctx, &ent.Approval{
		UserID:       user.ID,
		DeploymentID: d.ID,
	})
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to create a new approval.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if _, err := r.i.CreateEvent(ctx, &ent.Event{
		Kind:       event.KindApproval,
		Type:       event.TypeCreated,
		ApprovalID: ap.ID,
	}); err != nil {
		r.log.Error("Failed to create the event.", zap.Error(err))
	}

	// Get the approval with edges
	if ae, _ := r.i.FindApprovalByID(ctx, ap.ID); ae != nil {
		ap = ae
	}

	gb.Response(c, http.StatusCreated, ap)
}

func (r *Repo) UpdateMyApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "The number must be number.", nil),
		)
		return
	}

	p := &approvalPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Warn("failed to bind the payload.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "It has failed to bind the payload.", nil),
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

	a, err := r.i.FindApprovalOfUser(ctx, d, u)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the user's approval.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if p.Status != string(a.Status) {
		a.Status = approval.Status(p.Status)
		if a, err = r.i.UpdateApproval(ctx, a); err != nil {
			r.log.Check(gb.GetZapLogLevel(err), "Failed to update the approval.").Write(zap.Error(err))
			gb.ResponseWithError(c, err)
			return
		}

		if _, err := r.i.CreateEvent(ctx, &ent.Event{
			Kind:       event.KindApproval,
			Type:       event.TypeUpdated,
			ApprovalID: a.ID,
		}); err != nil {
			r.log.Error("Failed to create the event.", zap.Error(err))
		}
	}

	// Get the approval with edges
	if ae, _ := r.i.FindApprovalOfUser(ctx, d, u); ae != nil {
		a = ae
	}

	gb.Response(c, http.StatusOK, a)
}

func (r *Repo) DeleteApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		aid int
		err error
	)

	if aid, err = strconv.Atoi(c.Param("aid")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "The number must be number.", nil),
		)
		return
	}

	ap, err := r.i.FindApprovalByID(ctx, aid)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the approval.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if err := r.i.DeleteApproval(ctx, ap); err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to delete the approval.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if _, err := r.i.CreateEvent(ctx, &ent.Event{
		Kind:      event.KindApproval,
		Type:      event.TypeDeleted,
		DeletedID: aid,
	}); err != nil {
		r.log.Error("It has failed to create a new event.", zap.Error(err))
	}

	c.Status(http.StatusOK)
}
