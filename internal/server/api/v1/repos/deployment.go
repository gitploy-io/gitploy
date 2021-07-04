package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	deploymentPayload struct {
		Type string `json:"type"`
		Ref  string `json:"ref"`
		Env  string `json:"env"`
	}
)

func (r *Repo) ListDeployments(c *gin.Context) {
	var (
		env     = c.Query("env")
		status  = c.Query("status")
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)
	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	ds, err := r.i.ListDeployments(ctx, re, env, status, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list deployments.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list deployments.")
		return
	}

	gb.Response(c, http.StatusOK, ds)
}

func (r *Repo) GetConfig(c *gin.Context) {
	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	config, err := r.i.GetConfig(ctx, u, re)
	if vo.IsConfigNotFoundError(err) {
		r.log.Warn("failed to find the config file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "It has failed to find the configuraton file.")
		return
	} else if vo.IsConfigParseError(err) {
		r.log.Warn("failed to parse the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to parse the configuraton file.")
		return
	} else if err != nil {
		r.log.Error("failed to get the config file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the config file.")
		return
	}

	gb.Response(c, http.StatusOK, config)
}

func (r *Repo) CreateDeployment(c *gin.Context) {
	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	p := &deploymentPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	ctx := c.Request.Context()

	cf, err := r.i.GetConfig(ctx, u, re)
	if vo.IsConfigNotFoundError(err) {
		r.log.Warn("failed to get the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to find the configuraton file.")
		return
	} else if vo.IsConfigParseError(err) {
		r.log.Warn("failed to parse the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to parse the configuraton file.")
		return
	} else if err != nil {
		r.log.Error("failed to get the configuration file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the configuraton file.")
		return
	}

	if !cf.HasEnv(p.Env) {
		r.log.Warn("failed to get the env.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to get the env in the configuration file.")
		return
	}

	d, err := r.i.Deploy(ctx, u, re, &ent.Deployment{
		Type: deployment.Type(p.Type),
		Ref:  p.Ref,
		Env:  p.Env,
	}, cf.GetEnv(p.Env))
	if ent.IsConstraintError(err) {
		r.log.Warn("deployment number conflict.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusConflict, "The conflict occurs, please retry.")
		return
	} else if err != nil {
		r.log.Error("failed to deploy.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to deploy.")
		return
	}

	if err = r.i.Publish(ctx, d); err != nil {
		r.log.Warn("failed to notify the deployment.", zap.Error(err))
	}

	gb.Response(c, http.StatusCreated, d)
}
