// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package repos

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/extent"
	"github.com/gitploy-io/gitploy/internal/server/global"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	lockPostPayload struct {
		Env       string  `json:"env"`
		ExpiredAt *string `json:"expired_at,omitempty"`
	}

	lockPatchPayload struct {
		ExpiredAt *string `json:"expired_at,omitempty"`
	}
)

func (r *Repo) ListLocks(c *gin.Context) {
	ctx := c.Request.Context()

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	locks, err := r.i.ListLocksOfRepo(ctx, re)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to list locks.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, locks)
}

func (r *Repo) CreateLock(c *gin.Context) {
	ctx := c.Request.Context()

	p := &lockPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Error("It has failed to bind the payload.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", err),
		)
		return
	}

	var (
		expiredAt *time.Time
	)
	if p.ExpiredAt != nil {
		exp, err := time.Parse(time.RFC3339, *p.ExpiredAt)
		if err != nil {
			gb.ResponseWithError(
				c,
				e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"expired_at\" parameter, RFC3339 format only.", err),
			)
			return
		}

		expiredAt = pointer.ToTime(exp.UTC())
	}

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	vu, _ := c.Get(global.KeyUser)
	u := vu.(*ent.User)

	config, err := r.i.GetConfig(ctx, u, re)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to get the configuration.").Write(zap.Error(err))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, err)
		return
	}

	var env *extent.Env
	if env = config.GetEnv(p.Env); env == nil {
		r.log.Warn("The environment is not found.", zap.String("env", p.Env))
		gb.ResponseWithStatusAndError(c, http.StatusUnprocessableEntity, e.NewError(e.ErrorCodeConfigUndefinedEnv, nil))
		return
	}

	l, err := r.i.CreateLock(ctx, &ent.Lock{
		Env:       env.Name,
		ExpiredAt: expiredAt,
		UserID:    u.ID,
		RepoID:    re.ID,
	})
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to create a new lock.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if nl, err := r.i.FindLockByID(ctx, l.ID); err == nil {
		l = nl
	}

	r.log.Info("Lock the environment.", zap.String("repo", re.GetFullName()), zap.String("env", p.Env), zap.String("login", u.Login))
	gb.Response(c, http.StatusCreated, l)
}

func (r *Repo) UpdateLock(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(c.Param("lockID")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The ID must be number.", nil),
		)
		return
	}

	p := &lockPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the payload.", nil),
		)
		return
	}

	var expiredAt *time.Time
	if p.ExpiredAt != nil {
		exp, err := time.Parse(time.RFC3339, *p.ExpiredAt)
		if err != nil {
			gb.ResponseWithError(
				c,
				e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"expired_at\" parameter, RFC3339 format only.", err),
			)
			return
		}

		expiredAt = pointer.ToTime(exp.UTC())
	}

	l, err := r.i.FindLockByID(ctx, id)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "The lock is not found.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if p.ExpiredAt != nil {
		l.ExpiredAt = expiredAt
		r.log.Debug("Update the expired_at of the lock.", zap.Int("id", l.ID), zap.Timep("expired_at", l.ExpiredAt))
	}

	if _, err := r.i.UpdateLock(ctx, l); err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to update the lock.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if nl, err := r.i.FindLockByID(ctx, l.ID); err == nil {
		l = nl
	}

	r.log.Info("Patch the environment lock.", zap.Int("lock_id", l.ID))
	gb.Response(c, http.StatusOK, l)
}

func (r *Repo) DeleteLock(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(c.Param("lockID")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The ID must be number.", nil),
		)
		return
	}

	l, err := r.i.FindLockByID(ctx, id)
	if err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to find the lock.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if err := r.i.DeleteLock(ctx, l); err != nil {
		r.log.Check(gb.GetZapLogLevel(err), "Failed to delete the lock.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	r.log.Debug("Unlock the env.", zap.String("env", l.Env))
	gb.Response(c, http.StatusOK, nil)
}
