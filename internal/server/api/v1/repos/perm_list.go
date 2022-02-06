package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
)

func (s *PermsAPI) List(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		q       = c.DefaultQuery("q", "")
		page    = atoi(c.DefaultQuery("page", "1"))
		perPage = atoi(c.DefaultQuery("per_page", "30"))
	)

	v, _ := c.Get(KeyRepo)
	re := v.(*ent.Repo)

	if perPage > 100 {
		perPage = 100
	}

	perms, err := s.i.ListPermsOfRepo(ctx, re, q, page, perPage)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list permissions.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, perms)
}
