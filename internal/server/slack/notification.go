package slack

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/vo"
)

const (
	colorGray   = "#bfbfbf"
	colorPurple = "#722ed1"
	colorGreen  = "#52c41a"
	colorRed    = "#f5222d"
)

func (s *Slack) Notify(ctx context.Context, e *ent.Event) error {
	if s.i.CheckNotificationRecordOfEvent(ctx, e) {
		return nil
	}

	var (
		n     *vo.Notification
		users []*ent.User
		err   error
	)

	if n, err = s.i.ConvertEventToNotification(ctx, e); err != nil {
		return err
	}

	if users, err = s.i.ListUsersOfEvent(ctx, e); err != nil {
		return err
	}

	for _, u := range users {
		// Load edges.
		if u, err = s.i.FindUserByID(ctx, u.ID); err != nil {
			s.log.Error("It has failed to get the user.", zap.Error(err))
			continue
		}

		if u.Edges.ChatUser != nil {
			if err := s.notify(ctx, u.Edges.ChatUser, n); err != nil {
				s.log.Error("It has failed to notify the event.", zap.Error(err))
			}
		}
	}

	return nil
}

func (s *Slack) notify(ctx context.Context, cu *ent.ChatUser, n *vo.Notification) error {
	if n.Kind == vo.KindDeployment && n.Type == vo.TypeCreated {
		return s.notifyDeploymentCreated(ctx, cu, n)
	}

	if n.Kind == vo.KindDeployment && n.Type == vo.TypeUpdated {
		return s.notifyDeploymentUpdated(ctx, cu, n)
	}

	if n.Kind == vo.KindApproval && n.Type == vo.TypeCreated {
		return s.notifyApprovalRequested(ctx, cu, n)
	}

	if n.Kind == vo.KindApproval && n.Type == vo.TypeUpdated {
		return s.notifyApprovalResponded(ctx, cu, n)
	}

	return fmt.Errorf("It is out of cases - kind: %s, type: %s.", n.Kind, n.Type)
}

func (s *Slack) notifyDeploymentCreated(ctx context.Context, cu *ent.ChatUser, n *vo.Notification) error {
	var (
		repoName = fmt.Sprintf("%s/%s", n.Repo.Namespace, n.Repo.Name)
		number   = n.Deployment.Number
		ref      = n.Deployment.Ref
		env      = n.Deployment.Env
		deployer = n.Deployment.Login
		link     = s.buildLink(n)
	)

	if n.Deployment.Type == string(deployment.TypeCommit) && len(n.Deployment.Ref) > 7 {
		ref = n.Deployment.Ref[:7]
	}

	client := slack.New(cu.BotToken)
	_, _, err := client.
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color:   mapDeploymentStatusToColor(deployment.Status(n.Deployment.Status)),
			Pretext: fmt.Sprintf("*New Deployment #%d*", number),
			Text:    fmt.Sprintf("*%s* deploys `%s` to the `%s` environment of `%s`. <%s|• View Details> ", deployer, ref, env, repoName, link),
		}))

	return err
}

func (s *Slack) notifyDeploymentUpdated(ctx context.Context, cu *ent.ChatUser, n *vo.Notification) error {
	var (
		repoName = fmt.Sprintf("%s/%s", n.Repo.Namespace, n.Repo.Name)
		number   = n.Deployment.Number
		status   = n.Deployment.Status
		link     = s.buildLink(n)
	)

	client := slack.New(cu.BotToken)
	_, _, err := client.
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color:   mapDeploymentStatusToColor(deployment.Status(n.Deployment.Status)),
			Pretext: fmt.Sprintf("*Deployment Updated #%d*", number),
			Text:    fmt.Sprintf("The deployment <%s|#%d> of `%s` is updated %s.", link, number, repoName, status),
		}))

	return err
}

func (s *Slack) notifyApprovalRequested(ctx context.Context, cu *ent.ChatUser, n *vo.Notification) error {
	var (
		repoName = fmt.Sprintf("%s/%s", n.Repo.Namespace, n.Repo.Name)
		number   = n.Deployment.Number
		approver = n.Approval.Login
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

func (s *Slack) notifyApprovalResponded(ctx context.Context, cu *ent.ChatUser, n *vo.Notification) error {
	var (
		repoName = fmt.Sprintf("%s/%s", n.Repo.Namespace, n.Repo.Name)
		number   = n.Deployment.Number
		action   = string(n.Approval.Status)
		approver = n.Approval.Login
		link     = s.buildLink(n)
	)

	client := slack.New(cu.BotToken)

	_, _, err := client.
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
			Color:   mapApprovalStatusToColor(approval.Status(n.Approval.Status)),
			Pretext: "*Approval Responded*",
			Text:    fmt.Sprintf("%s has *%s* for the deployment <%s|#%d> of `%s`.", approver, action, link, number, repoName),
		}))

	return err
}

func (s *Slack) buildLink(n *vo.Notification) string {
	return fmt.Sprintf("%s://%s/%s/%s/deployments/%d", s.proto, s.host, n.Repo.Namespace, n.Repo.Name, n.Deployment.Number)
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
