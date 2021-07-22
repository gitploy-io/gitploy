package slack

import (
	"context"
	"crypto/rand"
	"encoding/base64"
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

func (s *Slack) Index(c *gin.Context) {
	_, ok := c.Get(gb.KeyUser)
	if !ok {
		c.Redirect(http.StatusFound, "/")
		return
	}

	s.redirectToAuth(c)
}

func (s *Slack) redirectToAuth(c *gin.Context) {
	const (
		secure   = false
		httpOnly = true
	)
	state := randState()
	c.SetCookie("state", state, 60, "/", "", secure, httpOnly)

	url := s.c.AuthCodeURL(state)
	c.Redirect(http.StatusFound, url)
}

// SigninSlack authenticate by Slack oAuth
// https://api.slack.com/authentication/oauth-v2#exchanging
func (s *Slack) SigninSlack(c *gin.Context) {
	var (
		state = c.Query("state")
		code  = c.Query("code")
	)

	ctx := c.Request.Context()

	sv, err := c.Cookie("state")
	if err != nil || state != sv {
		c.String(http.StatusInternalServerError, "The state of Slack is invalid. It's possible CSRF or cookies not enabled.")
		return
	}

	sr, err := s.exchangeSlackCode(ctx, code)
	if err != nil {
		s.log.Error("It has failed to exchange the code.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to exchange the code for Slack.")
		return
	}

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	_, err = s.i.SaveChatUser(ctx, u, &ent.ChatUser{
		ID:       sr.User.ID,
		Token:    sr.User.AccessToken,
		BotToken: sr.AccessToken,
	})
	if err != nil {
		s.log.Error("It has failed to save the chat user.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to save the chat user.")
		return
	}

	// TODO: redirect to settings page.
	c.Redirect(http.StatusFound, "/")
}

func (s *Slack) exchangeSlackCode(ctx context.Context, code string) (*SlackTokenResponse, error) {
	url := fmt.Sprintf("%s?code=%s&client_id=%s&client_secret=%s",
		s.c.Endpoint.TokenURL,
		code,
		s.c.ClientID,
		s.c.ClientSecret,
	)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	sr := &SlackTokenResponse{}
	if err = json.Unmarshal(body, sr); err != nil {
		return nil, err
	}

	return sr, nil
}

func randState() string {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	return state
}
