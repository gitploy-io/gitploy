package repos

type (
	ReposAPI struct {
		*service

		WebhookURL    string
		WebhookSSL    bool
		WebhookSecret string
	}
)
