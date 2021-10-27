package slack

import (
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type (
	Slack struct {
		host  string
		proto string

		c *oauth2.Config
		i Interactor

		log *zap.Logger
	}

	SlackConfig struct {
		ServerHost  string
		ServerProto string
		*oauth2.Config
		Interactor
	}
)
