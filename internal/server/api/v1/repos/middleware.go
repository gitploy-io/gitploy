package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/perm"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"go.uber.org/zap"
)

type (
	RepoMiddleware struct {
		i   Interactor
		log *zap.Logger
	}
)

const (
	KeyRepo = "gitploy.repo"
)

func NewRepoMiddleware(i Interactor) *RepoMiddleware {
	return &RepoMiddleware{
		i:   i,
		log: zap.L().Named("repo-middleware"),
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

		repo, err := rm.i.FindRepoByID(ctx, u, repoID)
		if ent.IsNotFound(err) {
			rm.log.Error("denied to access the repo.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusForbidden, "It has denied to access the repo.")
			return
		} else if err != nil {
			rm.log.Error("failed to get the repository.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
			return
		}

		c.Set(KeyRepo, repo)
	}
}

func (rm *RepoMiddleware) WritePerm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			repoID = c.Param("repoID")
		)

		v, _ := c.Get(gb.KeyUser)
		u := v.(*ent.User)

		ctx := c.Request.Context()

		p, err := rm.i.FindPermByRepoID(ctx, u, repoID)
		if ent.IsNotFound(err) {
			rm.log.Error("denied to access the repo.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusForbidden, "It has denied to access the repo.")
			return
		} else if err != nil {
			rm.log.Error("failed to get the repository.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
			return
		}

		if !(p.RepoPerm == perm.RepoPermWrite || p.RepoPerm == perm.RepoPermAdmin) {
			rm.log.Warn("denied to access the repo.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusForbidden, "It has denied to access the repo, only write permission can access.")
			return
		}
	}
}

func (rm *RepoMiddleware) AdminPerm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			repoID = c.Param("repoID")
		)

		v, _ := c.Get(gb.KeyUser)
		u := v.(*ent.User)

		ctx := c.Request.Context()

		p, err := rm.i.FindPermByRepoID(ctx, u, repoID)
		if ent.IsNotFound(err) {
			rm.log.Error("denied to access the repo.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusForbidden, "It has denied to access the repo.")
			return
		} else if err != nil {
			rm.log.Error("failed to get the repository.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
			return
		}

		if p.RepoPerm != perm.RepoPermAdmin {
			rm.log.Warn("denied to access the repo.", zap.String("repoID", repoID), zap.Error(err))
			gb.ErrorResponse(c, http.StatusForbidden, "It has denied to access the repo, only admin permission can access.")
			return
		}
	}
}
