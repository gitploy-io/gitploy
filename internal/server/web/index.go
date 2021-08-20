package web

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	Web struct {
		// OAuth Configuration to sign in.
		c *oauth2.Config

		// Identify user as admin when user sign in.
		adminUsers []string

		i Interactor

		log *zap.Logger
	}

	WebConfig struct {
		Config     *oauth2.Config
		Interactor Interactor
		AdminUsers []string
	}
)

func NewWeb(c *WebConfig) *Web {
	return &Web{
		c:          c.Config,
		adminUsers: c.AdminUsers,
		i:          c.Interactor,
		log:        zap.L().Named("web"),
	}
}

func (w *Web) IndexHTML(c *gin.Context) {
	_, ok := c.Get(gb.KeyUser)
	if !ok {
		w.redirectToAuth(c)
		return
	}

	c.HTML(http.StatusOK, "index.html", nil)
}

func (w *Web) IndexString(c *gin.Context) {
	_, ok := c.Get(gb.KeyUser)
	if !ok {
		w.redirectToAuth(c)
		return
	}

	c.String(http.StatusOK, "Ok")
}

func (w *Web) redirectToAuth(c *gin.Context) {
	const (
		secure   = false
		httpOnly = true
	)
	state := randState()
	c.SetCookie("state", state, 60, "/", "", secure, httpOnly)

	url := w.c.AuthCodeURL(state)
	c.Redirect(http.StatusFound, url)
}

func (w *Web) Signin(c *gin.Context) {
	var (
		state = c.Query("state")
		code  = c.Query("code")
	)
	s, err := c.Cookie("state")
	if err != nil || state != s {
		w.log.Error("The state is invalid")
		c.String(http.StatusInternalServerError, "The state is invalid. It's possible CSRF or cookies not enabled.")
		return
	}

	t, err := w.c.Exchange(c, code)
	if err != nil {
		w.log.Error("failed to exchange the code.", zap.Error(err))
		c.String(http.StatusInternalServerError, "There is an issue to exchange the code.")
		return
	}

	if !t.Valid() {
		w.log.Error("invalid token.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It's a invalid token.")
		return
	}

	ctx := c.Request.Context()

	ru, err := w.i.GetRemoteUserByToken(ctx, t.AccessToken)
	if err != nil {
		w.log.Error("failed to fetch a user from SCM.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to fetch a user from SCM.")
		return
	}

	// Synchronize from the remote user. It synchronizes
	// user information and save generated OAuth token.
	u := &ent.User{
		ID:      ru.ID,
		Login:   ru.Login,
		Avatar:  ru.AvatarURL,
		Token:   t.AccessToken,
		Refresh: t.RefreshToken,
		Expiry:  t.Expiry,
		Admin:   w.isAdmin(ru.Login),
	}

	if u, err = w.i.SaveUser(ctx, u); err != nil {
		w.log.Error("failed to save the user.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to save the user.")
		return
	}

	// Register cookie.
	const (
		secure   = false
		httpOnly = true
	)
	c.SetCookie(gb.CookieSession, u.Hash, 0, "/", "", secure, httpOnly)
	c.Redirect(http.StatusFound, "/")
}

func (w *Web) isAdmin(login string) bool {
	for _, au := range w.adminUsers {
		if login == au {
			return true
		}
	}

	return false
}

func (w *Web) SignOut(c *gin.Context) {
	const (
		secure   = false
		httpOnly = true
	)

	// Delete the cookie
	c.SetCookie(gb.CookieSession, "", -1, "/", "", secure, httpOnly)
	c.Redirect(http.StatusFound, "/")
}

func randState() string {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	return state
}
