package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"go.uber.org/zap"
)

type (
	RepoMiddleware struct {
		store Store
		log   *zap.Logger
	}
)

const (
	KeyRepo = "gitploy.repo"
)

func NewRepoMiddleware(store Store) *RepoMiddleware {
	return &RepoMiddleware{
		store: store,
		log:   zap.L().Named("repo-middleware"),
	}
}

func (rm *RepoMiddleware) Repo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			repoID = c.Param("repoID")
		)

		v, _ := c.Get(gb.KeyUser)
		u := v.(*ent.User)

		ctx := c.Request.Context()

		repo, err := rm.store.FindRepo(ctx, u, repoID)
		if ent.IsNotFound(err) {
			rm.log.Error("denied to access the repo.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It is denied to access the repo.")
			return

		} else if err != nil {
			rm.log.Error("failed to get the repository.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
			return
		}

		c.Set(KeyRepo, repo)
	}
}
