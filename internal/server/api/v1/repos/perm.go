package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

func (r *Repo) ListPerms(c *gin.Context) {
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

	perms, err := r.i.ListPermsOfRepo(ctx, re, q, page, perPage)
	if err != nil {
		r.log.Error("failed to get permissions.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get permissions.")
		return
	}

	gb.Response(c, http.StatusOK, perms)
}
