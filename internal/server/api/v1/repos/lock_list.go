// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss

package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
)

func (s *LockAPI) List(c *gin.Context) {
	ctx := c.Request.Context()

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	locks, err := s.i.ListLocksOfRepo(ctx, re)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list locks.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, locks)
}
