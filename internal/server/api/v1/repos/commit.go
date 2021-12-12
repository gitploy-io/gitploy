package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/extent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"go.uber.org/zap"
)

func (r *Repo) ListCommits(c *gin.Context) {
	var (
		branch  = c.Query("branch")
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	commits, err := r.i.ListCommits(ctx, u, repo, branch, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to list commits.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, commits)
}

func (r *Repo) GetCommit(c *gin.Context) {
	var (
		sha = c.Param("sha")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	commit, err := r.i.GetCommit(ctx, u, repo, sha)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to get the commit.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, commit)
}

func (r *Repo) ListStatuses(c *gin.Context) {
	var (
		sha = c.Param("sha")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	ss, err := r.i.ListCommitStatuses(ctx, u, repo, sha)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to list commit statuses.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, map[string]interface{}{
		"state":    mergeState(ss),
		"statuses": ss,
	})
}

func mergeState(ss []*extent.Status) string {
	// The state is failure if one of them is failure.
	for _, s := range ss {
		if s.State == extent.StatusStateFailure {
			return string(extent.StatusStateFailure)
		}
	}

	for _, s := range ss {
		if s.State == extent.StatusStatePending {
			return string(extent.StatusStatePending)
		}
	}

	return string(extent.StatusStateSuccess)
}
