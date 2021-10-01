package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/perm"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
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

func (rm *RepoMiddleware) RepoReadPerm() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var (
			namespace = c.Param("namespace")
			name      = c.Param("name")
		)

		v, _ := c.Get(gb.KeyUser)
		u := v.(*ent.User)

		r, err := rm.i.FindRepoOfUserByNamespaceName(ctx, u, namespace, name)
		if ent.IsNotFound(err) {
			rm.log.Warn("The repository is not found.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusNotFound, "The repository is not found.")
			return
		} else if err != nil {
			rm.log.Error("It has failed to get the repository.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
			return
		}

		_, err = rm.i.FindPermOfRepo(ctx, r, u)
		if ent.IsNotFound(err) {
			rm.log.Warn("It is denied to access the repository.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusForbidden, "It is denied to access the repository.")
			return
		} else if err != nil {
			rm.log.Error("It has failed to get the permission.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusInternalServerError, "It has failed to get the permission.")
			return
		}

		c.Set(KeyRepo, r)
	}
}

func (rm *RepoMiddleware) RepoWritePerm() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var (
			namespace = c.Param("namespace")
			name      = c.Param("name")
		)

		v, _ := c.Get(gb.KeyUser)
		u := v.(*ent.User)

		r, err := rm.i.FindRepoOfUserByNamespaceName(ctx, u, namespace, name)
		if ent.IsNotFound(err) {
			rm.log.Warn("The repository is not found.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusNotFound, "The repository is not found.")
			return
		} else if err != nil {
			rm.log.Error("It has failed to get the repository.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
			return
		}

		p, err := rm.i.FindPermOfRepo(ctx, r, u)
		if ent.IsNotFound(err) {
			rm.log.Warn("It is denied to access the repository.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusForbidden, "It is denied to access the repository.")
			return
		} else if err != nil {
			rm.log.Error("It has failed to get the repository.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusInternalServerError, "It has failed to get the permission.")
			return
		}

		if !(p.RepoPerm == perm.RepoPermWrite || p.RepoPerm == perm.RepoPermAdmin) {
			rm.log.Warn("The access is forbidden. Only write permission can access.", zap.String("repo", namespace+"/"+name))
			gb.AbortWithErrorResponse(c, http.StatusForbidden, "Only write permission can access.")
			return
		}

		c.Set(KeyRepo, r)
	}
}

func (rm *RepoMiddleware) RepoAdminPerm() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var (
			namespace = c.Param("namespace")
			name      = c.Param("name")
		)

		v, _ := c.Get(gb.KeyUser)
		u := v.(*ent.User)

		r, err := rm.i.FindRepoOfUserByNamespaceName(ctx, u, namespace, name)
		if ent.IsNotFound(err) {
			rm.log.Warn("The repository is not found.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusNotFound, "The repository is not found.")
			return
		} else if err != nil {
			rm.log.Error("It has failed to get the repository.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
			return
		}

		p, err := rm.i.FindPermOfRepo(ctx, r, u)
		if ent.IsNotFound(err) {
			rm.log.Warn("It is denied to access the repo.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusForbidden, "It is denied to access the repo.")
			return
		} else if err != nil {
			rm.log.Error("It has failed to get the permission.", zap.String("repo", namespace+"/"+name), zap.Error(err))
			gb.AbortWithErrorResponse(c, http.StatusInternalServerError, "It has failed to get the permission.")
			return
		}

		if p.RepoPerm != perm.RepoPermAdmin {
			rm.log.Warn("The access is forbidden. Only admin permission can access.", zap.String("repo", namespace+"/"+name))
			gb.AbortWithErrorResponse(c, http.StatusForbidden, "Only admin permission can access.")
			return
		}

		c.Set(KeyRepo, r)
	}
}
