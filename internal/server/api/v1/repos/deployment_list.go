package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *DeploymentAPI) List(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		env     = c.Query("env")
		status  = c.Query("status")
		page    int
		perPage int
		err     error
	)
	// Validate quries
	if page, err = strconv.Atoi(c.DefaultQuery("page", defaultQueryPage)); err != nil {
		s.log.Warn("Invalid parameter: page is not integer.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	if perPage, err = strconv.Atoi(c.DefaultQuery("per_page", defaultQueryPage)); err != nil {
		s.log.Warn("Invalid parameter: per_page is not integer.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ds, err := s.i.ListDeploymentsOfRepo(ctx, re, env, status, page, perPage)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	s.log.Debug("Success to list deployments.")
	gb.Response(c, http.StatusOK, ds)
}
