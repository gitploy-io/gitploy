package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"go.uber.org/zap"
)

type (
	Repo struct {
		store Store
		scm   SCM
		log   *zap.Logger
	}

	repoData struct {
		*ent.Repo
		FullName string `json:"full_name"`
	}
)

func NewRepo(store Store, scm SCM) *Repo {
	return &Repo{
		store: store,
		scm:   scm,
		log:   zap.L().Named("repos"),
	}
}

func (r *Repo) ListRepos(c *gin.Context) {
	var (
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)

	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	repos, err := r.store.ListRepos(ctx, u, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list repositories.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list repositories.")
		return
	}

	gb.Response(c, http.StatusOK, mapReposToRepoDatas(repos))
}

func (r *Repo) GetRepo(c *gin.Context) {
	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	gb.Response(c, http.StatusOK, mapRepoToRepoData(repo))
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
