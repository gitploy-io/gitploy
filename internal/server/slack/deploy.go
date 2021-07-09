package slack

import (
	"context"
	"fmt"
	"strings"

	"github.com/slack-go/slack"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/vo"
)

// handleDeployCmd handles deploy command: "/gitploy deploy OWNER/REPO".
// It opens a dialog with fields to submit the payload for deployment.
func (s *Slack) handleDeployCmd(ctx context.Context, cmd slack.SlashCommand) error {
	// user have to be exist if chat user is found.
	cu, err := s.i.FindChatUserWithUserByID(ctx, cmd.UserID)
	if ent.IsNotFound(err) {
		responseMessage(cmd, fmt.Sprint("Slack is not connected with Gitploy."))
		return nil
	} else if err != nil {
		return err
	}

	u := cu.Edges.User
	client := slack.New(cu.BotToken)

	fullname := trimDeployCommandPrefix(cmd.Text)
	ns, n, err := parseFullName(fullname)
	if err != nil {
		responseMessage(cmd, fmt.Sprintf("`%s` is invalid format.", fullname))
		return nil
	}

	r, err := s.i.FindRepoByNamespaceName(ctx, u, ns, n)
	if ent.IsNotFound(err) {
		responseMessage(cmd, fmt.Sprintf("The `%s` is not found.", fullname))
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to find the repo: %w", err)
	}

	config, err := s.i.GetConfig(ctx, u, r)
	if vo.IsConfigNotFoundError(err) {
		responseMessage(cmd, "The config file is not found.")
		return nil
	} else if vo.IsConfigParseError(err) {
		responseMessage(cmd, "The config file is invliad format.")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get the config: %w", err)
	}

	state := randstr()

	cb, err := s.i.CreateDeployChatCallback(ctx, cu, r, &ent.ChatCallback{
		Type:  chatcallback.TypeDeploy,
		State: state,
	})
	if err != nil {
		return fmt.Errorf("failed to save the callback: %w", err)
	}

	err = client.OpenDialogContext(ctx, cmd.TriggerID, slack.Dialog{
		CallbackID:     itoa(cb.ID),
		State:          state,
		Title:          "Deploy",
		SubmitLabel:    "Submit",
		NotifyOnCancel: true,
		Elements:       createDeployDialogElement(config),
	})
	if err != nil {
		return fmt.Errorf("failed to open a new dialog: %w", err)
	}

	return nil
}

// interactDeploy deploy with the submitted payload.
func (s *Slack) interactDeploy(ctx context.Context, scb slack.InteractionCallback, cb *ent.ChatCallback) error {
	var (
		typ = scb.Submission["type"]
		ref = scb.Submission["ref"]
		env = scb.Submission["env"]
	)

	if cb.Edges.ChatUser == nil {
		return fmt.Errorf("The chat_user edge is not found")
	}

	if cb.Edges.Repo == nil {
		return fmt.Errorf("The repo edge is not found")
	}

	cu := cb.Edges.ChatUser
	re := cb.Edges.Repo

	cu, _ = s.i.FindChatUserWithUserByID(ctx, cu.ID)
	u := cu.Edges.User

	cf, err := s.i.GetConfig(ctx, u, re)
	if vo.IsConfigNotFoundError(err) {
		responseMessage(scb, "The configuration file is not found.")
		return nil
	} else if vo.IsConfigParseError(err) {
		responseMessage(scb, "The configuration file is invalid format.")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get the config: %w", err)
	}

	if !cf.HasEnv(env) {
		responseMessage(scb, "The configuration file is invalid format.")
		return nil
	}

	// Prepare to deploy:
	// 1) commit SHA.
	// 2) next deployment number.
	var (
		sha    string
		number int
	)
	if sha, err = s.getCommitSha(ctx, u, re, typ, ref); err != nil {
		return fmt.Errorf("invalid ref: %w", err)
	}

	if number, err = s.i.GetNextDeploymentNumberOfRepo(ctx, re); err != nil {
		return fmt.Errorf("failed to get the next deployment number: %w", err)
	}

	d, err := s.i.Deploy(ctx, u, cb.Edges.Repo, &ent.Deployment{
		Number: number,
		Type:   deployment.Type(typ),
		Ref:    ref,
		Sha:    sha,
		Env:    env,
	}, cf.GetEnv(env))
	if err != nil {
		return fmt.Errorf("failed to deploy: %w", err)
	}

	s.i.Publish(ctx, d)
	return nil
}

func sendResponse(c *slack.Client, cmd slack.SlashCommand, message string) error {
	_, _, _, err := c.SendMessage(cmd.ChannelID, slack.MsgOptionText(message, true), slack.MsgOptionResponseURL(cmd.ResponseURL, "ephemeral"))
	return err
}

func trimDeployCommandPrefix(txt string) string {
	return strings.TrimPrefix(txt, "deploy ")
}

func parseFullName(n string) (string, string, error) {
	namespaceName := strings.Split(n, "/")
	if len(namespaceName) != 2 {
		return "", "", fmt.Errorf("invalid format")
	}

	return namespaceName[0], namespaceName[1], nil
}

func createDeployDialogElement(c *vo.Config) []slack.DialogElement {
	options := []slack.DialogSelectOption{}
	for _, env := range c.Envs {
		options = append(options, slack.DialogSelectOption{
			Label: strings.Title(env.Name),
			Value: env.Name,
		})
	}

	return []slack.DialogElement{
		slack.DialogInputSelect{
			DialogInput: slack.DialogInput{
				Type:  slack.InputTypeSelect,
				Label: "Environment",
				Name:  "env",
			},
			Options: options,
		},
		slack.DialogInputSelect{
			DialogInput: slack.DialogInput{
				Type:  slack.InputTypeSelect,
				Label: "Type",
				Name:  "type",
			},
			Options: []slack.DialogSelectOption{
				{
					Label: "Commit",
					Value: "commit",
				},
				{
					Label: "Branch",
					Value: "branch",
				},
				{
					Label: "Tag",
					Value: "tag",
				},
			},
		},
		slack.DialogInput{
			Type:  "text",
			Label: "Ref",
			Name:  "ref",
			Hint:  "E.g. Commit - 25a667d6, Branch - main, Tag - v0.1.2",
		},
	}
}

func (s *Slack) getCommitSha(ctx context.Context, u *ent.User, re *ent.Repo, typ, ref string) (string, error) {
	if typ == "commit" {
		return ref, nil
	} else if typ == "branch" {
		b, err := s.i.GetBranch(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return b.CommitSha, nil
	} else if typ == "tag" {
		t, err := s.i.GetTag(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return t.CommitSha, nil
	}

	return "", fmt.Errorf("Type must be one of commit, branch, tag.")
}
