package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

func (m *Middleware) OnlyAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get(gb.KeyUser)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized user",
			})
		}
	}
}
