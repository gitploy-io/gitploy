package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	User struct {
		i   Interactor
		log *zap.Logger
	}

	userPatchPayload struct {
		Admin *bool `json:"admin"`
	}
)

func NewUser(i Interactor) *User {
	return &User{
		i:   i,
		log: zap.L().Named("users"),
	}
}

func (u *User) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		q   = c.DefaultQuery("q", "")
		p   int
		pp  int
		err error
	)

	if p, err = strconv.Atoi(c.DefaultQuery("page", "1")); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format \"page\".")
	}

	if pp, err = strconv.Atoi(c.DefaultQuery("per_page", "30")); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format \"per_page\".")
	}

	us, err := u.i.ListUsers(ctx, q, p, pp)
	if err != nil {
		u.log.Error("It has failed to list users.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list users.")
		return
	}

	gb.Response(c, http.StatusOK, us)
}

func (u *User) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		u.log.Error("Invalid ID of user.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid ID of user.")
		return
	}

	p := &userPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		u.log.Error("It has failed to binding the payload.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to binding the payload.")
		return
	}

	du, err := u.i.FindUserByID(ctx, id)
	if ent.IsNotFound(err) {
		u.log.Warn("The deleting user is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The deleting user is not found.")
		return
	} else if err != nil {
		u.log.Error("It has failed to get the user.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the user.")
		return
	}

	if p.Admin != nil {
		du.Admin = *p.Admin
		if du, err = u.i.UpdateUser(ctx, du); err != nil {
			u.log.Error("It has failed to patch the user.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to delete the user.")
			return
		}
	}

	gb.Response(c, http.StatusOK, du)
}

func (u *User) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		u.log.Error("Invalid ID of user.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid ID of user.")
		return
	}

	du, err := u.i.FindUserByID(ctx, id)
	if ent.IsNotFound(err) {
		u.log.Warn("The deleting user is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The deleting user is not found.")
		return
	} else if err != nil {
		u.log.Error("It has failed to get the user.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the user.")
		return
	}

	if err := u.i.DeleteUser(ctx, du); err != nil {
		u.log.Error("It has failed to delete the user.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to delete the user.")
		return
	}

	c.Status(http.StatusOK)
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
