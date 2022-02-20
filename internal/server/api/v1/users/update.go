package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	userPatchPayload struct {
		Admin *bool `json:"admin"`
	}
)

func (u *UserAPI) Update(c *gin.Context) {
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

	p := &userPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		u.log.Warn("It has failed to binding the payload.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to binding the payload.", err),
		)
		return
	}

	du, err := u.i.FindUserByID(ctx, id)
	if err != nil {
		u.log.Check(gb.GetZapLogLevel(err), "Failed to find the user.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if p.Admin != nil {
		du.Admin = *p.Admin
		if du, err = u.i.UpdateUser(ctx, du); err != nil {
			u.log.Check(gb.GetZapLogLevel(err), "Failed to update the user.").Write(zap.Error(err))
			gb.ResponseWithError(c, err)
			return
		}
	}

	gb.Response(c, http.StatusOK, du)
}
