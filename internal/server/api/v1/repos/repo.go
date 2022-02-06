package repos

import (
	"go.uber.org/zap"
)

type (
	Repo struct {
		RepoConfig
		i   Interactor
		log *zap.Logger
	}

	RepoConfig struct {
		WebhookURL    string
		WebhookSSL    bool
		WebhookSecret string
	}

	RepoService struct {
		service

		WebhookURL    string
		WebhookSSL    bool
		WebhookSecret string
	}
)

func NewRepo(c RepoConfig, i Interactor) *Repo {
	return &Repo{
		RepoConfig: c,
		i:          i,
		log:        zap.L().Named("repos"),
	}
}
