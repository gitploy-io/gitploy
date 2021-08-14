package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"github.com/hanjunlee/gitploy/vo"
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

func (u *User) GetMyUser(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	uv, _ := v.(*ent.User)

	uv, err := u.i.FindUserByID(ctx, uv.ID)
	if err != nil {
		u.log.Error("failed to find the user.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the user.")
		return
	}

	gb.Response(c, http.StatusOK, uv)
}

func (u *User) GetRateLimit(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	uv, _ := v.(*ent.User)

	var (
		rl  *vo.RateLimit
		err error
	)

	if rl, err = u.i.GetRateLimit(ctx, uv); err != nil {
		u.log.Error("It has failed to get the rate-limit.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the rate-limit.")
		return
	}

	gb.Response(c, http.StatusOK, rl)
}
