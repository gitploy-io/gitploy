package global

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	Middleware struct {
		i Interactor
	}
)

func NewMiddleware(i Interactor) *Middleware {
	return &Middleware{
		i: i,
	}
}

func (s *Middleware) SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		u, err := s.i.FindUserByHash(ctx, FindHash(c))
		if ent.IsNotFound(err) {
			return
		}

		c.Set(KeyUser, u)
	}
}

func FindHash(c *gin.Context) string {
	s, _ := c.Cookie(CookieSession)
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
