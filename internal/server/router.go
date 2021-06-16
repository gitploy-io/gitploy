package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/hanjunlee/gitploy/internal/server/api/v1/repos"
	"github.com/hanjunlee/gitploy/internal/server/api/v1/sync"
	mw "github.com/hanjunlee/gitploy/internal/server/middlewares"
	s "github.com/hanjunlee/gitploy/internal/server/slack"
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
		Interactor
	}

	ServerConfig struct {
		Host  string
		Proto string

		WebhookHost   string
		WebhookProto  string
		WebhookSecret string
	}

	SCMType string

	SCMConfig struct {
		Type         SCMType
		ClientID     string
		ClientSecret string
		Scopes       []string
	}

	Store interface {
		repos.Store
	}

	SCM interface {
		repos.SCM
	}

	Interactor interface {
		s.Interactor
		sync.Interactor
		mw.Interactor
		web.Interactor
		repos.Interactor
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
		AllowHeaders:     []string{"Origin", "Authorization", "accept", "Content-Length", "Content-Type"},
	}))
	r.Use(mw.Session())

	root := r.Group("/")
	{
		w := web.NewWeb(newGithubOauthConfig(c), c.Interactor)
		root.GET("/", w.Index)
		root.GET("/signin", w.Signin)
	}

	v1 := r.Group("/api/v1")
	{
		sm := mw.NewSessMiddleware(c.Store)
		// Only authed user access to API.
		v1.Use(sm.User())
	}

	syncv1 := v1.Group("/sync")
	{
		s := sync.NewSyncher(c.Interactor)
		syncv1.POST("", s.Sync)
	}

	repov1 := v1.Group("/repos")
	{
		rm := repos.NewRepoMiddleware(c.Interactor)
		r := repos.NewRepo(
			repos.RepoConfig{
				WebhookURL:    fmt.Sprintf("%s://%s/hooks", c.WebhookProto, c.WebhookHost),
				WebhookSecret: c.WebhookSecret,
			},
			c.Store,
			c.SCM,
		)
		repov1.GET("", r.ListRepos)
		repov1.GET("/search", r.GetRepoByNamespaceName)
		repov1.GET("/:repoID", rm.Repo(), r.GetRepo)
		repov1.PATCH("/:repoID", rm.Repo(), r.UpdateRepo)
		repov1.GET("/:repoID/commits", rm.Repo(), r.ListCommits)
		repov1.GET("/:repoID/commits/:sha", rm.Repo(), r.GetCommit)
		repov1.GET("/:repoID/commits/:sha/statuses", rm.Repo(), r.ListStatuses)
		repov1.GET("/:repoID/branches", rm.Repo(), r.ListBranches)
		repov1.GET("/:repoID/branches/:branch", rm.Repo(), r.GetBranch)
		repov1.GET("/:repoID/tags", rm.Repo(), r.ListTags)
		repov1.GET("/:repoID/tags/:tag", rm.Repo(), r.GetTag)
		repov1.GET("/:repoID/deployments", rm.Repo(), r.ListDeployments)
		repov1.POST("/:repoID/deployments", rm.Repo(), r.CreateDeployment)
		repov1.GET("/:repoID/deployments/latest", rm.Repo(), r.GetLatestDeployment)
		repov1.GET("/:repoID/config", rm.Repo(), r.GetConfig)
		repov1.PATCH("/:repoID/activate", rm.Repo(), rm.AdminPerm(), r.Activate)
		repov1.PATCH("/:repoID/deactivate", rm.Repo(), rm.AdminPerm(), r.Deactivate)
	}

	// TODO: add webhook

	slackapi := r.Group("/slack")
	{
		slack := s.NewSlack(c.Interactor)
		slackapi.POST("/", slack.Interact)
		slackapi.POST("/deploy", slack.Deploy)
	}

	return r
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
