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
		PostMessageContext(ctx, cu.ID, slack.MsgOptionAttachments(slack.Attachment{
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
							Text: fmt.Sprintf("*%s* - *%s* deploy `%s` to *%s* environment. (status: `%s`)", fullname, u.Login, ref, d.Env, d.Status),
						},
					},
				},
			},
		}))

	return err
}

func mapDeploymentStatusToColor(d *ent.Deployment) string {
	const (
		gray   = "#bfbfbf"
		purple = "#722ed1"
		green  = "#52c41a"
		red    = "#f5222d"
	)
	switch d.Status {
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

func refstr(d *ent.Deployment) string {
	if d.Type == deployment.TypeCommit {
		return d.Ref[:7]
	}

	return d.Ref
}
