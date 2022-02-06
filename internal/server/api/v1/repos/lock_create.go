package repos

import (
	"net/http"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/internal/server/global"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	LockPostPayload struct {
		Env       string  `json:"env"`
		ExpiredAt *string `json:"expired_at,omitempty"`
	}
)

func (s *LocksAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()

	p := &LockPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		s.log.Warn("Failed to bind the payload.", zap.Error(err))
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", err))
		return
	}

	var expiredAt *time.Time
	if p.ExpiredAt != nil {
		exp, err := time.Parse(time.RFC3339, *p.ExpiredAt)
		if err != nil {
			gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"expired_at\" parameter, RFC3339 format only.", err))
			return
		}

		expiredAt = pointer.ToTime(exp.UTC())
	}

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	vu, _ := c.Get(global.KeyUser)
	u := vu.(*ent.User)

	config, err := s.i.GetConfig(ctx, u, re)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to get the configuration.").Write(zap.Error(err))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	}

	var env *extent.Env
	if env = config.GetEnv(p.Env); env == nil {
		s.log.Warn("The environment is not found.", zap.String("env", p.Env))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, e.NewError(e.ErrorCodeConfigUndefinedEnv, nil))
		return
	}

	l, err := s.i.CreateLock(ctx, &ent.Lock{
		Env:       env.Name,
		ExpiredAt: expiredAt,
		UserID:    u.ID,
		RepoID:    re.ID,
	})
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to create a new lock.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if nl, err := s.i.FindLockByID(ctx, l.ID); err == nil {
		l = nl
	}

	s.log.Info("Lock the environment.", zap.String("repo", re.GetFullName()), zap.String("env", p.Env), zap.String("login", u.Login))
	gb.Response(c, http.StatusCreated, l)
}
