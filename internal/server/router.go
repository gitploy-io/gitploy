package server

import (
	"github.com/gin-gonic/gin"
)

type (
	RouterConfig struct {
	}
)

func init() {
	// always release mode.
	gin.SetMode("release")
}

func NewRouter(c *RouterConfig) *gin.Engine {
	r := gin.New()

	r.Use(Session())

	return r
}
