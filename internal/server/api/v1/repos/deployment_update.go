package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"go.uber.org/zap"
)

// UpdateDeployment trigger to request for a waiting deployment, and
// it patches the deployment status 'created'.
func (s *DeploymentsAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := s.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	config, err := s.i.GetConfig(ctx, u, re)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the configuration.").Write(zap.Error(err))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	}

	if err := config.Eval(&extent.EvalValues{IsRollback: d.IsRollback}); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to evaludate the configuration.").Write(zap.Error(err))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	}

	var env *extent.Env
	if env = config.GetEnv(d.Env); env == nil {
		s.log.Warn("The environment is not found.", zap.String("env", d.Env))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, e.NewError(e.ErrorCodeConfigUndefinedEnv, nil))
		return
	}

	if d, err = s.i.DeployToRemote(ctx, u, re, d, env); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "It has failed to deploy to the remote.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if _, err := s.i.CreateEvent(ctx, &ent.Event{
		Kind:         event.KindDeployment,
		Type:         event.TypeUpdated,
		DeploymentID: d.ID,
	}); err != nil {
		s.log.Error("It has failed to create an event.", zap.Error(err))
	}

	// Get the deployment with edges.
	if de, _ := s.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	s.log.Info("Trigger to create a remote deployment.", zap.String("repo", re.GetFullName()), zap.Int("number", d.Number))
	gb.Response(c, http.StatusOK, d)
}
