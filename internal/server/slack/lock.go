package slack

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/perm"
)

func (s *Slack) handleLockCmd(c *gin.Context) {
	ctx := c.Request.Context()

	// SlashCommandParse hvae to be success because
	// it has parsed in the Cmd method.
	cmd, _ := slack.SlashCommandParse(c.Request)

	cu, err := s.i.FindChatUserByID(ctx, cmd.UserID)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "Slack is not connected with Gitploy.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get chat user.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	args := strings.Split(cmd.Text, " ")

	ns, n, err := parseFullName(args[1])
	if err != nil {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("`%s` is invalid repository format.", args[1]))
		c.Status(http.StatusOK)
		return
	}

	r, err := s.i.FindRepoOfUserByNamespaceName(ctx, cu.Edges.User, ns, n)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s` repository is not found.", args[1]))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	p, err := s.i.FindPermOfRepo(ctx, r, cu.Edges.User)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s` repository is not found.", args[1]))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the perm.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if p.RepoPerm != perm.RepoPermAdmin {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "Only admin can lock the repository.")
		c.Status(http.StatusOK)
		return
	}

	// r.Locked = true
	if r, err = s.i.UpdateRepo(ctx, r); err != nil {
		s.log.Error("It has failed to update the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("Lock the `%s` repository successfully.", args[1]))
	c.Status(http.StatusOK)
}

func (s *Slack) handleUnlockCmd(c *gin.Context) {
	ctx := c.Request.Context()

	// SlashCommandParse hvae to be success because
	// it has parsed in the Cmd method.
	cmd, _ := slack.SlashCommandParse(c.Request)

	cu, err := s.i.FindChatUserByID(ctx, cmd.UserID)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "Slack is not connected with Gitploy.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get chat user.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	args := strings.Split(cmd.Text, " ")

	ns, n, err := parseFullName(args[1])
	if err != nil {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("`%s` is invalid repository format.", args[1]))
		c.Status(http.StatusOK)
		return
	}

	r, err := s.i.FindRepoOfUserByNamespaceName(ctx, cu.Edges.User, ns, n)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s` repository is not found.", args[1]))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	p, err := s.i.FindPermOfRepo(ctx, r, cu.Edges.User)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s` repository is not found.", args[1]))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the perm.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if p.RepoPerm != perm.RepoPermAdmin {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "Only admin can unlock the repository.")
		c.Status(http.StatusOK)
		return
	}

	// r.Locked = false
	if r, err = s.i.UpdateRepo(ctx, r); err != nil {
		s.log.Error("It has failed to update the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("Unlock the `%s` repository successfully.", args[1]))
	c.Status(http.StatusOK)
}
