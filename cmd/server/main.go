package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/internal/server"
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

	e := server.NewRouter(newRouterConfig(c))
	e.Run()
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
		SCMConfig: &server.SCMConfig{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
		},
		Store: nil,
		SCM:   nil,
	}
}
