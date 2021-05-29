package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

func (r *Repo) ListBranches(c *gin.Context) {
	var (
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)
	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	branches, err := r.scm.ListBranches(ctx, u, repo, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list branches.", zap.String("repo", repo.Name), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list branches.")
		return
	}

	gb.Response(c, http.StatusOK, branches)
}

func (r *Repo) GetBranch(c *gin.Context) {
	var (
		branch = c.Param("branch")
	)
	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	b, err := r.scm.GetBranch(ctx, u, repo, branch)
	if err != nil {
		r.log.Error("failed to get the branch.", zap.String("repo", repo.Name), zap.String("branch", branch), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the branch.")
		return
	}

	gb.Response(c, http.StatusOK, b)
}
