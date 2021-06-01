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
