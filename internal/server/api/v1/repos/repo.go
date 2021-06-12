package repos

import (
	"net/http"
	"net/url"
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
		store Store
		scm   SCM
		log   *zap.Logger
	}

	RepoConfig struct {
		WebhookURL    string
		WebhookSecret string
	}

	repoPayload struct {
		ConfigPath string `json:"config_path"`
	}
)

func NewRepo(c RepoConfig, store Store, scm SCM) *Repo {
	return &Repo{
		RepoConfig: c,
		store:      store,
		scm:        scm,
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

	if sorted {
		repos, err = r.store.ListSortedRepos(ctx, u, q, atoi(page), atoi(perPage))
	} else {
		repos, err = r.store.ListRepos(ctx, u, q, atoi(page), atoi(perPage))
	}
	if err != nil {
		r.log.Error("failed to list repositories.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list repositories.")
		return
	}

	gb.Response(c, http.StatusOK, repos)
}

func (r *Repo) UpdateRepo(c *gin.Context) {
	rv, _ := c.Get(KeyRepo)
	re := rv.(*ent.Repo)

	p := &repoPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the body")
		return
	}

	ctx := c.Request.Context()

	re.ConfigPath = p.ConfigPath
	_, err := r.store.UpdateRepo(ctx, re)
	if err != nil {
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to update the repository.")
		return
	}

	gb.Response(c, http.StatusOK, nil)
}

func (r *Repo) GetRepo(c *gin.Context) {
	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	gb.Response(c, http.StatusOK, repo)
}

func (r *Repo) GetRepoByNamespaceName(c *gin.Context) {
	var (
		namespace = c.Query("namespace")
		name      = c.Query("name")
	)

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	ctx := c.Request.Context()

	repo, err := r.store.FindRepoByNamespaceName(ctx, u, namespace, name)
	if ent.IsNotFound(err) {
		r.log.Error("failed to access the repo.", zap.String("repo", name), zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "It has failed to search the repo.")
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
		return
	}

	re.WebhookID = hid
	re, err = r.store.Activate(ctx, re)
	if err != nil {
		r.log.Error("failed to activate the webhook.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to activate the webhook.")
		return
	}

	gb.Response(c, http.StatusOK, re)
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
		return
	}

	re, err = r.store.Deactivate(ctx, re)
	if err != nil {
		r.log.Error("failed to deactivate the webhook.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to deactivate the webhook.")
		return
	}

	gb.Response(c, http.StatusOK, re)
}

func isSecure(raw string) bool {
	u, _ := url.Parse(raw)
	if u.Scheme == "https" {
		return true
	}
	return false
}
