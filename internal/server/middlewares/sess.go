package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	SessMiddleware struct {
		i Interactor
	}
)

func NewSessMiddleware(i Interactor) *SessMiddleware {
	return &SessMiddleware{
		i: i,
	}
}

func (s *SessMiddleware) User() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		u, err := s.i.FindUserByHash(ctx, FindHash(c))
		if ent.IsNotFound(err) {
			return
		}

		c.Set(gb.KeyUser, u)
	}
}

func FindHash(c *gin.Context) string {
	s, _ := c.Cookie(gb.CookieSession)
	if s != "" {
		return s
	}

	header := c.GetHeader("Authorization")
	s = strings.TrimPrefix(header, "Bearer ")
	if s != "" {
		return s
	}

	return ""
}

func OnlyAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get(gb.KeyUser)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized user",
			})
		}
	}
}
