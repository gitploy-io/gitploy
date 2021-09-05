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

func (s *Slack) Notify(ctx context.Context, e *ent.Event) error {
	if s.i.CheckNotificationRecordOfEvent(ctx, e) {
		return nil
	}

	var (
		users []*ent.User
		err   error
	)

	if users, err = s.i.ListUsersOfEvent(ctx, e); err != nil {
		return err
	}

	for _, u := range users {
		// Eager loading for chat user.
		if u, err = s.i.FindUserByID(ctx, u.ID); err != nil {
			s.log.Error("It has failed to eager load.", zap.Error(err))
			continue
		}
		if u.Edges.ChatUser == nil {
			continue
		}

		if err := s.notify(ctx, u.Edges.ChatUser, e); err != nil {
			s.log.Error("It has failed to notify the event.", zap.Error(err))
		}
	}

	return nil
}

func (s *Slack) notify(ctx context.Context, cu *ent.ChatUser, e *ent.Event) error {
	var option slack.MsgOption

	// Check the event has processed eager loading.
	if err := e.CheckEagerLoading(); err != nil {
		return err
	}

	if e.Kind == event.KindDeployment {
		d := e.Edges.Deployment
		if err := d.CheckEagerLoading(); err != nil {
			return err
		}

		var (
			r = d.Edges.Repo
			u = d.Edges.User
		)

		if e.Type == event.TypeCreated {
			option = slack.MsgOptionAttachments(
				slack.Attachment{
					Color:   mapDeploymentStatusToColor(d.Status),
					Pretext: fmt.Sprintf("*New Deployment #%d*", d.Number),
					Text:    fmt.Sprintf("*%s* deploys `%s` to the `%s` environment of `%s`. <%s|• View Details> ", u.Login, d.GetShortRef(), d.Env, r.GetFullName(), s.buildDeploymentLink(r, d)),
				},
			)
		} else if e.Type == event.TypeUpdated {
			option = slack.MsgOptionAttachments(
				slack.Attachment{
					Color:   mapDeploymentStatusToColor(d.Status),
					Pretext: fmt.Sprintf("*Deployment Updated #%d*", d.Number),
					Text:    fmt.Sprintf("The deployment <%s|#%d> of `%s` is updated `%s`.", s.buildDeploymentLink(r, d), d.Number, r.GetFullName(), d.Status),
				},
			)
		}
	} else if e.Kind == event.KindApproval {
		a := e.Edges.Approval
		if err := a.CheckEagerLoading(); err != nil {
			return err
		}

		var (
			u *ent.User       = a.Edges.User
			d *ent.Deployment = a.Edges.Deployment
			r *ent.Repo
		)

		// Approval have to process eager loading for the deployment, too.
		if err := d.CheckEagerLoading(); err != nil {
			return err
		}
		r = d.Edges.Repo

		if e.Type == event.TypeCreated {
			option = slack.MsgOptionAttachments(slack.Attachment{
				Color:   colorPurple,
				Pretext: "*Approval Requested*",
				Text:    fmt.Sprintf("%s has requested the approval for the deployment <%s|#%d> of `%s`.", u.Login, s.buildDeploymentLink(r, d), d.Number, r.GetFullName()),
			})
		} else if e.Type == event.TypeUpdated {
			option = slack.MsgOptionAttachments(slack.Attachment{
				Color:   mapApprovalStatusToColor(a.Status),
				Pretext: "*Approval Responded*",
				Text:    fmt.Sprintf("%s has *%s* for the deployment <%s|#%d> of `%s`.", u.Login, a.Status, s.buildDeploymentLink(r, d), d.Number, r.GetFullName()),
			})
		}
	}

	_, _, err := slack.
		New(cu.BotToken).
		PostMessageContext(ctx, cu.ID, option)
	return err
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
