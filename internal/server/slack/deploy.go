package slack

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/callback"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/event"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/vo"
)

const (
// linkUnprocessalbeEntity = "https://github.com/gitploy-io/gitploy/discussions/64"
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

	av, _ := c.Get(KeyCmd)
	cmd := av.(slack.SlashCommand)

	bv, _ := c.Get(KeyChatUser)
	cu := bv.(*ent.ChatUser)

	s.log.Debug("Processing deploy command.", zap.String("command", cmd.Text))
	ns, n := parseCmd(cmd.Text)

	r, err := s.i.FindRepoOfUserByNamespaceName(ctx, cu.Edges.User, ns, n)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s/%s` repository is not found.", ns, n))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	config, err := s.i.GetConfig(ctx, cu.Edges.User, r)
	if err != nil {
		postResponseWithError(cmd.ChannelID, cmd.ResponseURL, err)
		c.Status(http.StatusOK)
		return
	}

	perms, err := s.i.ListPermsOfRepo(ctx, r, "", 1, 100)
	if err != nil {
		s.log.Error("It has failed to list permissions.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	perms = s.filterPerms(perms, cu)

	// Create a new callback to interact with submissions.
	cb, err := s.i.CreateCallback(ctx, &ent.Callback{
		Type:   callback.TypeDeploy,
		RepoID: r.ID,
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

func parseCmd(cmd string) (string, string) {
	words := strings.Fields(cmd)

	nn := strings.Split(words[1], "/")

	return nn[0], nn[1]
}

// filterPerms returns permissions except the deployer.
func (s *Slack) filterPerms(perms []*ent.Perm, cu *ent.ChatUser) []*ent.Perm {
	ret := []*ent.Perm{}

	for _, p := range perms {
		if p.Edges.User == nil {
			s.log.Warn("The user edge of perm is not found.", zap.Int("perm_id", p.ID))
			continue
		}

		if cu.Edges.User == nil {
			s.log.Warn("The user edge of chat-user is not found.", zap.String("chat_user_id", cu.ID))
			continue
		}

		if p.Edges.User.ID != cu.Edges.User.ID {
			ret = append(ret, p)
		}
	}

	return ret
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
			Value: strconv.FormatInt(u.ID, 10),
		})
	}

	set := []slack.Block{
		slack.NewInputBlock(
			blockEnv,
			slack.NewTextBlockObject(slack.PlainTextType, "Environment", false, false),
			slack.NewOptionsSelectBlockElement(
				slack.OptTypeStatic,
				slack.NewTextBlockObject(slack.PlainTextType, "Select target environment", false, false),
				actionEnv,
				envs...,
			),
		),
		slack.NewInputBlock(
			blockType,
			slack.NewTextBlockObject(slack.PlainTextType, "Type", false, false),
			slack.NewOptionsSelectBlockElement(
				slack.OptTypeStatic,
				slack.NewTextBlockObject(slack.PlainTextType, "Select your ref type", false, false),
				actionType,
				slack.NewOptionBlockObject(
					"commit",
					slack.NewTextBlockObject(slack.PlainTextType, "Commit", false, false),
					nil,
				),
				slack.NewOptionBlockObject(
					"branch",
					slack.NewTextBlockObject(slack.PlainTextType, "Branch", false, false),
					nil,
				),
				slack.NewOptionBlockObject(
					"tag",
					slack.NewTextBlockObject(slack.PlainTextType, "Tag", false, false),
					nil,
				),
			),
		),
		slack.NewInputBlock(
			blockRef,
			slack.NewTextBlockObject(slack.PlainTextType, "Ref", false, false),
			slack.NewPlainTextInputBlockElement(
				slack.NewTextBlockObject(slack.PlainTextType, "E.g. Commit - 25a667d6, Branch - main, Tag - v0.1.2", false, false),
				actionRef,
			),
		),
	}

	if len(approvers) > 0 {
		set = append(set, slack.InputBlock{
			Type:     slack.MBTInput,
			BlockID:  blockApprovers,
			Label:    slack.NewTextBlockObject(slack.PlainTextType, "Approvers", false, false),
			Optional: true,
			Element: slack.NewOptionsSelectBlockElement(
				slack.MultiOptTypeStatic,
				slack.NewTextBlockObject(slack.PlainTextType, "Select approvers", false, false),
				actionApprovers,
				approvers...,
			),
		})
	}

	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: callbackID,
		Title:      slack.NewTextBlockObject(slack.PlainTextType, "Deploy", false, false),
		Submit:     slack.NewTextBlockObject(slack.PlainTextType, "Submit", false, false),
		Close:      slack.NewTextBlockObject(slack.PlainTextType, "Close", false, false),
		Blocks: slack.Blocks{
			BlockSet: set,
		},
	}
}

// interactDeploy deploy with the submitted payload.
func (s *Slack) interactDeploy(c *gin.Context) {
	ctx := c.Request.Context()

	iv, _ := c.Get(KeyIntr)
	itr := iv.(slack.InteractionCallback)

	cv, _ := c.Get(KeyChatUser)
	cu := cv.(*ent.ChatUser)

	cb, _ := s.i.FindCallbackByHash(ctx, itr.View.CallbackID)

	// Parse view submission, and
	// validate values.
	sm := parseViewSubmissions(itr)

	// Validate the entity is processible.
	_, err := s.getCommitSha(ctx, cu.Edges.User, cb.Edges.Repo, sm.Type, sm.Ref)
	if e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
		c.JSON(http.StatusOK, buildErrorsPayload(map[string]string{
			blockRef: "The reference is not found.",
		}))
		return
	} else if err != nil {
		s.log.Error("It has failed to get the SHA of commit.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cfg, err := s.i.GetConfig(ctx, cu.Edges.User, cb.Edges.Repo)
	if err != nil {
		postMessageWithError(cu, err)
		c.Status(http.StatusOK)
		return
	}

	var env *vo.Env
	if env = cfg.GetEnv(sm.Env); env == nil {
		postBotMessage(cu, "The env is not defined in the config.")
		c.Status(http.StatusOK)
		return
	}

	d, err := s.i.Deploy(ctx, cu.Edges.User, cb.Edges.Repo,
		&ent.Deployment{
			Type: deployment.Type(sm.Type),
			Env:  sm.Env,
			Ref:  sm.Ref,
		},
		env,
	)
	if err != nil {
		s.log.Error("It has failed to deploy.", zap.Error(err))
		postMessageWithError(cu, err)
		c.Status(http.StatusOK)
		return
	}

	if _, err := s.i.CreateEvent(ctx, &ent.Event{
		Kind:         event.KindDeployment,
		Type:         event.TypeCreated,
		DeploymentID: d.ID,
	}); err != nil {
		s.log.Error("It has failed to create the event.", zap.Error(err))
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

		return c.SHA, nil
	case "branch":
		b, err := s.i.GetBranch(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return b.CommitSHA, nil
	case "tag":
		t, err := s.i.GetTag(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return t.CommitSHA, nil
	default:
		return "", fmt.Errorf("Type must be one of commit, branch, tag.")
	}
}
