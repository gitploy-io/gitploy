package hooks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/go-github/v32/github"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/event"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
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
	// Github webhook payload
	// https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
	headerGithubDelivery  = "X-GitHub-Delivery"
	headerGtihubEvent     = "X-GitHub-Event"
	headerGithubSignature = "X-Hub-Signature-256"
)

func NewHooks(c *ConfigHooks, i Interactor) *Hooks {
	return &Hooks{
		ConfigHooks: c,
		i:           i,
		log:         zap.L().Named("hooks"),
	}
}

// HandleHook handles deployment status event, basically,
// it creates a new deployment status for the deployment.
func (h *Hooks) HandleHook(c *gin.Context) {
	if isFromGithub(c) {
		h.handleGithubHook(c)
		return
	}

	gb.ErrorResponse(c, http.StatusBadRequest, "It is invalid request.")
}

func (h *Hooks) handleGithubHook(c *gin.Context) {
	ctx := c.Request.Context()

	if !isGithubDeploymentStatusEvent(c) {
		c.String(http.StatusOK, "It is not deployment status event.")
		return
	}

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

	// Convert event to the deployment status.
	ds := mapGithubDeploymentStatus(e)

	uid := *e.Deployment.ID
	d, err := h.i.FindDeploymentByUID(ctx, uid)
	if err != nil {
		h.log.Error("It has failed to find the deployment by UID.", zap.Int64("deployment_uid", uid), zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to find the deployment by UID.")
		return
	}

	ds.DeploymentID = d.ID
	if ds, err = h.i.CreateDeploymentStatus(ctx, ds); err != nil {
		h.log.Error("It has failed to create a new the deployment status.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to create a new the deployment status.")
		return
	}

	d.Status = mapGithubState(ds.Status)
	if _, err := h.i.UpdateDeployment(ctx, d); err != nil {
		h.log.Error("It has failed to update the deployment status.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to update the deployment status.")
		return
	}

	if _, err := h.i.CreateEvent(ctx, &ent.Event{
		Kind:         event.KindDeployment,
		Type:         event.TypeUpdated,
		DeploymentID: d.ID,
	}); err != nil {
		h.log.Error("It has failed to create the event.", zap.Error(err))
	}

	// Produce statistics when the deployment is success, and production environment.
	if d.Status == deployment.StatusSuccess &&
		d.ProductionEnvironment == true &&
		d.Edges.Repo != nil {
		if _, err := h.i.ProduceDeploymentStatisticsOfRepo(ctx, d.Edges.Repo, d); err != nil {
			h.log.Error("It has failed to produce the statistics of deployment.", zap.Error(err))
		}
	}

	gb.Response(c, http.StatusCreated, ds)
}

func isFromGithub(c *gin.Context) bool {
	return c.GetHeader(headerGithubDelivery) != ""
}

func isGithubDeploymentStatusEvent(c *gin.Context) bool {
	return c.GetHeader(headerGtihubEvent) == "deployment_status"
}

func mapGithubDeploymentStatus(gds *github.DeploymentStatusEvent) *ent.DeploymentStatus {
	var (
		state       = *gds.DeploymentStatus.State
		description = *gds.DeploymentStatus.Description
		logURL      string
	)

	// target_url is deprecated.
	if gds.DeploymentStatus.TargetURL != nil {
		logURL = *gds.DeploymentStatus.TargetURL
	}

	if gds.DeploymentStatus.LogURL != nil {
		logURL = *gds.DeploymentStatus.LogURL
	}

	ds := &ent.DeploymentStatus{
		Status:      state,
		Description: description,
		LogURL:      logURL,
	}

	return ds
}

// mapGithubState convert state into the status of deployment:
// "in_progress", "queued" to "running",
// "success" to "success", and "failure" to "failure".
func mapGithubState(state string) deployment.Status {
	switch state {
	case "success":
		return deployment.StatusSuccess
	case "failure":
		return deployment.StatusFailure
	default:
		return deployment.StatusRunning
	}
}
