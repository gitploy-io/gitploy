package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"go.uber.org/zap"
)

type (
	Repo struct {
		RepoConfig
		i   Interactor
		log *zap.Logger
	}

	RepoConfig struct {
		WebhookURL    string
		WebhookSSL    bool
		WebhookSecret string
	}

	repoPatchPayload struct {
		ConfigPath *string `json:"config_path"`
		Active     *bool   `json:"active"`
	}
)

func NewRepo(c RepoConfig, i Interactor) *Repo {
	return &Repo{
		RepoConfig: c,
		i:          i,
		log:        zap.L().Named("repos"),
	}
}

func (r *Repo) ListRepos(c *gin.Context) {
	var (
		sort      = c.DefaultQuery("sort", "false")
		q         = c.Query("q")
		namespace = c.Query("namespace")
		name      = c.Query("name")
		page      = c.DefaultQuery("page", "1")
		perPage   = c.DefaultQuery("per_page", "30")
	)

	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	sorted, err := strconv.ParseBool(sort)
	if err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The \"sort\" field must be boolean.", err),
		)
		return
	}

	repos, err := r.i.ListReposOfUser(ctx, u, q, namespace, name, sorted, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to list repositories.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, repos)
}

func (r *Repo) UpdateRepo(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	re := rv.(*ent.Repo)

	p := &repoPatchPayload{}
	var err error
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the body.", err),
		)
		return
	}

	// Activate (or Deactivate) the repository:
	// Create a new webhook when it activates the repository,
	// in contrast it remove the webhook when it deactivates.
	if p.Active != nil {
		if *p.Active && !re.Active {
			if re, err = r.i.ActivateRepo(ctx, u, re, &extent.WebhookConfig{
				URL:         r.WebhookURL,
				Secret:      r.WebhookSecret,
				InsecureSSL: r.WebhookSSL,
			}); err != nil {
				r.log.Check(gb.GetZapLogLevel(err), "Failed to activate the repository.").Write(zap.Error(err))
				gb.ResponseWithError(c, err)
				return
			}
		} else if !*p.Active && re.Active {
			if re, err = r.i.DeactivateRepo(ctx, u, re); err != nil {
				r.log.Check(gb.GetZapLogLevel(err), "Failed to deactivate the repository.").Write(zap.Error(err))
				gb.ResponseWithError(c, err)
				return
			}
		}
	}

	if p.ConfigPath != nil {
		if *p.ConfigPath != re.ConfigPath {
			re.ConfigPath = *p.ConfigPath

			if re, err = r.i.UpdateRepo(ctx, re); err != nil {
				r.log.Check(gb.GetZapLogLevel(err), "Failed to update the repository.").Write(zap.Error(err))
				gb.ResponseWithError(c, err)
				return
			}
		}
	}

	gb.Response(c, http.StatusOK, re)
}

func (r *Repo) GetRepo(c *gin.Context) {
	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	gb.Response(c, http.StatusOK, repo)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
