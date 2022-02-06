package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	DeploymentPostPayload struct {
		Type string `json:"type"`
		Ref  string `json:"ref"`
		Env  string `json:"env"`
	}
)

func (s *DeploymentAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()

	p := &DeploymentPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", nil),
		)
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	config, err := s.i.GetConfig(ctx, u, re)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the configuration.").Write(zap.Error(err))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	}

	if err := config.Eval(&extent.EvalValues{}); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to evaluate the configuration.").Write(zap.Error(err))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	}

	var env *extent.Env
	if env = config.GetEnv(p.Env); env == nil {
		s.log.Warn("The environment is not found.", zap.String("env", p.Env))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, e.NewError(e.ErrorCodeConfigUndefinedEnv, nil))
		return
	}

	d, err := s.i.Deploy(ctx, u, re,
		&ent.Deployment{
			Type: deployment.Type(p.Type),
			Env:  p.Env,
			Ref:  p.Ref,
		},
		env)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to deploy.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if _, err := s.i.CreateEvent(ctx, &ent.Event{
		Kind:         event.KindDeployment,
		Type:         event.TypeCreated,
		DeploymentID: d.ID,
	}); err != nil {
		s.log.Error("It has failed to create the event.", zap.Error(err))
	}

	// Get the deployment with edges.
	if de, _ := s.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	s.log.Info("Success to start to deploy.", zap.String("repo", re.GetFullName()), zap.String("env", p.Env))
	gb.Response(c, http.StatusCreated, d)
}
