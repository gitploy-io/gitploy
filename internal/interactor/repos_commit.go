package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error) {
	return i.scm.ListCommits(ctx, u, r, branch, page, perPage)
}

func (i *Interactor) GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error) {
	return i.scm.GetCommit(ctx, u, r, sha)
}

func (i *Interactor) ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error) {
	return i.scm.ListCommitStatuses(ctx, u, r, sha)
}
