package main

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		Store
		Github
	}

	Store struct {
		StoreDriver string `required:"true" default:"sqlite3"`
		StoreSource string `required:"true" default:"file:./data/sqlite3.db?cache=shared&_fk=1"`
	}

	Github struct {
		GithubClientID     string   `split_words:"true"`
		GithubClientSecret string   `split_words:"true"`
		GithubScopes       []string `split_words:"true" default:"user:email,repo"`
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
