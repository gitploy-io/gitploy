package slack

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/callback"
	"github.com/gitploy-io/gitploy/ent/perm"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	lockViewSubmission struct {
		Env string
	}
)

func (s *Slack) handleLockCmd(c *gin.Context) {
	ctx := c.Request.Context()

	av, _ := c.Get(KeyCmd)
	cmd := av.(slack.SlashCommand)

	bv, _ := c.Get(KeyChatUser)
	cu := bv.(*ent.ChatUser)

	s.log.Debug("Processing lock command.", zap.String("command", cmd.Text))
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

	// Validate the perm for the repo.
	if p, err := s.i.FindPermOfRepo(ctx, r, cu.Edges.User); !(p.RepoPerm == perm.RepoPermWrite || p.RepoPerm == perm.RepoPermAdmin) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "Write perm is required to lock the environment.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the perm.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	// Build the modal with unlocked envs.
	config, err := s.i.GetConfig(ctx, cu.Edges.User, r)
	if vo.IsConfigNotFoundError(err) || vo.IsConfigParseError(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "The config is invalid.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the config file.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	locks, err := s.i.ListLocksOfRepo(ctx, r)
	if err != nil {
		s.log.Error("It has failed to list locks.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cb, err := s.i.CreateCallback(ctx, &ent.Callback{
		Type:   callback.TypeLock,
		RepoID: r.ID,
	})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = slack.New(cu.BotToken).
		OpenViewContext(ctx, cmd.TriggerID, buildLockView(cb.Hash, config, locks))
	if err != nil {
		s.log.Error("It has failed to open a new view.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func buildLockView(callbackID string, c *vo.Config, locks []*ent.Lock) slack.ModalViewRequest {
	hasLocked := func(env string) bool {
		for _, lock := range locks {
			if lock.Env == env {
				return true
			}
		}

		return false
	}

	envs := []*slack.OptionBlockObject{}
	for _, env := range c.Envs {
		if hasLocked(env.Name) {
			continue
		}

		envs = append(envs, &slack.OptionBlockObject{
			Text: &slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: env.Name,
			},
			Value: env.Name,
		})
	}

	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: callbackID,
		Title:      slack.NewTextBlockObject(slack.PlainTextType, "Lock", false, false),
		Submit:     slack.NewTextBlockObject(slack.PlainTextType, "Submit", false, false),
		Close:      slack.NewTextBlockObject(slack.PlainTextType, "Close", false, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.NewInputBlock(
					blockEnv,
					slack.NewTextBlockObject(slack.PlainTextType, "Environment", false, false),
					slack.NewOptionsSelectBlockElement(
						slack.OptTypeStatic,
						slack.NewTextBlockObject(slack.PlainTextType, "Select the environment", false, false),
						actionEnv,
						envs...,
					),
				),
			},
		},
	}
}

func (s *Slack) handleUnlockCmd(c *gin.Context) {
	ctx := c.Request.Context()

	av, _ := c.Get(KeyCmd)
	cmd := av.(slack.SlashCommand)

	bv, _ := c.Get(KeyChatUser)
	cu := bv.(*ent.ChatUser)

	s.log.Debug("Processing lock command.", zap.String("command", cmd.Text))
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

	// Validate the perm for the repo.
	if p, err := s.i.FindPermOfRepo(ctx, r, cu.Edges.User); !(p.RepoPerm == perm.RepoPermWrite || p.RepoPerm == perm.RepoPermAdmin) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "Write perm is required to lock the environment.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the perm.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	// Build the modal with unlocked envs.
	locks, err := s.i.ListLocksOfRepo(ctx, r)
	if len(locks) == 0 {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("There is no locked envs in the `%s/%s` repository.", ns, n))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to list locks.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cb, err := s.i.CreateCallback(ctx, &ent.Callback{
		Type:   callback.TypeUnlock,
		RepoID: r.ID,
	})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = slack.New(cu.BotToken).
		OpenViewContext(ctx, cmd.TriggerID, buildUnlockView(cb.Hash, locks))
	if err != nil {
		s.log.Error("It has failed to open a new view.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func buildUnlockView(callbackID string, locks []*ent.Lock) slack.ModalViewRequest {
	envs := []*slack.OptionBlockObject{}
	for _, lock := range locks {

		envs = append(envs, &slack.OptionBlockObject{
			Text: &slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: lock.Env,
			},
			Value: lock.Env,
		})
	}

	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: callbackID,
		Title:      slack.NewTextBlockObject(slack.PlainTextType, "Unlock", false, false),
		Submit:     slack.NewTextBlockObject(slack.PlainTextType, "Submit", false, false),
		Close:      slack.NewTextBlockObject(slack.PlainTextType, "Close", false, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.NewInputBlock(
					blockEnv,
					slack.NewTextBlockObject(slack.PlainTextType, "Environment", false, false),
					slack.NewOptionsSelectBlockElement(
						slack.OptTypeStatic,
						slack.NewTextBlockObject(slack.PlainTextType, "Select the environment", false, false),
						actionEnv,
						envs...,
					),
				),
			},
		},
	}
}

func (s *Slack) interactLock(c *gin.Context) {
	ctx := c.Request.Context()

	iv, _ := c.Get(KeyIntr)
	itr := iv.(slack.InteractionCallback)

	cv, _ := c.Get(KeyChatUser)
	cu := cv.(*ent.ChatUser)

	cb, _ := s.i.FindCallbackByHash(ctx, itr.View.CallbackID)

	sm := parseLockViewSubmissions(itr)

	if _, err := s.i.CreateLock(ctx, &ent.Lock{
		Env:    sm.Env,
		UserID: cu.Edges.User.ID,
		RepoID: cb.Edges.Repo.ID,
	}); err != nil {
		s.log.Error("It has failed to lock the env.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	postBotMessage(cu, fmt.Sprintf("Success to lock the `%s` environment of the `%s` repository.", sm.Env, cb.Edges.Repo.GetFullName()))
	c.Status(http.StatusOK)
}

func (s *Slack) interactUnlock(c *gin.Context) {
	ctx := c.Request.Context()

	iv, _ := c.Get(KeyIntr)
	itr := iv.(slack.InteractionCallback)

	cv, _ := c.Get(KeyChatUser)
	cu := cv.(*ent.ChatUser)

	cb, _ := s.i.FindCallbackByHash(ctx, itr.View.CallbackID)

	sm := parseLockViewSubmissions(itr)

	lock, err := s.i.FindLockOfRepoByEnv(ctx, cb.Edges.Repo, sm.Env)
	if ent.IsNotFound(err) {
		postBotMessage(cu, fmt.Sprintf("The `%s` environment is not locked.", sm.Env))
		c.Status(http.StatusOK)
	} else if err != nil {
		s.log.Error("It has failed to find the lock.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := s.i.DeleteLock(ctx, lock); err != nil {
		s.log.Error("It has failed to unlock the env.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	postBotMessage(cu, fmt.Sprintf("Success to unlock the `%s` environment of the `%s` repository.", sm.Env, cb.Edges.Repo.GetFullName()))
	c.Status(http.StatusOK)
}

func parseLockViewSubmissions(itr slack.InteractionCallback) *lockViewSubmission {
	sm := &lockViewSubmission{}

	values := itr.View.State.Values
	if v, ok := values[blockEnv][actionEnv]; ok {
		sm.Env = v.SelectedOption.Value
	}

	return sm
}
