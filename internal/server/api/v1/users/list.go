package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (u *UserAPI) List(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		q   = c.DefaultQuery("q", "")
		p   int
		pp  int
		err error
	)

	if p, err = strconv.Atoi(c.DefaultQuery("page", "1")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The page must be number.", err),
		)
	}

	if pp, err = strconv.Atoi(c.DefaultQuery("per_page", "30")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The per_page must be number.", err),
		)
	}

	us, err := u.i.SearchUsers(ctx, &i.SearchUsersOptions{
		Query:       q,
		ListOptions: i.ListOptions{Page: p, PerPage: pp},
	})
	if err != nil {
		u.log.Check(gb.GetZapLogLevel(err), "Failed to list users.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, us)
}
