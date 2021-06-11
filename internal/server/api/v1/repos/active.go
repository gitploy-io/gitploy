package repos

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"github.com/hanjunlee/gitploy/vo"
)

func (r *Repo) Activate(c *gin.Context) {
	uv, _ := c.Get(gb.KeyUser)
	u, _ := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	re, _ := rv.(*ent.Repo)

	ctx := c.Request.Context()

	hid, err := r.scm.CreateWebhook(ctx, u, re, &vo.WebhookConfig{
		URL:         r.WebhookURL,
		Secret:      r.WebhookSecret,
		InsecureSSL: isSecure(r.WebhookURL),
	})
	if err != nil {
		r.log.Error("failed to create a new webhook.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to create a new webhook.")
	}

	re.WebhookID = hid
	_, err = r.store.Activate(ctx, re)
	if err != nil {
		r.log.Error("failed to activate the webhook.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to activate the webhook.")
	}

	gb.Response(c, http.StatusOK, nil)
}

func (r *Repo) Deactivate(c *gin.Context) {
	uv, _ := c.Get(gb.KeyUser)
	u, _ := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	re, _ := rv.(*ent.Repo)

	ctx := c.Request.Context()

	err := r.scm.DeleteWebhook(ctx, u, re, re.WebhookID)
	if err != nil {
		r.log.Error("failed to delete the webhook.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to delete the webhook.")
	}

	_, err = r.store.Deactivate(ctx, re)
	if err != nil {
		r.log.Error("failed to deactivate the webhook.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to deactivate the webhook.")
	}

	gb.Response(c, http.StatusOK, nil)
}

func isSecure(raw string) bool {
	u, _ := url.Parse(raw)
	if u.Scheme == "https" {
		return true
	}
	return false
}
