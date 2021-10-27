// +build oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

func (r *Repo) ListApprovals(c *gin.Context) {
	gb.Response(c, http.StatusOK, make([]*ent.Approval, 0))
}

func (r *Repo) GetApproval(c *gin.Context) {
	gb.Response(c, http.StatusNotFound, nil)
}

func (r *Repo) GetMyApproval(c *gin.Context) {
	gb.Response(c, http.StatusNotFound, nil)
}

func (r *Repo) CreateApproval(c *gin.Context) {
	gb.ErrorResponse(c, http.StatusPaymentRequired, "It is limited to the community edition.")
}

func (r *Repo) UpdateMyApproval(c *gin.Context) {
	gb.ErrorResponse(c, http.StatusPaymentRequired, "It is limited to the community edition.")
}

func (r *Repo) DeleteApproval(c *gin.Context) {
	gb.ErrorResponse(c, http.StatusPaymentRequired, "It is limited to the community edition.")
}
