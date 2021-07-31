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
	var (
		repoName = fmt.Sprintf("%s/%s", n.RepoNamespace, n.RepoName)
		number   = n.DeploymentNumber
		ref      = n.DeploymentRef
		env      = n.DeploymentEnv
		deployer = n.DeploymentLogin
		link     = s.buildLink(n)
	)

	if n.DeploymentType == string(deployment.TypeCommit) {
		ref = n.DeploymentRef[:7]
	}

	client := slack.New(cu.BotToken)
	_, _, err := client.
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color:   mapDeploymentStatusToColor(deployment.Status(n.DeploymentStatus)),
			Pretext: fmt.Sprintf("*New Deployment #%d*", number),
			Text:    fmt.Sprintf("*%s* deploys `%s` to the `%s` environment of `%s`. <%s|• View Details> ", deployer, ref, env, repoName, link),
		}))

	return err
}

func (s *Slack) notifyApprovalRequested(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	var (
		repoName = fmt.Sprintf("%s/%s", n.RepoNamespace, n.RepoName)
		number   = n.DeploymentNumber
		approver = n.ApprovalLogin
		link     = s.buildLink(n)
	)

	client := slack.New(cu.BotToken)

	_, _, err := client.
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color:   colorPurple,
			Pretext: "*Approval Requested*",
			Text:    fmt.Sprintf("%s has requested the approval for the deployment <%s|#%d> of `%s`.", approver, link, number, repoName),
		}))

	return err
}

func (s *Slack) notifyApprovalResponded(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	var (
		repoName = fmt.Sprintf("%s/%s", n.RepoNamespace, n.RepoName)
		number   = n.DeploymentNumber
		action   = string(n.ApprovalStatus)
		approver = n.ApprovalLogin
		link     = s.buildLink(n)
	)

	client := slack.New(cu.BotToken)

	_, _, err := client.
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color:   mapApprovalStatusToColor(approval.Status(n.ApprovalStatus)),
			Pretext: "*Approval Responded*",
			Text:    fmt.Sprintf("%s has *%s* for the deployment <%s|#%d> of `%s`.", approver, action, link, number, repoName),
		}))

	return err
}

func (s *Slack) buildLink(n *ent.Notification) string {
	return fmt.Sprintf("%s://%s/%s/%s/deployments/%d", s.proto, s.host, n.RepoNamespace, n.RepoName, n.DeploymentNumber)
}

func mapDeploymentStatusToColor(status deployment.Status) string {
	switch status {
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

func mapApprovalStatusToColor(status approval.Status) string {
	switch status {
	case approval.StatusPending:
		return colorGray
	case approval.StatusApproved:
		return colorGreen
	case approval.StatusDeclined:
		return colorRed
	default:
		return colorGray
	}
}
