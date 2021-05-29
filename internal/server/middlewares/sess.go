package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	SessMiddleware struct {
		store Store
	}
)

func NewSessMiddleware(store Store) *SessMiddleware {
	return &SessMiddleware{
		store: store,
	}
}

func (s *SessMiddleware) User() gin.HandlerFunc {
	return func(c *gin.Context) {
		hash := c.GetString(gb.KeySession)
		if hash == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized user",
			})
		}

		ctx := c.Request.Context()

		u, err := s.store.FindUserByHash(ctx, hash)
		if ent.IsNotFound(err) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized user",
			})
		}

		c.Set(gb.KeyUser, u)
	}
}

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := c.Cookie(gb.CookieSession)
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		}
		if s != "" {
			c.Set(gb.KeySession, s)
			return
		}

		header := c.GetHeader("Authorization")
		s = strings.TrimPrefix(header, "Bearer ")
		if s != "" {
			c.Set(gb.KeySession, s)
			return
		}
	}
}

func OnlyAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get(gb.KeySession)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized user",
			})
		}
	}
}
