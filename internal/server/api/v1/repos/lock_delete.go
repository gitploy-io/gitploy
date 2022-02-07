// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss

package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *LocksAPI) Delete(c *gin.Context) {
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

	l, err := s.i.FindLockByID(ctx, id)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to find the lock.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	if err := s.i.DeleteLock(ctx, l); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to delete the lock.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	s.log.Debug("Unlock the env.", zap.String("env", l.Env))
	gb.Response(c, http.StatusOK, nil)
}
