package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
)

func (s *ConfigService) Get(c *gin.Context) {
	ctx := c.Request.Context()

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	config, err := s.i.GetConfig(ctx, u, re)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the configuration.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, config)
}
