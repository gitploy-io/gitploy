package repos

type (
	RepoAPI struct {
		*service

		WebhookURL    string
		WebhookSSL    bool
		WebhookSecret string
	}
)
