package repos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/event"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	deploymentPostPayload struct {
		Type string `json:"type"`
		Ref  string `json:"ref"`
		Env  string `json:"env"`
	}

	deploymentPatchPayload struct {
		Status string `json:"status"`
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

	ds, err := r.i.ListDeploymentsOfRepo(ctx, re, env, status, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list deployments.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list deployments.")
		return
	}

	gb.Response(c, http.StatusOK, ds)
}

func (r *Repo) GetDeploymentByNumber(c *gin.Context) {
	var (
		number = c.Param("number")
	)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("the deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
	} else if err != nil {
		r.log.Error("failed to find deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "It has failed to find the deployment.")
	}

	gb.Response(c, http.StatusOK, d)
}

func (r *Repo) CreateDeployment(c *gin.Context) {
	ctx := c.Request.Context()

	p := &deploymentPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	cf, err := r.i.GetConfig(ctx, u, re)
	if e.HasErrorCode(err, e.ErrorCodeConfigNotFound) {
		r.log.Error("The configuration file is not found.", zap.Error(err))
		// To override the HTTP status 422.
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	} else if err != nil {
		r.log.Error("It has failed to get the configuration.")
		gb.ResponseWithError(c, err)
		return
	}

	var env *vo.Env
	if env = cf.GetEnv(p.Env); env == nil {
		r.log.Warn("The environment is not defined in the configuration.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The environment is not defined in the configuration.")
		return
	}

	if err := env.Eval(&vo.EvalValues{}); err != nil {
		r.log.Warn("It has failed to eval variables in the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to eval variables in the config.")
		return
	}

	d, err := r.i.Deploy(ctx, u, re,
		&ent.Deployment{
			Type: deployment.Type(p.Type),
			Env:  p.Env,
			Ref:  p.Ref,
		},
		cf.GetEnv(p.Env),
	)
	if err != nil {
		r.log.Error("It has failed to deploy.", zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if _, err := r.i.CreateEvent(ctx, &ent.Event{
		Kind:         event.KindDeployment,
		Type:         event.TypeCreated,
		DeploymentID: d.ID,
	}); err != nil {
		r.log.Error("It has failed to create the event.", zap.Error(err))
	}

	// Get the deployment with edges.
	if de, _ := r.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	gb.Response(c, http.StatusCreated, d)
}

// UpdateDeployment creates a new remote deployment and
// patch the deployment status 'created'.
func (r *Repo) UpdateDeployment(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	p := &deploymentPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("the deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
	}

	cf, err := r.i.GetConfig(ctx, u, re)
	if e.HasErrorCode(err, e.ErrorCodeConfigNotFound) {
		r.log.Error("The configuration file is not found.", zap.Error(err))
		// To override the HTTP status 422.
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	} else if err != nil {
		r.log.Error("It has failed to get the configuration.")
		gb.ResponseWithError(c, err)
		return
	}

	var env *vo.Env
	if env = cf.GetEnv(d.Env); env == nil {
		r.log.Warn("The environment is not defined in the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The environment is not defined in the config.")
		return
	}

	if err := env.Eval(&vo.EvalValues{IsRollback: d.IsRollback}); err != nil {
		r.log.Warn("It has failed to eval variables in the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to eval variables in the config.")
		return
	}

	if p.Status == string(deployment.StatusCreated) && d.Status == deployment.StatusWaiting {
		if d, err = r.i.DeployToRemote(ctx, u, re, d, env); err != nil {
			r.log.Error("It has failed to deploy to the remote.", zap.Error(err))
			gb.ResponseWithError(c, err)
			return
		}

		if _, err := r.i.CreateEvent(ctx, &ent.Event{
			Kind:         event.KindDeployment,
			Type:         event.TypeUpdated,
			DeploymentID: d.ID,
		}); err != nil {
			r.log.Error("It has failed to create an event.", zap.Error(err))
		}
	}

	// Get the deployment with edges.
	if de, _ := r.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	gb.Response(c, http.StatusOK, d)
}

func (r *Repo) RollbackDeployment(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("the deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	}

	cf, err := r.i.GetConfig(ctx, u, re)
	if e.HasErrorCode(err, e.ErrorCodeConfigNotFound) {
		r.log.Error("The configuration file is not found.", zap.Error(err))
		// To override the HTTP status 422.
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	} else if err != nil {
		r.log.Error("It has failed to get the configuration.")
		gb.ResponseWithError(c, err)
		return
	}

	var env *vo.Env
	if env = cf.GetEnv(d.Env); env == nil {
		r.log.Warn("The environment is not defined in the configuration.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The environment is not defined in the configuration.")
		return
	}

	if err := env.Eval(&vo.EvalValues{IsRollback: true}); err != nil {
		r.log.Warn("It has failed to eval variables in the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to eval variables in the config.")
		return
	}

	d, err = r.i.Deploy(ctx, u, re,
		&ent.Deployment{
			Type:       d.Type,
			Env:        d.Env,
			Ref:        d.Ref,
			IsRollback: true,
		},
		cf.GetEnv(d.Env),
	)
	if err != nil {
		r.log.Error("It has failed to deploy.", zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if _, err := r.i.CreateEvent(ctx, &ent.Event{
		Kind:         event.KindDeployment,
		Type:         event.TypeCreated,
		DeploymentID: d.ID,
	}); err != nil {
		r.log.Error("It has failed to create the event.", zap.Error(err))
	}

	// Get the deployment with edges.
	if de, _ := r.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	gb.Response(c, http.StatusCreated, d)
}

func (r *Repo) ListDeploymentChanges(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number  = c.Param("number")
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("The deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("It has failed to find the deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the deployment.")
		return
	}

	ld, err := r.i.FindPrevSuccessDeployment(ctx, d)
	if ent.IsNotFound(err) {
		gb.Response(c, http.StatusOK, []*vo.Commit{})
		return
	} else if err != nil {
		r.log.Error("It has failed to find the comparable deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the comparable deployment.")
		return
	}

	// Get SHA when the status of deployment is waiting.
	sha := d.Sha
	if sha == "" {
		sha, err = r.getCommitSha(ctx, u, re, d.Type, d.Ref)
		if err != nil {
			r.log.Error("It has failed to get the commit SHA.", zap.Error(err))
			gb.ResponseWithError(c, err)
			return
		}
	}

	commits, _, err := r.i.CompareCommits(ctx, u, re, ld.Sha, sha, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("It has failed to compare two commits.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to compare two commits.")
		return
	}

	gb.Response(c, http.StatusOK, commits)
}

func (r *Repo) getCommitSha(ctx context.Context, u *ent.User, re *ent.Repo, typ deployment.Type, ref string) (string, error) {
	switch typ {
	case deployment.TypeCommit:
		c, err := r.i.GetCommit(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return c.SHA, nil
	case deployment.TypeBranch:
		b, err := r.i.GetBranch(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return b.CommitSHA, nil
	case deployment.TypeTag:
		t, err := r.i.GetTag(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return t.CommitSHA, nil
	default:
		return "", fmt.Errorf("Type must be one of commit, branch, tag.")
	}
}

func (r *Repo) GetConfig(c *gin.Context) {
	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	config, err := r.i.GetConfig(ctx, u, re)
	if err != nil {
		r.log.Error("It has failed to get the configuration.", zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, config)
}
