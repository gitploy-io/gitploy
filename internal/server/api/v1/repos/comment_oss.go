// +build oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

func (r *Repo) ListComments(c *gin.Context) {
	gb.Response(c, http.StatusOK, make([]*ent.Comment, 0))
}

func (r *Repo) GetComment(c *gin.Context) {
	gb.Response(c, http.StatusNotFound, nil)
}

func (r *Repo) CreateComment(c *gin.Context) {
	gb.ErrorResponse(c, http.StatusPaymentRequired, "It is limited to the community edition.")
}
