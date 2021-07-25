package slack

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/notification"
)

const (
	colorGray   = "#bfbfbf"
	colorPurple = "#722ed1"
	colorGreen  = "#52c41a"
	colorRed    = "#f5222d"
)

func (s *Slack) Notify(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	if n.Type == notification.TypeDeployment {
		return s.notifyDeploymentCreated(ctx, cu, n)
	} else if n.Type == notification.TypeApprovalRequested {
		return s.notifyApprovalRequested(ctx, cu, n)
	} else if n.Type == notification.TypeApprovalResponded {
		return s.notifyApprovalResponded(ctx, cu, n)
	}

	return fmt.Errorf("type have to be one of \"deployment_created\".")
}

func (s *Slack) notifyDeploymentCreated(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	fullname := fmt.Sprintf("%s/%s", n.RepoNamespace, n.RepoName)

	var ref string
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
							Text: fmt.Sprintf("*%s* - %s deploy `%s` to `%s` environment.", fullname, n.DeploymentLogin, ref, n.DeploymentEnv),
						},
					},
				},
			},
		}))

	return err
}

func (s *Slack) notifyApprovalRequested(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	fullName := fmt.Sprintf("%s/%s", n.RepoNamespace, n.RepoName)
	_, _, err := s.Client(cu).
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color: colorPurple,
			Blocks: slack.Blocks{
				BlockSet: []slack.Block{
					slack.SectionBlock{
						Type: slack.MBTSection,
						Text: &slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: "*Approval Requested*",
						},
					},
					slack.SectionBlock{
						Type: slack.MBTSection,
						Text: &slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: fmt.Sprintf("*%s* - %s has requested the approval for the deployment(#%d).", fullName, n.DeploymentLogin, n.DeploymentNumber),
						},
					},
				},
			},
		}))

	return err
}

func (s *Slack) notifyApprovalResponded(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	fullName := fmt.Sprintf("%s/%s", n.RepoNamespace, n.RepoName)

	// Verb used in the message.
	var action string
	if n.ApprovalStatus == string(approval.StatusApproved) {
		action = "approved"
	} else {
		action = "declined"
	}

	_, _, err := s.Client(cu).
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color: colorPurple,
			Blocks: slack.Blocks{
				BlockSet: []slack.Block{
					slack.SectionBlock{
						Type: slack.MBTSection,
						Text: &slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: "*Approval Requested*",
						},
					},
					slack.SectionBlock{
						Type: slack.MBTSection,
						Text: &slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: fmt.Sprintf("*%s* - %s has *%s* the deployment(#%d.)", fullName, n.ApprovalLogin, action, n.DeploymentNumber),
						},
					},
				},
			},
		}))

	return err
}

func mapDeploymentStatusToColor(status string) string {
	switch deployment.Status(status) {
	case deployment.StatusWaiting:
		return colorGray
	case deployment.StatusCreated:
		return colorPurple
	case deployment.StatusRunning:
		return colorPurple
	case deployment.StatusSuccess:
		return colorGreen
	case deployment.StatusFailure:
		return colorRed
	default:
		return colorGray
	}
}
