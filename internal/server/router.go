package server

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/hanjunlee/gitploy/internal/server/web"
)

type (
	RouterConfig struct {
		*SCMConfig
		Store Store
		SCM   SCM
	}

	SCMConfig struct {
		ClientID     string
		ClientSEcret string
	}

	Store interface {
		web.Store
	}

	SCM interface {
		web.SCM
	}
)

func init() {
	// always release mode.
	gin.SetMode("release")
}

func NewRouter(c *RouterConfig) *gin.Engine {
	r := gin.New()

	r.Use(Session())

	root := r.Group("/")
	{
		w := web.NewWeb(newWebConfig(c))
		root.GET("/", w.Index)
		root.GET("/signin", w.Signin)
	}

	return r
}

func newWebConfig(c *RouterConfig) *web.WebConfig {
	return &web.WebConfig{
		Config: &oauth2.Config{},
		Store:  c.Store,
		SCM:    c.SCM,
	}
}
