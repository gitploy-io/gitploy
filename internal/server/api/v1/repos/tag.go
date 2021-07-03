package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"github.com/hanjunlee/gitploy/vo"
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
		r.log.Error("failed to list tags.", zap.String("repo", repo.Name), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list tags.")
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
	if vo.IsRefNotFoundError(err) {
		r.log.Warn("the tag is not found.", zap.String("repo", repo.Name), zap.String("tag", tag), zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "the tag is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the tag.", zap.String("repo", repo.Name), zap.String("tag", tag), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the tag.")
		return
	}

	gb.Response(c, http.StatusOK, t)
}
