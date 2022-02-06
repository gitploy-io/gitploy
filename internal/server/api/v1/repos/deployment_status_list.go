package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"go.uber.org/zap"
)

func (s *DeploymentStatusesAPI) List(c *gin.Context) {
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

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := s.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	dss, err := s.i.ListDeploymentStatuses(ctx, d)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list the deployment statuses.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	s.log.Debug("Success to list deployment statuses.")
	gb.Response(c, http.StatusOK, dss)
}
