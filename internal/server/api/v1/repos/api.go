package repos

import "go.uber.org/zap"

type (
	API struct {
		// Services used for talking to different parts of the entities.
		Repos              *RepoService
		Commits            *CommitService
		Branches           *BranchService
		Tags               *TagService
		Deployments        *DeploymentService
		Config             *ConfigService
		Reviews            *ReviewService
		DeploymentStatuses *DeploymentStatusService
		Locks              *LockService
		Perms              *PermsService
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
	api.Repos = &RepoService{
		service:       service{i: i, log: log},
		WebhookURL:    c.WebhookURL,
		WebhookSSL:    c.WebhookSSL,
		WebhookSecret: c.WebhookSecret,
	}
	api.Commits = &CommitService{i: i, log: log.Named("commits")}
	api.Branches = &BranchService{i: i, log: log.Named("branches")}
	api.Tags = &TagService{i: i, log: log.Named("tags")}
	api.Deployments = &DeploymentService{i: i, log: log.Named("deployments")}
	api.Config = &ConfigService{i: i, log: log.Named("config")}
	api.Reviews = &ReviewService{i: i, log: log.Named("reviews")}
	api.DeploymentStatuses = &DeploymentStatusService{i: i, log: log.Named("deployment_statuses")}
	api.Locks = &LockService{i: i, log: log.Named("locks")}
	api.Perms = &PermsService{i: i, log: log.Named("perms")}

	return api
}
