// +build oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (r *Repo) ListLocks(c *gin.Context) {
	gb.Response(c, http.StatusOK, make([]*ent.Lock, 0))
}

func (r *Repo) CreateLock(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}

func (r *Repo) UpdateLock(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}

func (r *Repo) DeleteLock(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}
