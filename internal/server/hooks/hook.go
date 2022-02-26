package hooks

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/go-github/v42/github"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/pkg/e"
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
	if !isFromGithub(c) {
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, nil))
		return
	}

	if h.WebhookSecret != "" {
		if err := h.validateGitHubSignature(c); err != nil {
			h.log.Warn("Failed to validate the signature.", zap.Error(err))
			gb.ResponseWithError(
				c,
				e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to validate the signature.", err),
			)
			return
		}
	}

	switch eventName := c.GetHeader(headerGtihubEvent); eventName {
	case "deployment_status":
		h.handleGithubDeploymentEvent(c)
		return
	case "push":
		h.handleGithubPushEvent(c)
		return
	default:
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, nil))
		return
	}
}

func (h *Hooks) validateGitHubSignature(c *gin.Context) error {
	// Read the payload which was set by the "ShouldBindBodyWith" method call.
	// https://github.com/gin-gonic/gin/issues/439
	var b interface{}
	c.ShouldBindBodyWith(b, binding.JSON)

	var payload []byte
	body, _ := c.Get(gin.BodyBytesKey)
	payload = body.([]byte)

	sig := c.GetHeader(headerGithubSignature)

	return github.ValidateSignature(sig, payload, []byte(h.WebhookSecret))
}

func (h *Hooks) handleGithubDeploymentEvent(c *gin.Context) {
	ctx := c.Request.Context()

	evt := &github.DeploymentStatusEvent{}
	if err := c.ShouldBindBodyWith(evt, binding.JSON); err != nil {
		h.log.Warn("Failed to bind the payload.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", err),
		)
		return
	}

	// Convert event to the deployment status.
	ds := mapGithubDeploymentStatus(evt)

	uid := *evt.Deployment.ID
	d, err := h.i.FindDeploymentByUID(ctx, uid)
	if err != nil {
		h.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployment by UID.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	ds.DeploymentID = d.ID
	if ds, err = h.i.SyncDeploymentStatus(ctx, ds); err != nil {
		h.log.Check(gb.GetZapLogLevel(err), "Failed to create a new the deployment status.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	d.Status = mapGithubState(ds.Status)
	if _, err := h.i.UpdateDeployment(ctx, d); err != nil {
		h.log.Check(gb.GetZapLogLevel(err), "Failed to update the deployment.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
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
		d.ProductionEnvironment &&
		d.Edges.Repo != nil {
		if _, err := h.i.ProduceDeploymentStatisticsOfRepo(ctx, d.Edges.Repo, d); err != nil {
			h.log.Error("It has failed to produce the statistics of deployment.", zap.Error(err))
		}
	}

	gb.Response(c, http.StatusCreated, ds)
}

func (h *Hooks) handleGithubPushEvent(c *gin.Context) {
	ctx := c.Request.Context()

	evt := &github.PushEvent{}
	if err := c.ShouldBindBodyWith(evt, binding.JSON); err != nil {
		h.log.Warn("Failed to bind the payload.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", err),
		)
		return
	}

	r, err := h.i.FindRepoByID(ctx, *evt.Repo.ID)
	if err != nil {
		h.log.Check(gb.GetZapLogLevel(err), "Failed to find the repository.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	} else if r.Edges.Owner == nil {
		h.log.Warn("The owner is not found.", zap.Int64("repo_id", r.ID))
		gb.ResponseWithError(c,
			e.NewErrorWithMessage(e.ErrorCodeInternalError, "The owner is not found.", nil),
		)
		return
	}

	config, err := h.i.GetConfig(ctx, r.Edges.Owner, r)
	if err != nil {
		h.log.Check(gb.GetZapLogLevel(err), "Failed to find the configuration file.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	for _, env := range config.Envs {
		ok, err := env.IsAutoDeployOn(*evt.Ref)
		if err != nil {
			h.log.Warn("Failed to validate the ref is matched with 'auto_deploy_on'.", zap.Error(err))
			continue
		}
		if !ok {
			continue
		}

		typ, ref, err := parseGithubRef(*evt.Ref)
		if err != nil {
			h.log.Error("Failed to parse the ref.", zap.Error(err))
			continue
		}

		h.log.Info("Trigger to deploy the ref.", zap.String("ref", *evt.Ref), zap.String("environment", env.Name))
		d := &ent.Deployment{
			Type: typ,
			Ref:  ref,
			Env:  env.Name,
		}
		d, err = h.i.Deploy(ctx, r.Edges.Owner, r, d, env)
		if err != nil {
			h.log.Error("Failed to deploy.", zap.Error(err))
			continue
		}

		if _, err := h.i.CreateEvent(ctx, &ent.Event{
			Kind:         event.KindDeployment,
			Type:         event.TypeCreated,
			DeploymentID: d.ID,
		}); err != nil {
			h.log.Error("It has failed to create the event.", zap.Error(err))
		}
	}

	c.Status(http.StatusOK)
}

func isFromGithub(c *gin.Context) bool {
	return c.GetHeader(headerGithubDelivery) != ""
}

func mapGithubDeploymentStatus(e *github.DeploymentStatusEvent) *ent.DeploymentStatus {
	var (
		logURL string
	)

	// target_url is deprecated.
	if e.DeploymentStatus.TargetURL != nil {
		logURL = *e.DeploymentStatus.TargetURL
	}

	if e.DeploymentStatus.LogURL != nil {
		logURL = *e.DeploymentStatus.LogURL
	}

	ds := &ent.DeploymentStatus{
		Status:      *e.DeploymentStatus.State,
		Description: *e.DeploymentStatus.Description,
		LogURL:      logURL,
		CreatedAt:   e.DeploymentStatus.CreatedAt.Time.UTC(),
		UpdatedAt:   e.DeploymentStatus.UpdatedAt.Time.UTC(),
	}

	return ds
}

// mapGithubState convert state into the status of deployment:
// "in_progress", "queued" to "running",
// "success" to "success", and "failure" to "failure".
func mapGithubState(state string) deployment.Status {
	switch state {
	case "queued":
		return deployment.StatusQueued
	case "success":
		return deployment.StatusSuccess
	case "failure":
		return deployment.StatusFailure
	case "error":
		return deployment.StatusFailure
	default:
		return deployment.StatusRunning
	}
}

func parseGithubRef(ref string) (deployment.Type, string, error) {
	const (
		prefixBranchRef = "refs/heads/"
		prefixTagRef    = "refs/tags/"
	)

	if strings.HasPrefix(ref, prefixBranchRef) {
		return deployment.TypeBranch, strings.TrimPrefix(ref, prefixBranchRef), nil
	}

	if strings.HasPrefix(ref, prefixTagRef) {
		return deployment.TypeTag, strings.TrimPrefix(ref, prefixTagRef), nil
	}

	return "", "", e.NewErrorWithMessage(e.ErrorCodeInternalError, "The ref must be one of branch or tag.", nil)
}
