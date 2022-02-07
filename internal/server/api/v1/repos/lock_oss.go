//go:build oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *LocksAPI) List(c *gin.Context) {
	gb.Response(c, http.StatusOK, make([]*ent.Lock, 0))
}

func (s *LocksAPI) Create(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}

func (s *LocksAPI) Update(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}

func (s *LocksAPI) Delete(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}
