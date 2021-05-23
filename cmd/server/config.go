package main

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		DebugMode bool `default:"false"`
		Server
		Store
		Github
	}

	Server struct {
		ServerHost  string `required:"true" split_words:"true"`
		ServerProto string `required:"true" default:"https" split_words:"true"`
		ServerPort  string `required:"true" default:"8080" split_words:"true"`
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
)

func NewConfigFromEnv() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("gitploy", c)
	return c, err
}

func (c *Config) isGithub() bool {
	return c.GithubClientID != "" && c.GithubClientSecret != ""
}
