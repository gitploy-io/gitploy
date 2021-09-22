package web

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

type (
	Web struct {
		// OAuth Configuration to sign in.
		c *oauth2.Config

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
		c:   c.Config,
		i:   c.Interactor,
		log: zap.L().Named("web"),
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

	c.String(http.StatusOK, "Index page")
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
		w.log.Error("It has failed to exchange the code.", zap.Error(err))
		c.String(http.StatusUnauthorized, "There is an issue to exchange the code.")
		return
	}

	if !t.Valid() {
		w.log.Error("invalid token.", zap.Error(err))
		c.String(http.StatusUnauthorized, "It's a invalid token.")
		return
	}

	ctx := c.Request.Context()

	ru, err := w.i.GetRemoteUserByToken(ctx, t.AccessToken)
	if err != nil {
		w.log.Error("It has failed to get the remote user.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to get the remote user.")
		return
	}
	w.log.Debug("Get user's login.", zap.String("login", ru.Login))

	orgs, err := w.i.ListRemoteOrgsByToken(ctx, t.AccessToken)
	if err != nil {
		w.log.Error("It has failed to list remote orgs.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to list remote orgs.")
		return
	}
	w.log.Debug("List remote orgs.", zap.Strings("orgs", orgs))

	// Check the login of user who is member and admin.
	if !(w.i.IsEntryMember(ctx, ru.Login) || w.i.IsOrgMember(ctx, orgs)) {
		w.log.Warn("This login not a member of an approved organization.", zap.String("login", ru.Login))
		c.String(http.StatusUnauthorized, "You are not a member of an approved organization.")
		return
	}

	admin := w.i.IsAdminUser(ctx, ru.Login)

	// Synchronize from the remote user. It synchronizes
	// user information and save generated OAuth token.
	u := &ent.User{
		ID:      ru.ID,
		Login:   ru.Login,
		Avatar:  ru.AvatarURL,
		Token:   t.AccessToken,
		Refresh: t.RefreshToken,
		Expiry:  t.Expiry,
		Admin:   admin,
	}

	if _, err = w.i.FindUserByID(ctx, u.ID); ent.IsNotFound(err) {
		lic, err := w.i.GetLicense(ctx)
		if err != nil {
			w.log.Error("It has failed to check the license.", zap.Error(err))
			c.String(http.StatusInternalServerError, "It has failed to check the license.")
			return
		}

		if lic.MemberCount >= lic.MemberLimit {
			w.log.Warn("There are no more seats. It prevents to over the limit.", zap.Int("member_count", lic.MemberCount), zap.Int("member_limit", lic.MemberLimit))
			c.String(http.StatusPaymentRequired, "There are no more seats.")
			return
		}

		u, _ = w.i.CreateUser(ctx, u)
	} else if err != nil {
		w.log.Error("It failed to save the user.", zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to save the user.")
		return
	} else {
		u, _ = w.i.UpdateUser(ctx, u)
	}

	// Register cookie.
	const (
		secure   = false
		httpOnly = true
	)
	c.SetCookie(gb.CookieSession, u.Hash, 0, "/", "", secure, httpOnly)
	c.Redirect(http.StatusFound, "/")
}

func (w *Web) SignOutHTML(c *gin.Context) {
	const (
		secure   = false
		httpOnly = true
	)

	// Delete the cookie
	c.SetCookie(gb.CookieSession, "", -1, "/", "", secure, httpOnly)
	c.HTML(http.StatusOK, "index.html", nil)
}

func (w *Web) SignOutString(c *gin.Context) {
	const (
		secure   = false
		httpOnly = true
	)

	// Delete the cookie
	c.SetCookie(gb.CookieSession, "", -1, "/", "", secure, httpOnly)
	c.String(http.StatusOK, "Signed out")
}

func randState() string {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	return state
}
