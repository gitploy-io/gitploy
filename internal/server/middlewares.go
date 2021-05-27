package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := c.Cookie(gb.CookieSession)
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			c.JSON(http.StatusInternalServerError, nil)
			c.Abort()
		}
		if s != "" {
			c.Set(gb.KeySession, s)
			c.Abort()
		}

		header := c.GetHeader("Authorization")
		s = strings.TrimPrefix(header, "Bearer: ")
		if s != "" {
			c.Set(gb.KeySession, s)
			c.Abort()
		}

		gb.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized user")
	}
}
