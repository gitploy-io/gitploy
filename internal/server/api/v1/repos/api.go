package repos

import "go.uber.org/zap"

type (
	API struct {
		common *service

		// APIs used for talking to different parts of the entities.
		Repo             *RepoAPI
		Commits          *CommitAPI
		Branch           *BranchAPI
		Tag              *TagAPI
		Deployment       *DeploymentAPI
		Config           *ConfigAPI
		Review           *ReviewAPI
		DeploymentStatus *DeploymentStatusAPI
		Lock             *LockAPI
		Perm             *PermAPI
	}

	APIConfig struct {
		Interactor

		WebhookURL    string
		WebhookSSL    bool
		WebhookSecret string
	}

	service struct {
		i   Interactor
		log *zap.Logger
	}
)

func NewAPI(c APIConfig) *API {
	api := &API{}

	api.common = &service{
		i:   c.Interactor,
		log: zap.L().Named("repos"),
	}
	api.Repo = &RepoAPI{
		service:       api.common,
		WebhookURL:    c.WebhookURL,
		WebhookSSL:    c.WebhookSSL,
		WebhookSecret: c.WebhookSecret,
	}
	api.Commits = (*CommitAPI)(api.common)
	api.Branch = (*BranchAPI)(api.common)
	api.Tag = (*TagAPI)(api.common)
	api.Deployment = (*DeploymentAPI)(api.common)
	api.Config = (*ConfigAPI)(api.common)
	api.Review = (*ReviewAPI)(api.common)
	api.DeploymentStatus = (*DeploymentStatusAPI)(api.common)
	api.Lock = (*LockAPI)(api.common)
	api.Perm = (*PermAPI)(api.common)

	return api
}
