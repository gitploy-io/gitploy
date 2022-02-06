package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"go.uber.org/zap"
)

func (s *CommitService) Get(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		sha = c.Param("sha")
	)

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	commit, err := s.i.GetCommit(ctx, u, repo, sha)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the commit.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, commit)
}
