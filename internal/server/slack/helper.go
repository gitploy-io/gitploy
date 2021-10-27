package slack

import (
	"strconv"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/slack-go/slack"
)

func postResponseMessage(channelID, responseURL, message string) error {
	_, _, _, err := slack.
		New("").
		SendMessage(
			channelID,
			slack.MsgOptionResponseURL(responseURL, "ephemeral"),
			slack.MsgOptionText(message, false),
		)
	return err
}

func postBotMessage(cu *ent.ChatUser, message string) error {
	_, _, _, err := slack.
		New(cu.BotToken).
		SendMessage(
			cu.ID,
			slack.MsgOptionText(message, false),
		)
	return err
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
