package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (i *Interactor) FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	return i.store.FindRepoByNamespaceName(ctx, u, namespace, name)
}
