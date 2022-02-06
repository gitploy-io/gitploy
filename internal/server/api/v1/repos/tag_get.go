package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"go.uber.org/zap"
)

func (s *TagService) Get(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		tag = c.Param("tag")
	)

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	t, err := s.i.GetTag(ctx, u, repo, tag)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the tag.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, t)
}
