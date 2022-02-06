package repos

type (
	RepoService struct {
		service

		WebhookURL    string
		WebhookSSL    bool
		WebhookSecret string
	}
)
