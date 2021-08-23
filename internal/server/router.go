package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/hanjunlee/gitploy/internal/server/api/v1/license"
	"github.com/hanjunlee/gitploy/internal/server/api/v1/repos"
	"github.com/hanjunlee/gitploy/internal/server/api/v1/search"
	"github.com/hanjunlee/gitploy/internal/server/api/v1/stream"
	"github.com/hanjunlee/gitploy/internal/server/api/v1/sync"
	"github.com/hanjunlee/gitploy/internal/server/api/v1/users"
	"github.com/hanjunlee/gitploy/internal/server/hooks"
	mw "github.com/hanjunlee/gitploy/internal/server/middlewares"
	s "github.com/hanjunlee/gitploy/internal/server/slack"
	"github.com/hanjunlee/gitploy/internal/server/web"
)

const (
	SCMTypeGithub SCMType  = "github"
	ChatTypeSlack ChatType = "slack"
)

type (
	RouterConfig struct {
		*ServerConfig
		*SCMConfig
		*ChatConfig
		Interactor
	}

	ServerConfig struct {
		Host       string
		Proto      string
		ProxyHost  string
		ProxyProto string

		WebhookSecret string
	}

	SCMType string

	SCMConfig struct {
		Type         SCMType
		ClientID     string
		ClientSecret string
		Scopes       []string
	}

	ChatType string

	ChatConfig struct {
		Type         ChatType
		ClientID     string
		ClientSecret string
		Secret       string
		BotScopes    []string
		UserScopes   []string
	}

	Interactor interface {
		s.Interactor
		sync.Interactor
		mw.Interactor
		web.Interactor
		repos.Interactor
		users.Interactor
		stream.Interactor
		hooks.Interactor
		search.Interactor
		license.Interactor
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
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Accept", "Content-Length", "Content-Type"},
	}))
	sm := mw.NewSessMiddleware(c.Interactor)
	lm := mw.NewLicenseMiddleware(c.Interactor)

	r.Use(sm.User())

	v1 := r.Group("/api/v1")
	{
		// Only authorized user can access to API.
		v1.Use(mw.OnlyAuthorized())
		// Check license expired.
		v1.Use(lm.IsExpired())
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
				WebhookURL:    fmt.Sprintf("%s://%s/hooks", c.ProxyProto, c.ProxyHost),
				WebhookSSL:    c.ProxyProto == "https",
				WebhookSecret: c.WebhookSecret,
			},
			c.Interactor,
		)
		repov1.GET("", r.ListRepos)
		repov1.GET("/search", r.GetRepoByNamespaceName)
		repov1.GET("/:id", rm.RepoReadPerm(), r.GetRepo)
		repov1.PATCH("/:id", rm.RepoAdminPerm(), r.UpdateRepo)
		repov1.GET("/:id/commits", rm.RepoReadPerm(), r.ListCommits)
		repov1.GET("/:id/commits/:sha", rm.RepoReadPerm(), r.GetCommit)
		repov1.GET("/:id/commits/:sha/statuses", rm.RepoReadPerm(), r.ListStatuses)
		repov1.GET("/:id/branches", rm.RepoReadPerm(), r.ListBranches)
		repov1.GET("/:id/branches/:branch", rm.RepoReadPerm(), r.GetBranch)
		repov1.GET("/:id/tags", rm.RepoReadPerm(), r.ListTags)
		repov1.GET("/:id/tags/:tag", rm.RepoReadPerm(), r.GetTag)
		repov1.GET("/:id/deployments", rm.RepoReadPerm(), r.ListDeployments)
		repov1.POST("/:id/deployments", rm.RepoWritePerm(), r.CreateDeployment)
		repov1.GET("/:id/deployments/:number", rm.RepoReadPerm(), r.GetDeploymentByNumber)
		repov1.PATCH("/:id/deployments/:number", rm.RepoWritePerm(), r.UpdateDeployment)
		repov1.GET("/:id/deployments/:number/changes", rm.RepoReadPerm(), r.ListDeploymentChanges)
		repov1.POST("/:id/deployments/:number/rollback", rm.RepoWritePerm(), r.RollbackDeployment)
		repov1.GET("/:id/deployments/:number/approvals", rm.RepoReadPerm(), r.ListApprovals)
		repov1.POST("/:id/deployments/:number/approvals", rm.RepoReadPerm(), r.CreateApproval)
		repov1.GET("/:id/deployments/:number/approval", rm.RepoReadPerm(), r.GetMyApproval)
		repov1.PATCH("/:id/deployments/:number/approval", rm.RepoReadPerm(), r.UpdateApproval)
		repov1.GET("/:id/approvals/:aid", rm.RepoReadPerm(), r.GetApproval)
		repov1.DELETE("/:id/approvals/:aid", rm.RepoReadPerm(), r.DeleteApproval)
		repov1.GET("/:id/perms", rm.RepoReadPerm(), r.ListPerms)
		repov1.GET("/:id/config", rm.RepoReadPerm(), r.GetConfig)
	}

	usersv1 := v1.Group("/users")
	userv1 := v1.Group("/user")
	{
		m := users.NewUserMiddleware()
		u := users.NewUser(c.Interactor)
		usersv1.GET("", m.AdminOnly(), u.ListUsers)
		usersv1.PATCH("/:id", m.AdminOnly(), u.UpdateUser)
		usersv1.DELETE("/:id", m.AdminOnly(), u.DeleteUser)
		userv1.GET("", u.GetMyUser)
		userv1.GET("/rate-limit", u.GetRateLimit)
	}

	streamv1 := v1.Group("/stream")
	{
		s := stream.NewStream(c.Interactor)
		streamv1.GET("/events", s.GetEvents)
	}

	searchapi := v1.Group("/search")
	{
		s := search.NewSearch(c.Interactor)
		searchapi.GET("/deployments", s.SearchDeployments)
		searchapi.GET("/approvals", s.SearchApprovals)
	}

	licenseapi := v1.Group("/license")
	{
		l := license.NewLicenser(c.Interactor)
		licenseapi.GET("", l.GetLicense)
	}

	hooksapi := r.Group("/hooks")
	{
		hc := &hooks.ConfigHooks{
			WebhookSecret: c.WebhookSecret,
		}
		h := hooks.NewHooks(hc, c.Interactor)
		hooksapi.POST("", h.HandleHook)
	}

	if isSlackEnabled(c) {
		slackapi := r.Group("/slack")
		{
			m := s.NewSlackMiddleware(c.ChatConfig.Secret)
			slack := s.NewSlack(&s.SlackConfig{
				ServerHost:  c.Host,
				ServerProto: c.Proto,
				Config:      newSlackOauthConfig(c),
				Interactor:  c.Interactor,
			})
			slackapi.GET("", slack.Index)
			slackapi.GET("/signin", slack.SigninSlack)
			// TODO: add signout
			slackapi.POST("/interact", m.Verify(), slack.Interact)
			slackapi.POST("/command", m.Verify(), slack.Cmd)
			// Check Slack is enabled or not.
			slackapi.GET("/ping", func(c *gin.Context) { c.Status(http.StatusOK) })
		}
	}

	root := r.Group("")
	{
		w := web.NewWeb(&web.WebConfig{
			Config:     newGithubOauthConfig(c),
			Interactor: c.Interactor,
		})

		if _, err := os.Stat("./index.html"); err == nil {
			r.LoadHTMLFiles("./index.html")
			root.GET("/", w.IndexHTML)
		} else {
			root.GET("/", w.IndexString)
		}

		root.GET("/signin", w.Signin)
		root.GET("/signout", w.SignOut)

		// Static files
		// Files in ui/public
		r.StaticFile("/favicon.ico", "./favicon.ico")
		r.StaticFile("/logo192.png", "./logo192.png")
		r.StaticFile("/logo512.png", "./logo512.png")
		r.StaticFile("/manifest.json", "./manifest.json")
		r.StaticFile("/robots.txt", "./robots.txt")
		r.Static("/static", "./static")
		r.NoRoute(w.IndexHTML)
	}

	return r
}

func newGithubOauthConfig(c *RouterConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.SCMConfig.ClientID,
		ClientSecret: c.SCMConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		RedirectURL: fmt.Sprintf("%s://%s/signin", c.Proto, c.Host),
		Scopes:      c.SCMConfig.Scopes,
	}
}

func newSlackOauthConfig(c *RouterConfig) *oauth2.Config {
	if c.ChatConfig == nil {
		return nil
	}

	return &oauth2.Config{
		ClientID:     c.ChatConfig.ClientID,
		ClientSecret: c.ChatConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://slack.com/oauth/v2/authorize",
			TokenURL: "https://slack.com/api/oauth.v2.access",
		},
		RedirectURL: fmt.Sprintf("%s://%s/slack/signin", c.Proto, c.Host),
		Scopes:      c.ChatConfig.BotScopes,
	}
}

func isSlackEnabled(c *RouterConfig) bool {
	return c.ChatConfig != nil && c.ChatConfig.Type == ChatTypeSlack
}
