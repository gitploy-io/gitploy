package main

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		Store
		SCM
	}

	Store struct {
		StoreDriver string `required:"true" default:"sqlite3"`
		StoreSource string `required:"true" default:"file:./data/sqlite3.db?cache=shared&_fk=1"`
	}

	SCM struct {
		ClientID     string `required:"true" split_words:"true"`
		ClientSecret string `required:"true" split_words:"true"`
	}
)

func NewConfigFromEnv() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("gitploy", c)
	return c, err
}
