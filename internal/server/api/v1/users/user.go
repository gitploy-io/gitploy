package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
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

func (u *User) UpdateUser(c *gin.Context) {
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

func (u *User) DeleteUser(c *gin.Context) {
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

func (u *User) GetMyUser(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	uv, _ := v.(*ent.User)

	uv, err := u.i.FindUserByID(ctx, uv.ID)
	if err != nil {
		u.log.Check(gb.GetZapLogLevel(err), "Failed to find the user.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, uv)
}

func (u *User) GetRateLimit(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	uv, _ := v.(*ent.User)

	var (
		rl  *extent.RateLimit
		err error
	)

	if rl, err = u.i.GetRateLimit(ctx, uv); err != nil {
		u.log.Check(gb.GetZapLogLevel(err), "Failed to get the rate-limit.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, rl)
}
