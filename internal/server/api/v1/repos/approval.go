package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	approvalPayload struct {
		IsApproved bool `json:"is_approved"`
	}
)

func (r *Repo) ListApprovals(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentWithEdgesOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("The deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the deployment.")
		return
	}

	as, err := r.i.ListApprovals(ctx, d)
	if err != nil {
		r.log.Error("failed to list approvals.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list approvals.")
		return
	}

	gb.Response(c, http.StatusOK, as)
}

func (r *Repo) GetApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentWithEdgesOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("The deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the deployment.")
		return
	}

	a, err := r.i.GetApprovalOfUser(ctx, d, u)
	if ent.IsNotFound(err) {
		r.log.Warn("The approval is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The approval is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the approval.")
		return
	}

	gb.Response(c, http.StatusOK, a)
}

func (r *Repo) UpdateApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	p := &approvalPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Warn("failed to bind the payload.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentWithEdgesOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("The deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the deployment.")
		return
	}

	a, err := r.i.GetApprovalOfUser(ctx, d, u)
	if ent.IsNotFound(err) {
		r.log.Warn("The approval is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The approval is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the approval.")
		return
	}

	if p.IsApproved != a.IsApproved {
		if a, err = r.i.UpdateApproval(ctx, a); err != nil {
			r.log.Error("failed to update the approval.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to update the approval.")
			return
		}
	}

	gb.Response(c, http.StatusOK, a)
}
