package repos

import "go.uber.org/zap"

type (
	API struct {
		// APIs used for talking to different parts of the entities.
		Repos              *RepoAPI
		Commits            *CommitAPI
		Branches           *BranchAPI
		Tags               *TagAPI
		Deployments        *DeploymentAPI
		Config             *ConfigAPI
		Reviews            *ReviewAPI
		DeploymentStatuses *DeploymentStatusAPI
		Locks              *LockAPI
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
	api.Repos = &RepoAPI{
		service:       service{i: i, log: log},
		WebhookURL:    c.WebhookURL,
		WebhookSSL:    c.WebhookSSL,
		WebhookSecret: c.WebhookSecret,
	}
	api.Commits = &CommitAPI{i: i, log: log.Named("commits")}
	api.Branches = &BranchAPI{i: i, log: log.Named("branches")}
	api.Tags = &TagAPI{i: i, log: log.Named("tags")}
	api.Deployments = &DeploymentAPI{i: i, log: log.Named("deployments")}
	api.Config = &ConfigAPI{i: i, log: log.Named("config")}
	api.Reviews = &ReviewAPI{i: i, log: log.Named("reviews")}
	api.DeploymentStatuses = &DeploymentStatusAPI{i: i, log: log.Named("deployment_statuses")}
	api.Locks = &LockAPI{i: i, log: log.Named("locks")}
	api.Perms = &PermsAPI{i: i, log: log.Named("perms")}

	return api
}
