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

func (s *DeploymentStatusAPI) CreateRemote(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		s.log.Warn("Invalid parameter: number is not integer.", zap.String("number", c.Param("number")))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	// Bind the request body.
	p := &DeploymentStatusPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		s.log.Error("Failed to bind the body.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := s.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}
	rds := &extent.RemoteDeploymentStatus{
		Status:      p.Status,
		Description: p.Description,
		LogURL:      p.LogURL,
	}
	if rds, err = s.i.CreateRemoteDeploymentStatus(ctx, u, re, d, rds); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to create a new deployment status.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	s.log.Debug("Success to create a remote deployment status.")
	gb.Response(c, http.StatusCreated, rds)
}
