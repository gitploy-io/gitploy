// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss

package slack

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"go.uber.org/zap"
)

func NewSlack(c *SlackConfig) *Slack {
	s := &Slack{
		host:  c.ServerHost,
		proto: c.ServerProto,
		c:     c.Config,
		i:     c.Interactor,
		log:   zap.L().Named("slack"),
	}

	s.i.SubscribeEvent(func(e *ent.Event) {
		s.Notify(context.Background(), e)
	})

	return s
}
