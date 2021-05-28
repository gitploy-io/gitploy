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

	u, _ := r.store.FindUserByHash(ctx, c.GetString(gb.KeySession))

	repos, err := r.store.ListRepos(ctx, u, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list repositories.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list repositories.")
		return
	}

	gb.Response(c, http.StatusOK, mapReposToRepoDatas(repos))
}

func (r *Repo) GetRepo(c *gin.Context) {
	var (
		id = c.Param("id")
	)
	ctx := c.Request.Context()

	u, _ := r.store.FindUserByHash(ctx, c.GetString(gb.KeySession))

	repo, err := r.store.FindRepo(ctx, u, id)
	if err != nil {
		r.log.Error("failed to get the repository.", zap.String("id", id), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the repository.")
		return
	}

	gb.Response(c, http.StatusOK, mapRepoToRepoData(repo))
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
