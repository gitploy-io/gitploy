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

func (s *DeploymentsAPI) Get(c *gin.Context) {
	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		s.log.Warn("Invalid parameter: number must be integer.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	d, err := s.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, d)
}
