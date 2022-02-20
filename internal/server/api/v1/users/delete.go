package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (u *UserAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		u.log.Warn("The id must be number.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The id must be number.", err),
		)
		return
	}

	du, err := u.i.FindUserByID(ctx, id)
	if err != nil {
		u.log.Check(gb.GetZapLogLevel(err), "Failed to find the user.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if err := u.i.DeleteUser(ctx, du); err != nil {
		u.log.Check(gb.GetZapLogLevel(err), "Failed to delete the user.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	c.Status(http.StatusOK)
}
