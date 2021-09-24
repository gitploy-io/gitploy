package main

import (
	"log"

	"entgo.io/ent/dialect"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Server
		Store
		Github
		Slack
		Webhook
	}

	Server struct {
		DebugMode bool `default:"false" split_words:"true"`

		ServerHost       string `required:"true" split_words:"true"`
		ServerProto      string `required:"true" default:"https" split_words:"true"`
		ServerProxyHost  string `split_words:"true"`
		ServerProxyProto string `default:"https" split_words:"true"`

		OrganizationEntries []string `split_words:"true"`
		MemberEntries       []string `split_words:"true"`
		AdminUsers          []string `split_words:"true"`

		License string `split_words:"true"`

		TLSCert string `split_words:"true"`
		TLSKey  string `split_words:"true"`
	}

	Store struct {
		StoreDriver string `required:"true" default:"sqlite3" split_words:"true"`
		StoreSource string `required:"true" default:"file:/data/sqlite3.db?cache=shared&_fk=1" split_words:"true"`
	}

	Github struct {
		GithubClientID     string   `split_words:"true"`
		GithubClientSecret string   `split_words:"true"`
		GithubScopes       []string `split_words:"true" default:"repo,read:user,read:org"`
	}

	Slack struct {
		SlackClientID      string   `split_words:"true"`
		SlackClientSecret  string   `split_words:"true"`
		SlackSigningSecret string   `split_words:"true"`
		SlackUserScopes    []string `split_words:"true" default:""`
		SlackBotScopes     []string `split_words:"true" default:"commands,chat:write"`
	}

	Webhook struct {
		WebhookSecret string `split_words:"true"`
	}
)

func NewConfigFromEnv() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("gitploy", c)
	return c, err
}

func (c *Config) Validate() {
	if !(c.ServerProto == "http" || c.ServerProto == "https") {
		log.Fatal("GITPLOY_SERVER_PROTO have to be \"http\" or \"https\".")
	}

	if driver := c.Store.StoreDriver; driver != dialect.SQLite &&
		driver != dialect.MySQL &&
		driver != dialect.Postgres {
		log.Fatal("GITPLOY_STORE_DRIVER have to be one of them: sqlite3, mysql, or postgres.")
	}
}

func (c *Config) isGithubEnabled() bool {
	return c.GithubClientID != "" && c.GithubClientSecret != ""
}

func (c *Config) isSlackEnabled() bool {
	return c.SlackClientID != "" && c.SlackClientSecret != "" && c.SlackSigningSecret != ""
}

func (c *Config) hasTLS() bool {
	return c.TLSCert != "" && c.TLSKey != ""
}
