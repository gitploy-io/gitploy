package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
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
