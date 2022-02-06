package repos

import "go.uber.org/zap"

type (
	API struct {
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
	i := c.Interactor
	log := zap.L().Named("repos")

	api := &API{}
	api.Repos = &ReposAPI{
		service:       service{i: i, log: log},
		WebhookURL:    c.WebhookURL,
		WebhookSSL:    c.WebhookSSL,
		WebhookSecret: c.WebhookSecret,
	}
	api.Commits = &CommitsAPI{i: i, log: log.Named("commits")}
	api.Branches = &BranchesAPI{i: i, log: log.Named("branches")}
	api.Tags = &TagsAPI{i: i, log: log.Named("tags")}
	api.Deployments = &DeploymentsAPI{i: i, log: log.Named("deployments")}
	api.Config = &ConfigAPI{i: i, log: log.Named("config")}
	api.Reviews = &ReviewsAPI{i: i, log: log.Named("reviews")}
	api.DeploymentStatuses = &DeploymentStatusesAPI{i: i, log: log.Named("deployment_statuses")}
	api.Locks = &LocksAPI{i: i, log: log.Named("locks")}
	api.Perms = &PermsAPI{i: i, log: log.Named("perms")}

	return api
}
