//go:build oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *ReviewsAPI) List(c *gin.Context) {
	gb.Response(c, http.StatusOK, make([]*ent.Review, 0))
}

func (s *ReviewsAPI) GetMine(c *gin.Context) {
	gb.Response(c, http.StatusNotFound, nil)
}

func (s *ReviewsAPI) UpdateMine(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}
