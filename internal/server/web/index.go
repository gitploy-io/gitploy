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
		c     *oauth2.Config
		store Store
		scm   SCM
		log   *zap.Logger
	}

	WebConfig struct {
		Config *oauth2.Config
		Store  Store
		SCM    SCM
	}
)

func NewWeb(wc *WebConfig) *Web {
	return &Web{
		c:     wc.Config,
		store: wc.Store,
		scm:   wc.SCM,
		log:   zap.L().Named("web"),
	}
}

func (w *Web) Index(c *gin.Context) {
	s := c.GetString(gb.KeySession)
	if s != "" {
		w.redirectToAuth(c)
		return
	}

	c.String(http.StatusOK, "OK")
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

	token, err := w.c.Exchange(c, code)
	if err != nil {
		w.log.Error("failed to exchange the code.", zap.Error(err))
		c.String(http.StatusInternalServerError, "There is an issue to exchange the code.")
		return
	}

	if !token.Valid() {
		w.log.Error("invalid token.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It's a invalid token.")
		return
	}

	ctx := c.Request.Context()

	u, err := w.scm.GetUser(ctx, token.AccessToken)
	if err != nil {
		w.log.Error("failed to fetch a user from SCM.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It's failed to fetch a user from SCM.")
		return
	}

	// Set authorization token, and
	// save the user with the token.
	u.Token = token.AccessToken
	u.Refresh = token.RefreshToken
	u.Expiry = token.Expiry

	u, err = w.store.FindUserByID(ctx, u.ID)
	if ent.IsNotFound(err) {
		w.store.CreateUser(ctx, u)
	}

	w.store.UpdateUser(ctx, u)
	c.Redirect(http.StatusFound, "/")
}

func randState() string {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	return state
}
