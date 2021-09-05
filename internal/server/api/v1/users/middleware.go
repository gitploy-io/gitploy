package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

type (
	UserMiddleware struct{}
)

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (m *UserMiddleware) AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		v, _ := c.Get(gb.KeyUser)
		u, _ := v.(*ent.User)

		if !u.Admin {
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{
				"message": "Only admin can access.",
			})
		}
	}
}
