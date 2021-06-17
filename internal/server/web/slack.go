package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"go.uber.org/zap"
)

type (
	// SlackTokenResponse is the result from exchanging a authorization code
	// for user's access token in Slack.
	// https://api.slack.com/authentication/oauth-v2#exchanging
	SlackTokenResponse struct {
		// Bot access token
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
		// Authenticated user access token.
		User *SlackAuthedUser `json:"authed_user"`
	}

	SlackAuthedUser struct {
		ID          string `json:"id"`
		Scope       string `json:"scope"`
		AccessToken string `json:"access_token"`
	}
)

func (w *Web) SlackIndex(c *gin.Context) {
	s := c.GetString(gb.KeySession)
	if s == "" {
		c.Redirect(http.StatusFound, "/")
		return
	}

	w.redirectToAuthSlack(c)
}

func (w *Web) redirectToAuthSlack(c *gin.Context) {
	const (
		secure   = false
		httpOnly = true
	)
	state := randState()
	c.SetCookie("state", state, 60, "/", "", secure, httpOnly)

	url := w.cc.AuthCodeURL(state)
	c.Redirect(http.StatusFound, url)
}

// SigninSlack authenticate by Slack oAuth
// https://api.slack.com/authentication/oauth-v2#exchanging
func (w *Web) SigninSlack(c *gin.Context) {
	var (
		state = c.Query("state")
		code  = c.Query("code")
	)

	ctx := c.Request.Context()

	s, err := c.Cookie("state")
	if err != nil || state != s {
		c.String(http.StatusInternalServerError, "The state of Slack is invalid. It's possible CSRF or cookies not enabled.")
		return
	}

	sr, err := w.exchangeSlackCode(ctx, code)
	if err != nil {
		w.log.Error("It has failed to exchange the code.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to exchange the code for Slack.")
		return
	}

	hash := c.GetString(gb.KeySession)
	u, _ := w.i.FindUserByHash(ctx, hash)

	_, err = w.i.SaveChatUser(ctx, u, &ent.ChatUser{
		ID:       sr.User.ID,
		Token:    sr.User.AccessToken,
		BotToken: sr.AccessToken,
	})
	if err != nil {
		w.log.Error("It has failed to save the chat user.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to save the chat user.")
		return
	}

	// TODO: redirect to settings page.
	c.Redirect(http.StatusFound, "/")
	return
}

func (w *Web) exchangeSlackCode(ctx context.Context, code string) (*SlackTokenResponse, error) {
	url := fmt.Sprintf("%s?code=%s&client_id=%s&client_secret=%s",
		w.cc.Endpoint.TokenURL,
		code,
		w.cc.ClientID,
		w.cc.ClientSecret,
	)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	s := &SlackTokenResponse{}
	if err = json.Unmarshal(body, s); err != nil {
		return nil, err
	}

	return s, nil
}
