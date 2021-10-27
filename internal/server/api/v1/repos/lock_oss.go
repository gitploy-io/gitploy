// +build oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

func (r *Repo) ListLocks(c *gin.Context) {
	gb.Response(c, http.StatusOK, make([]*ent.Lock, 0))
}

func (r *Repo) CreateLock(c *gin.Context) {
	gb.ErrorResponse(c, http.StatusPaymentRequired, "It is limited to the community edition.")
}

func (r *Repo) UpdateLock(c *gin.Context) {
	gb.ErrorResponse(c, http.StatusPaymentRequired, "It is limited to the community edition.")
}

func (r *Repo) DeleteLock(c *gin.Context) {
	gb.ErrorResponse(c, http.StatusPaymentRequired, "It is limited to the community edition.")
}
