package main

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		SCM
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
