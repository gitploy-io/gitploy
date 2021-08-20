package repos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"github.com/hanjunlee/gitploy/vo"
	"go.uber.org/zap"
)

type (
	Repo struct {
		RepoConfig
		i   Interactor
		log *zap.Logger
	}

	RepoConfig struct {
		ProxyProto    string
		ProxyHost     string
		WebhookPath   string
		WebhookSecret string
	}

	RepoPayload struct {
		ConfigPath string `json:"config_path"`
		Active     bool   `json:"active"`
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
		sort    = c.DefaultQuery("sort", "false")
		q       = c.Query("q")
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")

		repos []*ent.Repo
	)

	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	sorted, err := strconv.ParseBool(sort)
	if err != nil {
		r.log.Error("invalid sort value.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "It was invalid request.")
		return
	}

	repos, err = r.i.ListReposOfUser(ctx, u, sorted, q, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list repositories.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list repositories.")
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

	p := &RepoPayload{}
	var err error
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the body")
		return
	}

	// Activate (or Deactivate) the repository:
	// Create a new webhook when it activates the repository,
	// in contrast it remove the webhook when it deactivates.
	if p.Active && !re.Active {
		url := fmt.Sprintf("%s://%s%s", r.ProxyProto, r.ProxyHost, r.WebhookPath)

		if re, err = r.i.ActivateRepo(ctx, u, re, &vo.WebhookConfig{
			URL:         url,
			Secret:      r.WebhookSecret,
			InsecureSSL: isSecure(r.ProxyProto),
		}); err != nil {
			r.log.Error("It has failed to activate the repo.", zap.Error(err), zap.String("url", url))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to activate the repository.")
			return
		}
	} else if !p.Active && re.Active {
		if re, err = r.i.DeactivateRepo(ctx, u, re); err != nil {
			r.log.Error("failed to deactivate the repo.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to deactivate the repository.")
			return
		}
	}

	if re.ConfigPath != p.ConfigPath {
		re.ConfigPath = p.ConfigPath

		if re, err = r.i.UpdateRepo(ctx, re); err != nil {
			r.log.Error("failed to update the repo", zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to update the repository.")
			return
		}
	}

	gb.Response(c, http.StatusOK, re)
}

func (r *Repo) GetRepo(c *gin.Context) {
	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	gb.Response(c, http.StatusOK, repo)
}

func (r *Repo) GetRepoByNamespaceName(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		namespace = c.Query("namespace")
		name      = c.Query("name")
	)

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	repo, err := r.i.FindRepoOfUserByNamespaceName(ctx, u, namespace, name)
	if ent.IsNotFound(err) {
		r.log.Error("It has failed to find the repo.", zap.String("repo", name), zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The repository is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the repository.", zap.String("repo", name), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
		return
	}

	gb.Response(c, http.StatusOK, repo)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func isSecure(proto string) bool {
	return proto == "https"
}
