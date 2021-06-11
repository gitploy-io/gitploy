package vo

type (
	WebhookConfig struct {
		URL         string
		Secret      string
		InsecureSSL bool
	}
)
