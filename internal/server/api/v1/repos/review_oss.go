//go:build oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *ReviewAPI) List(c *gin.Context) {
	gb.Response(c, http.StatusOK, make([]*ent.Review, 0))
}

func (s *ReviewAPI) GetMine(c *gin.Context) {
	gb.Response(c, http.StatusNotFound, nil)
}

func (s *ReviewAPI) UpdateMine(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}
