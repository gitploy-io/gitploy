package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (i *Interactor) FindRepoByID(ctx context.Context, u *ent.User, id string) (*ent.Repo, error) {
	return i.store.FindRepo(ctx, u, id)
}

func (i *Interactor) FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	return i.store.FindRepoByNamespaceName(ctx, u, namespace, name)
}

func (i *Interactor) FindPermByRepoID(ctx context.Context, u *ent.User, repoID string) (*ent.Perm, error) {
	return i.store.FindPerm(ctx, u, repoID)
}
