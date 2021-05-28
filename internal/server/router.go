package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/hanjunlee/gitploy/internal/server/v1/sync"
	"github.com/hanjunlee/gitploy/internal/server/web"
)

const (
	SCMTypeGithub SCMType = "github"
)

type (
	RouterConfig struct {
		*ServerConfig
		*SCMConfig
		Store Store
		SCM   SCM
	}

	ServerConfig struct {
		Host  string
		Proto string
	}

	SCMType string

	SCMConfig struct {
		Type         SCMType
		ClientID     string
		ClientSecret string
		Scopes       []string
	}

	Store interface {
		web.Store
		sync.Store
	}

	SCM interface {
		web.SCM
		sync.SCM
	}
)

func init() {
	// always release mode.
	gin.SetMode("release")
}

func NewRouter(c *RouterConfig) *gin.Engine {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Length", "Content-Type"},
	}))
	r.Use(Session())

	root := r.Group("/")
	{
		w := web.NewWeb(newWebConfig(c))
		root.GET("/", w.Index)
		root.GET("/signin", w.Signin)
	}

	v1 := r.Group("/v1", OnlyAuthorized())

	syncv1 := v1.Group("/sync")
	{
		s := sync.NewSyncher(c.Store, c.SCM)
		syncv1.POST("", s.Sync)
	}

	return r
}

func newWebConfig(c *RouterConfig) *web.WebConfig {
	return &web.WebConfig{
		Config: newGithubOauthConfig(c),
		Store:  c.Store,
		SCM:    c.SCM,
	}
}

func newGithubOauthConfig(c *RouterConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		RedirectURL: fmt.Sprintf("%s://%s/signin", c.Proto, c.Host),
		Scopes:      c.Scopes,
	}
}
