package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

func (r *Repo) ListTags(c *gin.Context) {
	var (
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)
	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	tags, err := r.i.ListTags(ctx, u, repo, atoi(page), atoi(perPage))
	if err != nil {
		gb.LogWithError(r.log, "Failed to list tags.", err)
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, tags)
}

func (r *Repo) GetTag(c *gin.Context) {
	var (
		tag = c.Param("tag")
	)
	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	t, err := r.i.GetTag(ctx, u, repo, tag)
	if err != nil {
		gb.LogWithError(r.log, "Failed to get the tag.", err)
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, t)
}
