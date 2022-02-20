package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	// extendedUserData includes the 'hash' field.
	extendedUserData struct {
		*ent.User

		Hash string `json:"hash"`
	}
)

func (u *UserAPI) GetMe(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	uv, _ := v.(*ent.User)

	uv, err := u.i.FindUserByID(ctx, uv.ID)
	if err != nil {
		u.log.Check(gb.GetZapLogLevel(err), "Failed to find the user.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, extendedUserData{
		User: uv,
		Hash: uv.Hash,
	})
}
