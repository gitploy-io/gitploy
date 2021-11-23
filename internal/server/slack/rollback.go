package slack

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nleeper/goment"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/callback"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/event"
	"github.com/gitploy-io/gitploy/vo"
)

const (
	blockDeployment  = "block_deployment"
	actionDeployment = "aciton_deployment"
)

type (
	rollbackViewSubmission struct {
		DeploymentID int
		ApproverIDs  []string
	}

	deploymentAggregation struct {
		envName     string
		deployments []*ent.Deployment
	}
)

// handleRollbackCmd handles rollback command: "/gitploy rollback OWNER/REPO".
func (s *Slack) handleRollbackCmd(c *gin.Context) {
	ctx := c.Request.Context()

	av, _ := c.Get(KeyCmd)
	cmd := av.(slack.SlashCommand)

	bv, _ := c.Get(KeyChatUser)
	cu := bv.(*ent.ChatUser)

	s.log.Debug("Processing rollback command.", zap.String("command", cmd.Text))
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

	cb, err := s.i.CreateCallback(ctx, &ent.Callback{
		Type:   callback.TypeRollback,
		RepoID: r.ID,
	})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	as := s.getSuccessfulDeploymentAggregation(ctx, r, config)

	_, err = slack.New(cu.BotToken).
		OpenViewContext(ctx, cmd.TriggerID, buildRollbackView(cb.Hash, as, perms))
	if err != nil {
		s.log.Error("It has failed to open a dialog.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func buildRollbackView(callbackID string, as []*deploymentAggregation, perms []*ent.Perm) slack.ModalViewRequest {
	groups := []*slack.OptionGroupBlockObject{}

	for _, a := range as {
		options := []*slack.OptionBlockObject{}

		for _, d := range a.deployments {
			created, _ := goment.New(d.CreatedAt)

			options = append(options, slack.NewOptionBlockObject(
				strconv.Itoa(d.ID),
				slack.NewTextBlockObject(
					slack.PlainTextType,
					fmt.Sprintf("#%d - %s deployed %s", d.ID, d.GetShortRef(), created.FromNow()),
					false, false),
				nil))
		}

		groups = append(groups, slack.NewOptionGroupBlockElement(
			slack.NewTextBlockObject(slack.PlainTextType, string(a.envName), false, false),
			options...))
	}

	approvers := []*slack.OptionBlockObject{}
	for _, perm := range perms {
		u := perm.Edges.User
		if u == nil {
			continue
		}

		approvers = append(approvers, slack.NewOptionBlockObject(
			strconv.FormatInt(u.ID, 10),
			slack.NewTextBlockObject(slack.PlainTextType, u.Login, false, false),
			nil))
	}

	sets := []slack.Block{
		slack.NewInputBlock(
			blockDeployment,
			slack.NewTextBlockObject(slack.PlainTextType, "Deployments", false, false),
			slack.NewOptionsGroupSelectBlockElement(
				slack.OptTypeStatic,
				slack.NewTextBlockObject(slack.PlainTextType, "Select target deployment", false, false),
				actionDeployment,
				groups...,
			),
		),
	}

	if len(approvers) > 0 {
		sets = append(sets, slack.InputBlock{
			Type:     slack.MBTInput,
			BlockID:  blockApprovers,
			Optional: true,
			Label:    slack.NewTextBlockObject(slack.PlainTextType, "Approvers", false, false),
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
		Title:      slack.NewTextBlockObject(slack.PlainTextType, "Rollback", false, false),
		Submit:     slack.NewTextBlockObject(slack.PlainTextType, "Submit", false, false),
		Close:      slack.NewTextBlockObject(slack.PlainTextType, "Close", false, false),
		Blocks: slack.Blocks{
			BlockSet: sets,
		},
	}
}

func (s *Slack) getSuccessfulDeploymentAggregation(ctx context.Context, r *ent.Repo, cf *vo.Config) []*deploymentAggregation {
	a := []*deploymentAggregation{}

	for _, env := range cf.Envs {
		ds, _ := s.i.ListDeploymentsOfRepo(ctx, r, env.Name, string(deployment.StatusSuccess), 1, 5)
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

func (s *Slack) interactRollback(c *gin.Context) {
	ctx := c.Request.Context()

	iv, _ := c.Get(KeyIntr)
	itr := iv.(slack.InteractionCallback)

	cv, _ := c.Get(KeyChatUser)
	cu := cv.(*ent.ChatUser)

	cb, _ := s.i.FindCallbackByHash(ctx, itr.View.CallbackID)

	sm := parseRollbackSubmissions(itr)

	d, err := s.i.FindDeploymentByID(ctx, sm.DeploymentID)
	if err != nil {
		s.log.Error("It has failed to find the deployment.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	config, err := s.i.GetConfig(ctx, cu.Edges.User, cb.Edges.Repo)
	if err != nil {
		postMessageWithError(cu, err)
		c.Status(http.StatusOK)
		return
	}

	if err := config.Eval(&vo.EvalValues{}); err != nil {
		postMessageWithError(cu, err)
		c.Status(http.StatusOK)
		return
	}

	var env *vo.Env
	if env = config.GetEnv(d.Env); env == nil {
		postBotMessage(cu, "The env is not defined in the config.")
		c.Status(http.StatusOK)
		return
	}

	d, err = s.i.Deploy(ctx, cu.Edges.User, cb.Edges.Repo, &ent.Deployment{
		Type:       deployment.Type(d.Type),
		Ref:        d.Ref,
		Sha:        d.Sha,
		Env:        d.Env,
		IsRollback: true,
	}, env)
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

func parseRollbackSubmissions(itr slack.InteractionCallback) *rollbackViewSubmission {
	sm := &rollbackViewSubmission{}

	values := itr.View.State.Values
	if v, ok := values[blockDeployment][actionDeployment]; ok {
		sm.DeploymentID = atoi(v.SelectedOption.Value)
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
