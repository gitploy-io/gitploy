package slack

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/notification"
)

func (s *Slack) Notify(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	if n.Type == notification.TypeDeployment {
		return s.notifyDeploymentCreated(ctx, cu, n)
	}

	return fmt.Errorf("type have to be one of \"deployment_created\".")
}

func (s *Slack) notifyDeploymentCreated(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	const (
		title = "New Deployment"
	)

	var (
		fullname = fmt.Sprintf("%s/%s", n.RepoNamespace, n.RepoName)
		ref      string
	)

	if n.DeploymentType == string(deployment.TypeCommit) {
		ref = n.DeploymentRef[:7]
	} else {
		ref = n.DeploymentRef
	}

	_, _, err := s.Client(cu).
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color: mapDeploymentStatusToColor(n.DeploymentStatus),
			Blocks: slack.Blocks{
				BlockSet: []slack.Block{
					slack.SectionBlock{
						Type: slack.MBTSection,
						Text: &slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: fmt.Sprintf("*New Deployment #%d*", n.DeploymentNumber),
						},
					},
					slack.SectionBlock{
						Type: slack.MBTSection,
						Text: &slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: fmt.Sprintf("*%s* - *%s* deploy `%s` to *%s* environment.", fullname, n.DeploymentLogin, ref, n.DeploymentEnv),
						},
					},
				},
			},
		}))

	return err
}

func mapDeploymentStatusToColor(status string) string {
	const (
		gray   = "#bfbfbf"
		purple = "#722ed1"
		green  = "#52c41a"
		red    = "#f5222d"
	)
	switch deployment.Status(status) {
	case deployment.StatusWaiting:
		return gray
	case deployment.StatusCreated:
		return purple
	case deployment.StatusRunning:
		return purple
	case deployment.StatusSuccess:
		return green
	case deployment.StatusFailure:
		return red
	default:
		return gray
	}
}
