package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"go.uber.org/zap"
)

func (r *Repo) ListCommits(c *gin.Context) {
	var (
		repoID  = c.Param("repoID")
		branch  = c.Query("branch")
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	commits, err := r.scm.ListCommits(ctx, u, repo, branch, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list commits.", zap.String("repoID", repoID), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list commits.")
		return
	}

	gb.Response(c, http.StatusOK, commits)
}

func (r *Repo) GetCommit(c *gin.Context) {
	var (
		repoID = c.Param("repoID")
		sha    = c.Param("sha")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	commit, err := r.scm.GetCommit(ctx, u, repo, sha)
	if err != nil {
		r.log.Error("failed to list commits.", zap.String("repoID", repoID), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list commits.")
		return
	}

	gb.Response(c, http.StatusOK, commit)
}
