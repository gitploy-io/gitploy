package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *DeploymentAPI) Rollback(c *gin.Context) {
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
		s.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	config, err := s.i.GetEvaluatedConfig(ctx, u, re, &extent.EvalValues{IsRollback: true})
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the configuration.").Write(zap.Error(err))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	}

	var env *extent.Env
	if env = config.GetEnv(d.Env); env == nil {
		s.log.Warn("The environment is not found.", zap.String("env", d.Env))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, e.NewError(e.ErrorCodeConfigUndefinedEnv, nil))
		return
	}

	d, err = s.i.Deploy(ctx, u, re,
		&ent.Deployment{
			Type:       d.Type,
			Env:        d.Env,
			Ref:        d.Ref,
			IsRollback: true,
		},
		env)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to deploy.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	// Get the deployment with edges.
	if de, _ := s.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	s.log.Info("Start to rollback.", zap.String("repo", re.GetFullName()), zap.Int("number", d.Number))
	gb.Response(c, http.StatusCreated, d)
}
