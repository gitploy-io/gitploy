package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error) {
	return i.scm.ListTags(ctx, u, r, page, perPage)
}

func (i *Interactor) GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error) {
	return i.scm.GetTag(ctx, u, r, tag)
}
