package slack

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/nleeper/goment"
	"github.com/slack-go/slack"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	deploymentAggregation struct {
		envName     string
		deployments []*ent.Deployment
	}
)

// handleRollbackCmd handles rollback command: "/gitploy rollback OWNER/REPO".
func (s *Slack) handleRollbackCmd(ctx context.Context, cmd slack.SlashCommand) error {
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

	fullname := trimRollbackCommandPrefix(cmd.Text)
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
		responseMessage(cmd, "The configuration file is not found")
		return nil
	} else if vo.IsConfigParseError(err) {
		responseMessage(cmd, "The configuration file is invliad format.")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get the config: %w", err)
	}

	state := randstr()
	as := s.getSucceedDeploymentAggregation(ctx, r, config)

	cb, err := s.i.CreateDeployChatCallback(ctx, cu, r, &ent.ChatCallback{
		Type:  chatcallback.TypeRollback,
		State: state,
	})
	if err != nil {
		return fmt.Errorf("failed to save the callback: %w", err)
	}

	err = client.OpenDialogContext(ctx, cmd.TriggerID, slack.Dialog{
		CallbackID:     itoa(cb.ID),
		State:          state,
		Title:          "Rollback",
		SubmitLabel:    "Submit",
		NotifyOnCancel: true,
		Elements:       createRollbackDialogElement(as),
	})
	if err != nil {
		return fmt.Errorf("failed to open a new dialog: %w", err)
	}

	return nil
}

func (s *Slack) getSucceedDeploymentAggregation(ctx context.Context, r *ent.Repo, cf *vo.Config) []*deploymentAggregation {
	a := []*deploymentAggregation{}

	for _, env := range cf.Envs {
		ds, _ := s.i.ListDeployments(ctx, r, env.Name, string(deployment.StatusSuccess), 1, 5)
		if len(ds) == 0 {
			continue
		}

		a = append(a, &deploymentAggregation{
			envName:     env.Name,
			deployments: ds,
		})
	}

	return a
}

func (s *Slack) interactRollback(ctx context.Context, scb slack.InteractionCallback, cb *ent.ChatCallback) error {
	var (
		id = scb.Submission["deployment_id"]
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

	d, err := s.i.FindDeploymentWithEdgesByID(ctx, atoi(id))
	if err != nil {
		return fmt.Errorf("failed to find the deployment: %w", err)
	}

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

	if !cf.HasEnv(d.Env) {
		responseMessage(scb, "The configuration file is invalid format.")
		return nil
	}

	d, err = s.i.Rollback(ctx, u, cb.Edges.Repo, &ent.Deployment{
		Type: deployment.Type(d.Type),
		Ref:  d.Ref,
		Env:  d.Env,
	}, cf.GetEnv(d.Env))
	if err != nil {
		return fmt.Errorf("failed to deploy: %w", err)
	}

	s.i.Publish(ctx, d)
	return nil
}

func createRollbackDialogElement(as []*deploymentAggregation) []slack.DialogElement {
	groups := []slack.DialogOptionGroup{}
	for _, a := range as {
		options := []slack.DialogSelectOption{}

		for _, d := range a.deployments {
			created, _ := goment.New(d.CreatedAt)
			options = append(options, slack.DialogSelectOption{
				Label: fmt.Sprintf("#%d - %s deployed at %s", d.ID, strRef(d), created.FromNow()),
				Value: strconv.Itoa(d.ID),
			})
		}

		groups = append(groups, slack.DialogOptionGroup{
			Label:   strings.Title(a.envName),
			Options: options,
		})
	}

	return []slack.DialogElement{
		slack.DialogInputSelect{
			DialogInput: slack.DialogInput{
				Type:  slack.InputTypeSelect,
				Label: "Deployment",
				Name:  "deployment_id",
			},
			OptionGroups: groups,
		},
	}
}

func strRef(d *ent.Deployment) string {
	ref := d.Ref

	if d.Type == deployment.TypeCommit {
		ref = d.Ref[:7]
	}

	return ref
}

func trimRollbackCommandPrefix(txt string) string {
	return strings.TrimPrefix(txt, "rollback ")
}
