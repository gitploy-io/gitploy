package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
)

func (s *RepoAPI) Get(c *gin.Context) {
	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	gb.Response(c, http.StatusOK, repo)
}
