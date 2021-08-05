package slack

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/notification"
	"github.com/hanjunlee/gitploy/vo"
)

const (
	// When creating a view, set unique block_ids for all blocks
	// and unique action_ids for each block element.
	blockEnv        = "block_env"
	blockType       = "block_type"
	blockRef        = "block_ref"
	blockApprovers  = "block_approvers"
	actionEnv       = "action_env"
	actionType      = "action_type"
	actionRef       = "action_ref"
	actionApprovers = "action_approver_ids"
)

type (
	deployViewSubmission struct {
		Env         string
		Type        string
		Ref         string
		ApproverIDs []string
	}

	ErrorsPayload struct {
		ResponseAction string            `json:"response_action"`
		Errors         map[string]string `json:"errors"`
	}
)

// handleDeployCmd handles deploy command: "/gitploy deploy OWNER/REPO".
// It opens a dialog with fields to submit the payload for deployment.
func (s *Slack) handleDeployCmd(c *gin.Context) {
	ctx := c.Request.Context()

	// SlashCommandParse hvae to be success because
	// it was called in the Cmd method.
	cmd, _ := slack.SlashCommandParse(c.Request)

	cu, err := s.i.FindChatUserByID(ctx, cmd.UserID)
	if ent.IsNotFound(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "Slack is not connected with Gitploy.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get chat user.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	args := strings.Split(cmd.Text, " ")

	// The length of args is always equal to two.
	ns, n, err := parseFullName(args[1])
	if err != nil {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("`%s` is invalid repository format.", args[1]))
		c.Status(http.StatusOK)
		return
	}

	r, err := s.i.FindRepoOfUserByNamespaceName(ctx, cu.Edges.User, ns, n)
	if ent.IsNotFound(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s` repository is not found.", args[1]))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	config, err := s.i.GetConfig(ctx, cu.Edges.User, r)
	if vo.IsConfigNotFoundError(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "The config file is not found.")
		c.Status(http.StatusOK)
		return
	} else if vo.IsConfigParseError(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "The config file is invliad format.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the config.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	perms, err := s.i.ListPermsOfRepo(ctx, r, "", 1, 100)
	if err != nil {
		s.log.Error("It has failed to list permissions.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	// Create a new callback to interact with submissions.
	cb, err := s.i.CreateChatCallback(ctx, cu, r, &ent.ChatCallback{
		Type: chatcallback.TypeDeploy,
	})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = slack.New(cu.BotToken).
		OpenViewContext(ctx, cmd.TriggerID, buildDeployView(cb.Hash, config, perms))
	if err != nil {
		s.log.Error("It has failed to open a new view.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func parseFullName(fullname string) (string, string, error) {
	namespaceName := strings.Split(fullname, "/")
	if len(namespaceName) != 2 {
		return "", "", fmt.Errorf("It is a invalid formatted command.")
	}

	return namespaceName[0], namespaceName[1], nil
}

func buildDeployView(callbackID string, c *vo.Config, perms []*ent.Perm) slack.ModalViewRequest {
	envs := []*slack.OptionBlockObject{}
	for _, env := range c.Envs {
		envs = append(envs, &slack.OptionBlockObject{
			Text: &slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: env.Name,
			},
			Value: env.Name,
		})
	}

	approvers := []*slack.OptionBlockObject{}
	for _, perm := range perms {
		u := perm.Edges.User
		if u == nil {
			continue
		}

		approvers = append(approvers, &slack.OptionBlockObject{
			Text: &slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: u.Login,
			},
			Value: u.ID,
		})
	}

	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: callbackID,
		Title: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Deploy",
		},
		Submit: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Submit",
		},
		Close: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Close",
		},
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: blockEnv,
					Label: &slack.TextBlockObject{
						Type: slack.PlainTextType,
						Text: "Environment",
					},
					Element: slack.SelectBlockElement{
						Type:     slack.OptTypeStatic,
						ActionID: actionEnv,
						Placeholder: &slack.TextBlockObject{
							Type: slack.PlainTextType,
							Text: "Select target environment",
						},
						Options: envs,
					},
				},
				slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: blockType,
					Label: &slack.TextBlockObject{
						Type: slack.PlainTextType,
						Text: "Reference Type",
					},
					Element: slack.SelectBlockElement{
						Type:     slack.OptTypeStatic,
						ActionID: actionType,
						Placeholder: &slack.TextBlockObject{
							Type: slack.PlainTextType,
							Text: "Select your reference type",
						},
						Options: []*slack.OptionBlockObject{
							{
								Text: &slack.TextBlockObject{
									Type: slack.PlainTextType,
									Text: "Commit",
								},
								Value: "commit",
							},
							{
								Text: &slack.TextBlockObject{
									Type: slack.PlainTextType,
									Text: "Branch",
								},
								Value: "branch",
							},
							{
								Text: &slack.TextBlockObject{
									Type: slack.PlainTextType,
									Text: "Tag",
								},
								Value: "tag",
							},
						},
					},
				},
				slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: blockRef,
					Label: &slack.TextBlockObject{
						Type: slack.PlainTextType,
						Text: "Reference",
					},
					Element: slack.PlainTextInputBlockElement{
						Type:     slack.METPlainTextInput,
						ActionID: actionRef,
						Placeholder: &slack.TextBlockObject{
							Type: slack.PlainTextType,
							Text: "E.g. Commit - 25a667d6, Branch - main, Tag - v0.1.2",
						},
					},
				},
				slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: blockApprovers,
					Label: &slack.TextBlockObject{
						Type: slack.PlainTextType,
						Text: "Approvers",
					},
					Optional: true,
					Element: slack.SelectBlockElement{
						Type:     slack.MultiOptTypeStatic,
						ActionID: actionApprovers,
						Placeholder: &slack.TextBlockObject{
							Type: slack.PlainTextType,
							Text: "Select approvers",
						},
						Options: approvers,
					},
				},
			},
		},
	}
}

// interactDeploy deploy with the submitted payload.
func (s *Slack) interactDeploy(c *gin.Context) {
	ctx := c.Request.Context()

	// InteractionCallbackParse always to be parsed successfully because
	// it was called in the Interact method.
	itr, _ := s.InteractionCallbackParse(c.Request)
	cb, _ := s.i.FindChatCallbackByHash(ctx, itr.View.CallbackID)

	cu, _ := s.i.FindChatUserByID(ctx, cb.Edges.ChatUser.ID)

	// Parse view submission, and
	// validate values.
	sm := parseViewSubmissions(itr)

	_, err := s.getCommitSha(ctx, cu.Edges.User, cb.Edges.Repo, sm.Type, sm.Ref)
	if vo.IsRefNotFoundError(err) {
		c.JSON(http.StatusOK, buildErrorsPayload(map[string]string{
			blockRef: "The reference is not found.",
		}))
		return
	} else if err != nil {
		s.log.Error("It has failed to get the SHA of commit.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cf, err := s.i.GetConfig(ctx, cu.Edges.User, cb.Edges.Repo)
	if err != nil {
		s.log.Error("It has failed to get the config file.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if !cf.HasEnv(sm.Env) {
		s.log.Error("It has failed to find the environment.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	env := cf.GetEnv(sm.Env)

	number, err := s.i.GetNextDeploymentNumberOfRepo(ctx, cb.Edges.Repo)
	if err != nil {
		s.log.Error("It has failed to get the next deployment number.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	d, err := s.i.Deploy(ctx, cu.Edges.User, cb.Edges.Repo, &ent.Deployment{
		Number: number,
		Type:   deployment.Type(sm.Type),
		Ref:    sm.Ref,
		Env:    sm.Env,
	}, env)
	if err != nil {
		s.log.Error("It has failed to deploy.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if err = s.i.Publish(ctx, notification.TypeDeploymentCreated, cb.Edges.Repo, d, nil); err != nil {
		s.log.Warn("It has failed to publish the deployment.", zap.Error(err))
	}

	if env.IsApprovalEabled() {
		for _, id := range sm.ApproverIDs {
			a, err := s.i.CreateApproval(ctx, &ent.Approval{
				UserID:       id,
				DeploymentID: d.ID,
			})
			if err != nil {
				s.log.Error("It has failed to create a new approval.", zap.Error(err))
				continue
			}

			if err := s.i.Publish(ctx, notification.TypeApprovalRequested, cb.Edges.Repo, d, a); err != nil {
				s.log.Warn("It has failed to publish the approval.", zap.Error(err))
			}
		}
	}

	c.Status(http.StatusOK)
}

func parseViewSubmissions(itr slack.InteractionCallback) *deployViewSubmission {
	sm := &deployViewSubmission{}

	values := itr.View.State.Values
	if v, ok := values[blockEnv][actionEnv]; ok {
		sm.Env = v.SelectedOption.Value
	}

	if v, ok := values[blockType][actionType]; ok {
		sm.Type = v.SelectedOption.Value
	}

	if v, ok := values[blockRef][actionRef]; ok {
		sm.Ref = v.Value
	}

	ids := make([]string, 0)
	if v, ok := values[blockApprovers][actionApprovers]; ok {
		for _, option := range v.SelectedOptions {
			ids = append(ids, option.Value)
		}

		sm.ApproverIDs = ids
	}

	return sm
}

func buildErrorsPayload(errors map[string]string) *ErrorsPayload {
	return &ErrorsPayload{
		ResponseAction: "errors",
		Errors:         errors,
	}
}

func (s *Slack) getCommitSha(ctx context.Context, u *ent.User, re *ent.Repo, typ, ref string) (string, error) {
	switch typ {
	case "commit":
		c, err := s.i.GetCommit(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return c.Sha, nil
	case "branch":
		b, err := s.i.GetBranch(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return b.CommitSha, nil
	case "tag":
		t, err := s.i.GetTag(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return t.CommitSha, nil
	default:
		return "", fmt.Errorf("Type must be one of commit, branch, tag.")
	}
}
