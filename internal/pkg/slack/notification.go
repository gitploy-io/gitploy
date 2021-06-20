package slack

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
)

func (s *Slack) NotifyDeployment(ctx context.Context, cu *ent.ChatUser, d *ent.Deployment) error {
	const title = "New Deployment"

	if d.Edges.User == nil || d.Edges.Repo == nil {
		return fmt.Errorf("Edges are not found - user, repo.")
	}

	u, r := d.Edges.User, d.Edges.Repo

	fullname := fmt.Sprintf("%s/%s", r.Namespace, r.Name)
	ref := refstr(d)

	_, _, err := s.Client(cu).
		PostMessageContext(ctx, cu.UserID, slack.MsgOptionAttachments(slack.Attachment{
			Color: mapDeploymentStatusToColor(d),
			Blocks: slack.Blocks{
				BlockSet: []slack.Block{
					slack.SectionBlock{
						Type: slack.MBTSection,
						Text: &slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: fmt.Sprintf("*%s*", title),
						},
					},
					slack.SectionBlock{
						Type: slack.MBTSection,
						Text: &slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: fmt.Sprintf("*%s* - *%s* deploy `%s` to *%s* environment.", fullname, u.Login, ref, d.Env),
						},
					},
				},
			},
		}))

	return err
}

func mapDeploymentStatusToColor(d *ent.Deployment) string {
	switch d.Status {
	case deployment.StatusCreated:
		return "good"
	case deployment.StatusFailure:
		return "danger"
	default:
		return "good"
	}
}

func refstr(d *ent.Deployment) string {
	if d.Type == deployment.TypeCommit {
		return d.Ref[:7]
	}

	return d.Ref
}
