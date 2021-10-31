package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

func (r *Repo) ListBranches(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	branches, err := r.i.ListBranches(ctx, u, repo, atoi(page), atoi(perPage))
	if err != nil {
		gb.LogWithError(r.log, "Failed to list branches.", err)
		gb.ResponseWithError(c, err)
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

	b, err := r.i.GetBranch(ctx, u, repo, branch)
	if err != nil {
		gb.LogWithError(r.log, "Failed to get the branch.", err)
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, b)
}
