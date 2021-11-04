// +build oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (r *Repo) ListReviews(c *gin.Context) {
	gb.Response(c, http.StatusOK, make([]*ent.Review, 0))
}

func (r *Repo) GetUserReview(c *gin.Context) {
	gb.Response(c, http.StatusNotFound, nil)
}

func (r *Repo) UpdateUserReview(c *gin.Context) {
	gb.ResponseWithError(
		c,
		e.NewError(e.ErrorCodeLicenseRequired, nil),
	)
}
