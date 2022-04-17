package slack

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/model/ent/review"
)

func (s *Slack) Notify(ctx context.Context, e *ent.Event) {
	if e.Type == event.TypeDeleted {
		s.log.Debug("Skip the deleted type event.")
		return
	}

	switch e.Kind {
	case event.KindDeploymentStatus:
		ds, err := s.i.FindDeploymentStatusByID(ctx, e.DeploymentStatusID)
		if err != nil {
			s.log.Error("Failed to find the deployment status.", zap.Error(err))
			break
		}

		s.notifyDeploymentStatusEvent(ctx, ds)
	case event.KindReview:
		r, err := s.i.FindReviewByID(ctx, e.ReviewID)
		if err != nil {
			s.log.Error("Failed to find the review.", zap.Error(err))
			break
		}

		s.notifyReviewEvent(ctx, r)
	}
}

func (s *Slack) notifyDeploymentStatusEvent(ctx context.Context, ds *ent.DeploymentStatus) {
	d, err := s.i.FindDeploymentByID(ctx, ds.DeploymentID)
	if err != nil {
		s.log.Error("Failed to find the deployment.", zap.Error(err))
		return
	}

	owner, err := s.i.FindUserByID(ctx, d.UserID)
	if err != nil {
		s.log.Error("Failed to find the owner of the deployment.", zap.Error(err))
		return
	} else if owner.Edges.ChatUser == nil {
		s.log.Debug("Skip the notification. The owner is not connected with Slack.")
		return
	}

	// Build the message and post it.
	options := []slack.MsgOption{
		slack.MsgOptionBlocks(
			slack.NewSectionBlock(
				slack.NewTextBlockObject(
					slack.MarkdownType,
					fmt.Sprintf("*%s* *#%d* %s", d.Edges.Repo.GetFullName(), d.Number, buildLink(s.buildDeploymentLink(d.Edges.Repo, d), " • View Details")),
					false, false,
				),
				[]*slack.TextBlockObject{
					slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*Status:*\n`%s`", ds.Status), false, false),
					slack.NewTextBlockObject(
						slack.MarkdownType,
						fmt.Sprintf("*Description:*\n>%s %s", ds.Description, buildLink(ds.LogURL, " • View Log")),
						false, false,
					),
				}, nil,
			),
		),
	}

	if _, _, err := slack.New(owner.Edges.ChatUser.BotToken).
		PostMessageContext(ctx, owner.Edges.ChatUser.ID, options...); err != nil {
		s.log.Error("Failed to post the message.", zap.Error(err))
	}
}

func (s *Slack) notifyReviewEvent(ctx context.Context, r *ent.Review) {
	d, err := s.i.FindDeploymentByID(ctx, r.DeploymentID)
	if err != nil {
		s.log.Error("Failed to find the deployment.", zap.Error(err))
		return
	}

	deployer := d.Edges.User
	if deployer == nil {
		s.log.Error("Failed to find the deployer.", zap.Error(err))
	}

	reviewer, err := s.i.FindUserByID(ctx, r.UserID)
	if err != nil {
		s.log.Error("Failed to find the reviewer.", zap.Error(err))
		return
	}

	switch r.Status {
	case review.StatusPending:
		option := slack.MsgOptionBlocks(
			slack.NewSectionBlock(
				slack.NewTextBlockObject(
					slack.MarkdownType,
					fmt.Sprintf("%s requested a review in `%s` *#%d* %s", deployer.Login, d.Edges.Repo.GetFullName(), d.Number, buildLink(s.buildDeploymentLink(d.Edges.Repo, d), " • View Details")),
					false, false,
				),
				nil, nil,
			),
		)

		if reviewer.Edges.ChatUser == nil {
			s.log.Debug("Skip the notification. The reviewer is not connected with Slack.")
			return
		}

		if _, _, err := slack.New(reviewer.Edges.ChatUser.BotToken).
			PostMessageContext(ctx, reviewer.Edges.ChatUser.ID, option); err != nil {
			s.log.Error("Failed to post the message.", zap.Error(err))
		}

	default:
		option := slack.MsgOptionBlocks(
			slack.NewSectionBlock(
				slack.NewTextBlockObject(
					slack.MarkdownType,
					fmt.Sprintf("%s %s in `%s` *#%d* %s", reviewer.Login, r.Status, d.Edges.Repo.GetFullName(), d.Number, buildLink(s.buildDeploymentLink(d.Edges.Repo, d), " • View Details")),
					false, false,
				),
				nil, nil,
			),
		)

		if deployer.Edges.ChatUser == nil {
			s.log.Debug("Skip the notification. The deployer is not connected with Slack.")
			return
		}

		if _, _, err := slack.New(deployer.Edges.ChatUser.BotToken).
			PostMessageContext(ctx, deployer.Edges.ChatUser.ID, option); err != nil {
			s.log.Error("Failed to post the message.", zap.Error(err))
		}
	}
}

func (s *Slack) buildDeploymentLink(r *ent.Repo, d *ent.Deployment) string {
	return fmt.Sprintf("%s://%s/%s/deployments/%d", s.proto, s.host, r.GetFullName(), d.Number)
}

func buildLink(link string, msg string) string {
	if msg == "" || link == "" {
		return ""
	}

	return fmt.Sprintf("<%s|%s>", link, msg)
}
