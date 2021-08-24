package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/internal/interactor"
	"github.com/hanjunlee/gitploy/internal/pkg/github"
	"github.com/hanjunlee/gitploy/internal/pkg/store"
	"github.com/hanjunlee/gitploy/internal/server"

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

	setGlobalLogger(true)

	r := server.NewRouter(newRouterConfig(c))
	log.Printf("Run server with port %d ...", c.ServerPort)
	r.Run(fmt.Sprintf(":%d", c.ServerPort))
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
			RepoEntries: c.RepositoryEntries,
			AdminUsers:  c.AdminUsers,
			LicenseKey:  c.License,
			Store:       newStore(c),
			SCM:         newSCM(c),
		},
	)
}

func newStore(c *Config) interactor.Store {
	client, err := ent.Open(c.StoreDriver, c.StoreSource)
	if err != nil {
		log.Fatalf("failed create the connection for store: %v", err)
	}

	err = client.Schema.Create(
		context.Background(),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
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
