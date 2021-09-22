package slack

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/approval"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/event"
)

const (
	colorGray   = "#bfbfbf"
	colorPurple = "#722ed1"
	colorGreen  = "#52c41a"
	colorRed    = "#f5222d"
)

func (s *Slack) Notify(ctx context.Context, e *ent.Event) {
	if e.Kind == event.KindDeployment {
		s.notifyDeploymentEvent(ctx, e)
	}

	if e.Kind == event.KindApproval {
		s.notifyApprovalEvent(ctx, e)
	}
}

func (s *Slack) notifyDeploymentEvent(ctx context.Context, e *ent.Event) {
	if err := e.CheckEagerLoading(); err != nil {
		s.log.Error("The eager loading of event has failed.")
		return
	}

	d := e.Edges.Deployment
	if err := d.CheckEagerLoading(); err != nil {
		s.log.Error("The eager loading of deployment has failed.")
		return
	}

	owner, err := s.i.FindUserByID(ctx, d.Edges.User.ID)
	if err != nil {
		s.log.Error("It has failed to find the owner of the deployment.", zap.Error(err))
		return
	}
	if owner.Edges.ChatUser == nil {
		s.log.Debug("Skip the notification. The owner is not connected with Slack.")
		return
	}

	// Build the message and post it.
	var option slack.MsgOption

	if e.Type == event.TypeCreated {
		option = slack.MsgOptionAttachments(slack.Attachment{
			Color:   mapDeploymentStatusToColor(d.Status),
			Pretext: fmt.Sprintf("*New Deployment #%d*", d.Number),
			Text:    fmt.Sprintf("*%s* deploys `%s` to the `%s` environment of `%s`. <%s|• View Details> ", owner.Login, d.GetShortRef(), d.Env, d.Edges.Repo.GetFullName(), s.buildDeploymentLink(d.Edges.Repo, d)),
		})
	} else if e.Type == event.TypeUpdated {
		option = slack.MsgOptionAttachments(slack.Attachment{
			Color:   mapDeploymentStatusToColor(d.Status),
			Pretext: fmt.Sprintf("*Deployment Updated #%d*", d.Number),
			Text:    fmt.Sprintf("The deployment <%s|#%d> of `%s` is updated `%s`.", s.buildDeploymentLink(d.Edges.Repo, d), d.Number, d.Edges.Repo.GetFullName(), d.Status),
		})
	}

	if _, _, err := slack.
		New(owner.Edges.ChatUser.BotToken).
		PostMessageContext(ctx, owner.Edges.ChatUser.ID, option); err != nil {
		s.log.Error("It has failed to post the message.", zap.Error(err))
	}
}

func (s *Slack) notifyApprovalEvent(ctx context.Context, e *ent.Event) {
	if err := e.CheckEagerLoading(); err != nil {
		s.log.Error("The eager loading of event has failed.")
		return
	}

	a := e.Edges.Approval
	if err := a.CheckEagerLoading(); err != nil {
		s.log.Error("The eager loading of approval has failed.")
		return
	}

	d := e.Edges.Deployment
	if err := d.CheckEagerLoading(); err != nil {
		s.log.Error("The eager loading of deployment has failed.")
		return
	}

	if e.Type == event.TypeCreated {
		option := slack.MsgOptionAttachments(slack.Attachment{
			Color:   colorPurple,
			Pretext: "*Approval Requested*",
			Text:    fmt.Sprintf("%s has requested the approval for the deployment <%s|#%d> of `%s`.", d.Edges.User.Login, s.buildDeploymentLink(d.Edges.Repo, d), d.Number, d.Edges.Repo.GetFullName()),
		})

		recipient, err := s.i.FindUserByID(ctx, a.Edges.User.ID)
		if err != nil {
			s.log.Error("It has failed to find the recipient of the approval.", zap.Error(err))
			return
		}
		if recipient.Edges.ChatUser == nil {
			s.log.Debug("Skip the notification. The recipient is not connected with Slack.")
			return
		}

		if _, _, err := slack.
			New(recipient.Edges.ChatUser.BotToken).
			PostMessageContext(ctx, recipient.Edges.ChatUser.ID, option); err != nil {
			s.log.Error("It has failed to post the message.", zap.Error(err))
		}
	}

	if e.Type == event.TypeUpdated {
		option := slack.MsgOptionAttachments(slack.Attachment{
			Color:   mapApprovalStatusToColor(a.Status),
			Pretext: "*Approval Responded*",
			Text:    fmt.Sprintf("%s has *%s* for the deployment <%s|#%d> of `%s`.", a.Edges.User.Login, a.Status, s.buildDeploymentLink(d.Edges.Repo, d), d.Number, d.Edges.Repo.GetFullName()),
		})

		requester, err := s.i.FindUserByID(ctx, d.Edges.User.ID)
		if err != nil {
			s.log.Error("It has failed to find the requester of the approval.", zap.Error(err))
			return
		}
		if requester.Edges.ChatUser == nil {
			s.log.Debug("Skip the notification. The requester is not connected with Slack.")
			return
		}

		if _, _, err := slack.
			New(requester.Edges.ChatUser.BotToken).
			PostMessageContext(ctx, requester.Edges.ChatUser.ID, option); err != nil {
			s.log.Error("It has failed to post the message.", zap.Error(err))
		}
	}
}

func (s *Slack) buildDeploymentLink(r *ent.Repo, d *ent.Deployment) string {
	return fmt.Sprintf("%s://%s/%s/deployments/%d", s.proto, s.host, r.GetFullName(), d.Number)
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
