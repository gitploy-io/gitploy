package hooks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/go-github/v32/github"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	Hooks struct {
		*ConfigHooks

		i   Interactor
		log *zap.Logger
	}

	ConfigHooks struct {
		WebhookSecret string
	}
)

const (
	headerGithubEvent     = "X-GitHub-Event"
	headerGithubDelivery  = "X-GitHub-Delivery"
	headerGithubSignature = "X-Hub-Signature-256"
)

func NewHooks(c *ConfigHooks, i Interactor) *Hooks {
	return &Hooks{
		ConfigHooks: c,
		i:           i,
		log:         zap.L().Named("hooks"),
	}
}

func (h *Hooks) HandleHook(c *gin.Context) {
	if isFromGithub(c) {
		h.handleGithubHook(c)
		return
	}

	gb.ErrorResponse(c, http.StatusBadRequest, "It is invalid request.")
}

func (h *Hooks) handleGithubHook(c *gin.Context) {
	ctx := c.Request.Context()

	e := &github.DeploymentStatusEvent{}
	if err := c.ShouldBindBodyWith(e, binding.JSON); err != nil {
		h.log.Error("failed to bind the payload.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "It is invalid request.")
		return
	}

	// Validate Signature if the secret is exist.
	if secret := h.WebhookSecret; secret != "" {
		// Read the payload which was set by the "ShouldBindBodyWith" method call.
		// https://github.com/gin-gonic/gin/issues/439
		var payload []byte
		body, _ := c.Get(gin.BodyBytesKey)
		payload = body.([]byte)

		sig := c.GetHeader(headerGithubSignature)

		if err := github.ValidateSignature(sig, payload, []byte(secret)); err != nil {
			h.log.Error("failed to validate the signature.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to validate the signature.")
			return
		}
	}

	uid := *e.Deployment.ID
	d, err := h.i.FindDeploymentByUID(ctx, uid)
	if err != nil {
		h.log.Error("failed to find the deployment by UID.", zap.Int64("uid", uid))
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to find the deployment by UID.")
		return
	}

	ds := mapGithubDeploymentStatus(e)
	ds.DeploymentID = d.ID

	if ds, err = h.i.CreateDeploymentStatus(ctx, ds); err != nil {
		h.log.Error("failed to create a new the deployment status.", zap.Int64("uid", uid))
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to create a new the deployment status.")
		return
	}

	gb.Response(c, http.StatusCreated, ds)
}

func isFromGithub(c *gin.Context) bool {
	return c.GetHeader(headerGithubDelivery) != ""
}

func mapGithubDeploymentStatus(gds *github.DeploymentStatusEvent) *ent.DeploymentStatus {
	state := *gds.DeploymentStatus.State
	description := *gds.DeploymentStatus.Description
	logURL := *gds.DeploymentStatus.LogURL

	ds := &ent.DeploymentStatus{
		Status:      state,
		Description: description,
		LogURL:      logURL,
	}

	return ds
}
