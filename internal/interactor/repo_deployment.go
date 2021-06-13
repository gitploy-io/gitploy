package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error) {
	return i.scm.GetConfig(ctx, u, r)
}
