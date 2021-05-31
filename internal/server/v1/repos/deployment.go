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
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the configuraton file.")
			return
		}

		r.log.Error("failed to deploy.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to deploy.")
		return
	}

	gb.Response(c, http.StatusCreated, nil)
}

func (r *Repo) Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment) error {
	d, err := r.store.CreateDeployment(ctx, u, re, d)
	if err != nil {
		return fmt.Errorf("failed to create a new deployment on the store: %w", err)
	}

	// Deploy on the strategy
	c, err := r.scm.GetConfig(ctx, u, re)
	if err != nil {
		return err
	}

	if c.HasEnv(d.Env) {
		return fmt.Errorf("The environment is not defined in the configuration file")
	}

	env := c.GetEnv(d.Env)

	return r.deploy(ctx, u, re, d, env)
}

func (r *Repo) deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, e *vo.Env) error {
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

	return fmt.Errorf("Not implemented yet.")
}
