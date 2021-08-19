package main

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		DebugMode bool `default:"false"`
		Server
		Store
		Github
		Slack
		Webhook
	}

	Server struct {
		ServerHost       string `required:"true" split_words:"true"`
		ServerProto      string `required:"true" default:"https" split_words:"true"`
		ServerPort       string `required:"true" default:"8080" split_words:"true"`
		ServerProxyHost  string `split_words:"true"`
		ServerProxyProto string `default:"https" split_words:"true"`
		ServerProxyPort  string `default:"8081" split_words:"true"`

		AdminUsers []string `split_words:"true"`
	}

	Store struct {
		StoreDriver string `required:"true" default:"sqlite3" split_words:"true"`
		StoreSource string `required:"true" default:"file:./data/sqlite3.db?cache=shared&_fk=1" split_words:"true"`
	}

	Github struct {
		GithubClientID     string   `split_words:"true"`
		GithubClientSecret string   `split_words:"true"`
		GithubScopes       []string `split_words:"true" default:"repo"`
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

func (c *Config) isGithubEnabled() bool {
	return c.GithubClientID != "" && c.GithubClientSecret != ""
}

func (c *Config) isSlackEnabled() bool {
	return c.SlackClientID != "" && c.SlackClientSecret != "" && c.SlackSigningSecret != ""
}
