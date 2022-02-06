package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

func (s *CommitAPI) ListStatuses(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		sha = c.Param("sha")
	)

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	ss, err := s.i.ListCommitStatuses(ctx, u, repo, sha)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list commit statuses.").Write(zap.Error(err))
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
		if s.State == extent.StatusStateFailure || s.State == extent.StatusStateCancelled {
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
