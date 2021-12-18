//go:build oss

package slack

func NewSlack(c *SlackConfig) *Slack {
	return &Slack{}
}
