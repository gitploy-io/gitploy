package repos

import (
	"context"
	"fmt"
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
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)
	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	ds, err := r.store.ListDeployments(ctx, re, env, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list deployments.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list deployments.")
		return
	}

	gb.Response(c, http.StatusOK, ds)
}

func (r *Repo) GetLatestDeployment(c *gin.Context) {
	var (
		env = c.Query("env")
	)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	if env == "" {
		r.log.Warn("\"env\" query is required.")
		gb.ErrorResponse(c, http.StatusBadRequest, "\"env\" is not exist in the query.")
		return
	}

	ctx := c.Request.Context()

	d, err := r.store.FindLatestDeployment(ctx, re, env)
	if ent.IsNotFound(err) {
		r.log.Warn("the latest deployment is not found.", zap.String("repo", re.Name), zap.String("env", env))
		gb.ErrorResponse(c, http.StatusNotFound, "the latest deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the latest deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the latest deployment.")
		return
	}

	gb.Response(c, http.StatusOK, d)
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

	d := &ent.Deployment{
		Type: deployment.Type(p.Type),
		Ref:  p.Ref,
		Env:  p.Env,
	}
	if err := r.Deploy(ctx, u, re, d); err != nil {
		if IsConfigNotFoundError(err) {
			r.log.Warn("failed to get the config.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to find the configuraton file.")
			return
		} else if IsConfigParseError(err) {
			r.log.Warn("failed to parse the config.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to parse the configuraton file.")
			return
		} else if IsEnvNotFoundError(err) {
			r.log.Warn("failed to get the env.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to get the env in the configuration file.")
			return
		}

		r.log.Error("failed to deploy.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to deploy.")
		return
	}

	gb.Response(c, http.StatusCreated, nil)
}

func (r *Repo) Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment) error {
	c, err := r.scm.GetConfig(ctx, u, re)
	if err != nil {
		return err
	}

	if !c.HasEnv(d.Env) {
		return &EnvNotFoundError{
			RepoName: re.Name,
		}
	}

	env := c.GetEnv(d.Env)

	d, err = r.store.CreateDeployment(ctx, u, re, d)
	if err != nil {
		return fmt.Errorf("failed to create a new deployment on the store: %w", err)
	}

	return r.deployToSCM(ctx, u, re, d, env)
}

func (r *Repo) deployToSCM(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, e *vo.Env) error {
	if !e.HasApproval() {
		d, err := r.scm.CreateDeployment(ctx, u, re, d, e)
		if err != nil {
			d.Status = deployment.StatusFailure
			r.store.UpdateDeployment(ctx, d)
			return nil
		}

		d.Status = deployment.StatusCreated
		r.store.UpdateDeployment(ctx, d)
		return nil
	}

	// TODO: handling approval.
	return fmt.Errorf("Not implemented yet.")
}
