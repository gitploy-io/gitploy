package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

func (u *UserAPI) GetRateLimit(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	uv, _ := v.(*ent.User)

	var (
		rl  *extent.RateLimit
		err error
	)

	if rl, err = u.i.GetRateLimit(ctx, uv); err != nil {
		u.log.Check(gb.GetZapLogLevel(err), "Failed to get the rate-limit.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, rl)
}
