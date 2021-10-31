package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/perm"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
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
		if err != nil {
			rm.log.Check(gb.GetZapLogLevel(err), "Failed to find the repository.").Write(zap.Error(err))
			gb.AbortWithError(c, err)
			return
		}

		_, err = rm.i.FindPermOfRepo(ctx, r, u)
		if e.HasErrorCode(err, e.ErrorCodeNotFound) {
			rm.log.Check(gb.GetZapLogLevel(err), "It is denied to acess the repository.").Write(zap.Error(err))
			gb.AbortWithStatusAndError(c, http.StatusForbidden, err)
			return
		} else if err != nil {
			rm.log.Check(gb.GetZapLogLevel(err), "Failed to find the permission.").Write(zap.Error(err))
			gb.AbortWithError(c, err)
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
		if err != nil {
			rm.log.Check(gb.GetZapLogLevel(err), "Failed to find the repository.").Write(zap.Error(err))
			gb.AbortWithError(c, err)
			return
		}

		p, err := rm.i.FindPermOfRepo(ctx, r, u)
		if e.HasErrorCode(err, e.ErrorCodeNotFound) {
			rm.log.Check(gb.GetZapLogLevel(err), "It is denied to acess the repository.").Write(zap.Error(err))
			gb.AbortWithStatusAndError(c, http.StatusForbidden, err)
			return
		} else if err != nil {
			rm.log.Check(gb.GetZapLogLevel(err), "Failed to find the permission.").Write(zap.Error(err))
			gb.AbortWithError(c, err)
			return
		}

		if !(p.RepoPerm == perm.RepoPermWrite || p.RepoPerm == perm.RepoPermAdmin) {
			rm.log.Warn("The access is forbidden. Only write permission can access.", zap.String("repo", namespace+"/"+name))
			gb.AbortWithStatusAndError(
				c,
				http.StatusForbidden,
				e.NewErrorWithMessage(e.ErrorPermissionRequired, "Only write permission can access.", nil),
			)
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
		if err != nil {
			rm.log.Check(gb.GetZapLogLevel(err), "Failed to find the repository.").Write(zap.Error(err))
			gb.AbortWithError(c, err)
			return
		}

		p, err := rm.i.FindPermOfRepo(ctx, r, u)
		if e.HasErrorCode(err, e.ErrorCodeNotFound) {
			rm.log.Check(gb.GetZapLogLevel(err), "It is denied to acess the repository.").Write(zap.Error(err))
			gb.AbortWithStatusAndError(c, http.StatusForbidden, err)
			return
		} else if err != nil {
			rm.log.Check(gb.GetZapLogLevel(err), "Failed to find the permission.").Write(zap.Error(err))
			gb.AbortWithError(c, err)
			return
		}

		if p.RepoPerm != perm.RepoPermAdmin {
			rm.log.Warn("The access is forbidden. Only admin permission can access.", zap.String("repo", namespace+"/"+name))
			gb.AbortWithStatusAndError(
				c,
				http.StatusForbidden,
				e.NewErrorWithMessage(e.ErrorPermissionRequired, "Only admin permission can access.", nil),
			)
			return
		}

		c.Set(KeyRepo, r)
	}
}
