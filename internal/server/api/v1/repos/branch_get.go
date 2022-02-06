package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
)

func (s *BranchesAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		branch = c.Param("branch")
	)

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	b, err := s.i.GetBranch(ctx, u, repo, branch)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the branch.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, b)
}
