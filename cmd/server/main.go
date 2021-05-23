package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/migrate"
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
	log.Printf("Run server with port %s ...", c.ServerPort)
	r.Run(fmt.Sprintf(":%s", c.ServerPort))
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
		Store:        newStore(c),
		SCM:          newSCM(c),
	}
}

func newServerConfig(c *Config) *server.ServerConfig {
	return &server.ServerConfig{
		Host:  c.ServerHost,
		Proto: c.ServerProto,
	}
}

func newSCMConfig(c *Config) *server.SCMConfig {
	var sc *server.SCMConfig

	if c.isGithub() {
		sc = &server.SCMConfig{
			Type:         server.SCMTypeGithub,
			ClientID:     c.GithubClientID,
			ClientSecret: c.GithubClientSecret,
			Scopes:       c.GithubScopes,
		}
	}

	return sc
}

func newStore(c *Config) server.Store {
	client, err := ent.Open(c.StoreDriver, c.StoreSource)
	if err != nil {
		log.Fatalf("failed create the connection for store: %v", err)
	}

	err = client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return store.NewStore(client)
}

func newSCM(c *Config) server.SCM {
	var scm server.SCM

	if c.isGithub() {
		scm = github.NewGithub()
	}

	return scm
}
