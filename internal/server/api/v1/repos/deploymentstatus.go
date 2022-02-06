package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	DeploymentStatusPostPayload struct {
		Status      string `json:"status"`
		Description string `json:"description"`
		LogURL      string `json:"log_url"`
	}
)

func (r *Repo) ListDeploymentStatuses(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		r.log.Warn("Invalid parameter: number is not integer.", zap.String("number", c.Param("number")))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to get the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	dss, err := r.i.ListDeploymentStatuses(ctx, d)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to list the deployment statuses.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	r.log.Debug("Success to list deployment statuses.")
	gb.Response(c, http.StatusOK, dss)
}

func (r *Repo) CreateRemoteDeploymentStatus(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		r.log.Warn("Invalid parameter: number is not integer.", zap.String("number", c.Param("number")))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	// Bind the request body.
	p := &DeploymentStatusPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Error("Failed to bind the body.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to get the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}
	rds := &extent.RemoteDeploymentStatus{
		Status:      p.Status,
		Description: p.Description,
		LogURL:      p.LogURL,
	}
	if rds, err = r.i.CreateRemoteDeploymentStatus(ctx, u, re, d, rds); err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to create a new deployment status.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	r.log.Debug("Success to create a remote deployment status.")
	gb.Response(c, http.StatusCreated, rds)
}
