package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
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
