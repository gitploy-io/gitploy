package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/pkg/github"
	"github.com/gitploy-io/gitploy/internal/pkg/store"
	"github.com/gitploy-io/gitploy/internal/server"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var envfile string
	flag.StringVar(&envfile, "env-file", ".env", "Read in a file of environment variables")
	flag.Parse()

	godotenv.Load(envfile)
	c, err := NewConfigFromEnv()
	if err != nil {
		log.Fatalf("main: invalid configuration: %s", err)
	}

	setGlobalLogger(c.DebugMode)

	r := server.NewRouter(newRouterConfig(c))
	if err := runServer(r, c); err != nil {
		log.Printf("The server is down: %s.", err)
	}
}

func runServer(r *gin.Engine, c *Config) error {
	if !c.hasTLS() {
		return endless.ListenAndServe(":http", r)
	}

	// Redirect http request to https server.
	endless.ListenAndServe(":http", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		target := "https://" + req.Host + req.URL.Path
		http.Redirect(w, req, target, http.StatusTemporaryRedirect)
	}))

	return endless.ListenAndServeTLS(":https", c.TLSCert, c.TLSKey, r)
}

func setGlobalLogger(debug bool) {
	var config zap.Config
	if debug {
		config = zap.NewDevelopmentConfig()
		config.Encoding = "json"
	} else {
		config = zap.NewProductionConfig()
		config.DisableStacktrace = true
	}

	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
}

func newRouterConfig(c *Config) *server.RouterConfig {
	return &server.RouterConfig{
		ServerConfig: newServerConfig(c),
		SCMConfig:    newSCMConfig(c),
		ChatConfig:   newChatConfig(c),
		Interactor:   NewInteractor(c),
	}
}

func newServerConfig(c *Config) *server.ServerConfig {
	var (
		proxyHost  string
		proxyProto string
	)
	if c.ServerProxyHost != "" {
		proxyHost = c.ServerProxyHost
		proxyProto = c.ServerProxyProto
	} else {
		proxyHost = c.ServerHost
		proxyProto = c.ServerProto
	}

	return &server.ServerConfig{
		Host:          c.ServerHost,
		Proto:         c.ServerProto,
		ProxyHost:     proxyHost,
		ProxyProto:    proxyProto,
		WebhookSecret: c.WebhookSecret,
	}
}

func newSCMConfig(c *Config) *server.SCMConfig {
	var sc *server.SCMConfig

	if c.isGithubEnabled() {
		sc = &server.SCMConfig{
			Type:         server.SCMTypeGithub,
			ClientID:     c.GithubClientID,
			ClientSecret: c.GithubClientSecret,
			Scopes:       c.GithubScopes,
		}
	}

	return sc
}

func newChatConfig(c *Config) *server.ChatConfig {
	var cc *server.ChatConfig

	if c.isSlackEnabled() {
		cc = &server.ChatConfig{
			Type:         server.ChatTypeSlack,
			ClientID:     c.SlackClientID,
			ClientSecret: c.SlackClientSecret,
			Secret:       c.SlackSigningSecret,
			BotScopes:    c.SlackBotScopes,
			UserScopes:   c.SlackUserScopes,
		}
	}

	return cc
}

func NewInteractor(c *Config) server.Interactor {
	return interactor.NewInteractor(
		&interactor.InteractorConfig{
			ServerHost:  c.ServerHost,
			ServerProto: c.ServerProto,
			OrgEntries:  c.OrganizationEntries,
			AdminUsers:  c.AdminUsers,
			LicenseKey:  c.License,
			Store:       newStore(c),
			SCM:         newSCM(c),
		},
	)
}

func newStore(c *Config) interactor.Store {
	client, err := OpenDB(c.StoreDriver, c.StoreSource)
	if err != nil {
		log.Fatalf("It has failed to open the DB: %v", err)
	}

	err = client.Schema.Create(
		context.Background(),
	)
	if err != nil {
		log.Fatalf("It has failed to migrate the table schema: %v", err)
	}

	return store.NewStore(client)
}

func newSCM(c *Config) interactor.SCM {
	var scm interactor.SCM

	if c.isGithubEnabled() {
		scm = github.NewGithub()
	}

	return scm
}
