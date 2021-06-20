package slack

import (
	"github.com/slack-go/slack"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Slack struct{}
)

func NewSlack() *Slack {
	return &Slack{}
}

func (s *Slack) Client(cu *ent.ChatUser) *slack.Client {
	return slack.New(cu.BotToken)
}
