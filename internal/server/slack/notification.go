package slack

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/event"
	"github.com/gitploy-io/gitploy/ent/review"
)

const (
	colorGray   = "#bfbfbf"
	colorPurple = "#722ed1"
	colorGreen  = "#52c41a"
	colorRed    = "#f5222d"
)

func (s *Slack) Notify(ctx context.Context, e *ent.Event) {
	if e.Type == event.TypeDeleted {
		s.log.Debug("Skip the deleted type event.")
		return
	}

	if e.Kind == event.KindDeployment {
		s.notifyDeploymentEvent(ctx, e)
	}

	if e.Kind == event.KindReview {
		s.notifyReviewEvent(ctx, e)
	}
}

func (s *Slack) notifyDeploymentEvent(ctx context.Context, e *ent.Event) {
	var d *ent.Deployment

	if d = e.Edges.Deployment; d == nil {
		s.log.Error("The eager loading of event has failed.")
		return
	}

	if d.Edges.User == nil || d.Edges.Repo == nil {
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

func (s *Slack) notifyReviewEvent(ctx context.Context, e *ent.Event) {
	var (
		r *ent.Review
		d *ent.Deployment
	)

	if r = e.Edges.Review; r == nil {
		s.log.Error("The eager loading of review has failed.")
		return
	}

	if d = r.Edges.Deployment; d == nil {
		s.log.Error("The eager loading of deployment has failed.")
		return
	}

	if e.Type == event.TypeCreated {
		option := slack.MsgOptionAttachments(slack.Attachment{
			Color:   colorPurple,
			Pretext: "*Review Requested*",
			Text:    fmt.Sprintf("%s requested the review for the deployment <%s|#%d> of `%s`.", d.Edges.User.Login, s.buildDeploymentLink(d.Edges.Repo, d), d.Number, d.Edges.Repo.GetFullName()),
		})

		recipient, err := s.i.FindUserByID(ctx, r.Edges.User.ID)
		if err != nil {
			s.log.Error("It has failed to find the recipient of the review.", zap.Error(err))
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
			Color:   mapReviewStatusToColor(r.Status),
			Pretext: "*Review Responded*",
			Text:    fmt.Sprintf("%s *%s* the deployment <%s|#%d> of `%s`.", r.Edges.User.Login, r.Status, s.buildDeploymentLink(d.Edges.Repo, d), d.Number, d.Edges.Repo.GetFullName()),
		})

		requester, err := s.i.FindUserByID(ctx, d.Edges.User.ID)
		if err != nil {
			s.log.Error("It has failed to find the requester of the review.", zap.Error(err))
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

func mapReviewStatusToColor(status review.Status) string {
	switch status {
	case review.StatusPending:
		return colorGray
	case review.StatusApproved:
		return colorGreen
	case review.StatusRejected:
		return colorRed
	default:
		return colorGray
	}
}
