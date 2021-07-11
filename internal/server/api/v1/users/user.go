package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	User struct {
		i   Interactor
		log *zap.Logger
	}
)

func NewUser(i Interactor) *User {
	return &User{
		i:   i,
		log: zap.L().Named("users"),
	}
}

func (u *User) Me(c *gin.Context) {
	v, _ := c.Get(gb.KeyUser)
	uv, _ := v.(*ent.User)

	ctx := c.Request.Context()

	uv, err := u.i.FindUserByID(ctx, uv.ID)
	if err != nil {
		u.log.Error("failed to find the user.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the user.")
		return
	}

	gb.Response(c, http.StatusOK, uv)
}
