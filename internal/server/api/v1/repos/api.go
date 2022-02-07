package repos

import "go.uber.org/zap"

type (
	API struct {
		common *service

		// APIs used for talking to different parts of the entities.
		Repos              *ReposAPI
		Commits            *CommitsAPI
		Branches           *BranchesAPI
		Tags               *TagsAPI
		Deployments        *DeploymentsAPI
		Config             *ConfigAPI
		Reviews            *ReviewsAPI
		DeploymentStatuses *DeploymentStatusesAPI
		Locks              *LocksAPI
		Perms              *PermsAPI
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
	api.Repos = &ReposAPI{
		service:       api.common,
		WebhookURL:    c.WebhookURL,
		WebhookSSL:    c.WebhookSSL,
		WebhookSecret: c.WebhookSecret,
	}
	api.Commits = (*CommitsAPI)(api.common)
	api.Branches = (*BranchesAPI)(api.common)
	api.Tags = (*TagsAPI)(api.common)
	api.Deployments = (*DeploymentsAPI)(api.common)
	api.Config = (*ConfigAPI)(api.common)
	api.Reviews = (*ReviewsAPI)(api.common)
	api.DeploymentStatuses = (*DeploymentStatusesAPI)(api.common)
	api.Locks = (*LocksAPI)(api.common)
	api.Perms = (*PermsAPI)(api.common)

	return api
}
