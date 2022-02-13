package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/gitploy-io/gitploy/internal/server/api/shared"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/license"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/search"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/stream"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/sync"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/users"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/internal/server/hooks"
	"github.com/gitploy-io/gitploy/internal/server/metrics"
	s "github.com/gitploy-io/gitploy/internal/server/slack"
	"github.com/gitploy-io/gitploy/internal/server/web"
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

		PrometheusEnabled    bool
		PrometheusAuthSecret string
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
		BotScopes    []string
		UserScopes   []string
	}

	Interactor interface {
		gb.Interactor
		s.Interactor
		sync.Interactor
		shared.Interactor
		web.Interactor
		repos.Interactor
		users.Interactor
		stream.Interactor
		hooks.Interactor
		search.Interactor
		license.Interactor
		metrics.Interactor
	}
)

func init() {
	// always release mode.
	gin.SetMode("release")
}

func NewRouter(c *RouterConfig) *gin.Engine {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		// AllowOrigins: []string{"http://localhost:3000"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Accept", "Content-Length", "Content-Type"},
	}))

	gm := gb.NewMiddleware(c.Interactor)
	r.Use(
		gm.SetUser(),
	)

	v1 := r.Group("/api/v1")
	{
		m := shared.NewMiddleware(c.Interactor)
		v1.Use(
			m.OnlyAuthorized(),
			m.IsLicenseExpired(),
		)
	}

	syncv1 := v1.Group("/sync")
	{
		s := sync.NewSyncher(c.Interactor)
		syncv1.POST("", s.Sync)
	}

	repov1 := v1.Group("/repos")
	{
		rm := repos.NewRepoMiddleware(c.Interactor)
		api := repos.NewAPI(repos.APIConfig{
			Interactor:    c.Interactor,
			WebhookURL:    fmt.Sprintf("%s://%s/hooks", c.ProxyProto, c.ProxyHost),
			WebhookSSL:    c.ProxyProto == "https",
			WebhookSecret: c.WebhookSecret,
		})
		repov1.GET("", api.Repo.List)
		repov1.GET("/:namespace/:name", rm.RepoReadPerm(), api.Repo.Get)
		repov1.PATCH("/:namespace/:name", rm.RepoAdminPerm(), api.Repo.Update)
		repov1.GET("/:namespace/:name/commits", rm.RepoReadPerm(), api.Commits.List)
		repov1.GET("/:namespace/:name/commits/:sha", rm.RepoReadPerm(), api.Commits.Get)
		repov1.GET("/:namespace/:name/commits/:sha/statuses", rm.RepoReadPerm(), api.Commits.ListStatuses)
		repov1.GET("/:namespace/:name/branches", rm.RepoReadPerm(), api.Branch.List)
		repov1.GET("/:namespace/:name/branches/:branch", rm.RepoReadPerm(), api.Branch.Get)
		repov1.GET("/:namespace/:name/tags", rm.RepoReadPerm(), api.Tag.List)
		repov1.GET("/:namespace/:name/tags/:tag", rm.RepoReadPerm(), api.Tag.Get)
		repov1.GET("/:namespace/:name/deployments", rm.RepoReadPerm(), api.Deployment.List)
		repov1.POST("/:namespace/:name/deployments", rm.RepoWritePerm(), api.Deployment.Create)
		repov1.GET("/:namespace/:name/deployments/:number", rm.RepoReadPerm(), api.Deployment.Get)
		repov1.PUT("/:namespace/:name/deployments/:number", rm.RepoWritePerm(), api.Deployment.Update)
		repov1.GET("/:namespace/:name/deployments/:number/changes", rm.RepoReadPerm(), api.Deployment.ListChanges)
		repov1.POST("/:namespace/:name/deployments/:number/rollback", rm.RepoWritePerm(), api.Deployment.Rollback)
		repov1.GET("/:namespace/:name/deployments/:number/reviews", rm.RepoReadPerm(), api.Review.List)
		repov1.GET("/:namespace/:name/deployments/:number/review", rm.RepoReadPerm(), api.Review.GetMine)
		repov1.PATCH("/:namespace/:name/deployments/:number/review", rm.RepoReadPerm(), api.Review.UpdateMine)
		repov1.GET("/:namespace/:name/deployments/:number/statuses", rm.RepoReadPerm(), api.DeploymentStatus.List)
		repov1.POST("/:namespace/:name/deployments/:number/remote-statuses", rm.RepoReadPerm(), api.DeploymentStatus.CreateRemote)
		repov1.GET("/:namespace/:name/locks", rm.RepoReadPerm(), api.Lock.List)
		repov1.POST("/:namespace/:name/locks", rm.RepoWritePerm(), api.Lock.Create)
		repov1.PATCH("/:namespace/:name/locks/:lockID", rm.RepoWritePerm(), api.Lock.Update)
		repov1.DELETE("/:namespace/:name/locks/:lockID", rm.RepoWritePerm(), api.Lock.Delete)
		repov1.GET("/:namespace/:name/perms", rm.RepoReadPerm(), api.Perm.List)
		repov1.GET("/:namespace/:name/config", rm.RepoReadPerm(), api.Config.Get)
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

	searchv1 := v1.Group("/search")
	{
		s := search.NewSearch(c.Interactor)
		searchv1.GET("/deployments", s.SearchDeployments)
		searchv1.GET("/reviews", s.SearchAssignedReviews)
	}

	licensev1 := v1.Group("/license")
	{
		l := license.NewLicenser(c.Interactor)
		licensev1.GET("", l.GetLicense)
	}

	hooksapi := r.Group("/hooks")
	{
		hc := &hooks.ConfigHooks{
			WebhookSecret: c.WebhookSecret,
		}
		h := hooks.NewHooks(hc, c.Interactor)
		hooksapi.POST("", h.HandleHook)
	}

	metricsapi := r.Group("/metrics")
	{
		r.Use(metrics.CollectRequestMetrics())

		m := metrics.NewMetric(&metrics.MetricConfig{
			Interactor:           c.Interactor,
			PrometheusAuthSecret: c.PrometheusAuthSecret,
		})
		metricsapi.GET("", hasOptIn(c.PrometheusEnabled), m.CollectMetrics)
	}

	r.HEAD("/slack", func(gc *gin.Context) {
		if isSlackEnabled(c) {
			gc.Status(http.StatusOK)
			return
		}
		gc.Status(http.StatusNotFound)
	})
	if isSlackEnabled(c) {
		slackapi := r.Group("/slack")
		{
			slack := s.NewSlack(&s.SlackConfig{
				ServerHost:  c.Host,
				ServerProto: c.Proto,
				Config:      newSlackOauthConfig(c),
				Interactor:  c.Interactor,
			})
			slackapi.GET("", slack.Index)
			slackapi.GET("/signin", slack.Signin)
			slackapi.GET("/signout", slack.Signout)
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
			root.GET("/signout", w.SignOutHTML)
		} else {
			root.GET("/", w.IndexString)
			root.GET("/signout", w.SignOutString)
		}

		root.GET("/signin", w.Signin)
		root.GET("/link/:namespace/:name/config", w.RedirectToConfig)
		root.GET("/link/:namespace/:name/config/new", w.RedirectToNewConfig)

		// Static files located at the 'ui/public' directory.
		r.StaticFile("/favicon.ico", "./favicon.ico")
		r.StaticFile("/spinner.ico", "./spinner.ico")
		r.StaticFile("/manifest.json", "./manifest.json")
		r.StaticFile("/robots.txt", "./robots.txt")
		r.StaticFile("/logo192.png", "./logo192.png")
		r.Static("/static", "./static")
		if _, err := os.Stat("./index.html"); err == nil {
			r.NoRoute(w.IndexHTML)
		} else {
			r.NoRoute(w.IndexString)
		}
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

func hasOptIn(enabled bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enabled {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
}
