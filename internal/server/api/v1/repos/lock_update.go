package repos

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	LockPatchPayload struct {
		ExpiredAt *string `json:"expired_at,omitempty"`
	}
)

func (s *LockAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(c.Param("lockID")); err != nil {
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The ID must be number.", nil))
		return
	}

	p := &LockPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", nil))
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

	l, err := s.i.FindLockByID(ctx, id)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "The lock is not found.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if p.ExpiredAt != nil {
		l.ExpiredAt = expiredAt
		s.log.Debug("Update the expired_at of the lock.", zap.Int("id", l.ID), zap.Timep("expired_at", l.ExpiredAt))
	}

	if _, err := s.i.UpdateLock(ctx, l); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to update the lock.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if nl, err := s.i.FindLockByID(ctx, l.ID); err == nil {
		l = nl
	}

	s.log.Info("Patch the environment lock.", zap.Int("lock_id", l.ID))
	gb.Response(c, http.StatusOK, l)
}
