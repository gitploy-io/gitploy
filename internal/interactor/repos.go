package interactor

import (
	"context"
	"fmt"

	"github.com/hanjunlee/gitploy/ent"
	reposv1 "github.com/hanjunlee/gitploy/internal/server/api/v1/repos"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) ListRepos(ctx context.Context, u *ent.User, sorted bool, q string, page, perPage int) (repos []*ent.Repo, err error) {
	if sorted {
		repos, err = i.store.ListSortedRepos(ctx, u, q, page, perPage)
	} else {
		repos, err = i.store.ListRepos(ctx, u, q, page, perPage)
	}

	return repos, err
}

func (i *Interactor) FindRepoByID(ctx context.Context, u *ent.User, id string) (*ent.Repo, error) {
	return i.store.FindRepo(ctx, u, id)
}

func (i *Interactor) FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	return i.store.FindRepoByNamespaceName(ctx, u, namespace, name)
}

func (i *Interactor) FindPermByRepoID(ctx context.Context, u *ent.User, repoID string) (*ent.Perm, error) {
	return i.store.FindPerm(ctx, u, repoID)
}

func (i *Interactor) PatchRepo(ctx context.Context, r *ent.Repo, p *reposv1.RepoPayload) (*ent.Repo, error) {
	r.ConfigPath = p.ConfigPath
	return i.store.UpdateRepo(ctx, r)
}

func (i *Interactor) ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (*ent.Repo, error) {
	hid, err := i.scm.CreateWebhook(ctx, u, r, c)
	if err != nil {
		return nil, fmt.Errorf("failed to create a webhook: %s", err)
	}

	r.WebhookID = hid
	r, err = i.store.Activate(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to activate the webhook: %w", err)
	}

	return r, nil
}

func (i *Interactor) DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error) {
	err := i.scm.DeleteWebhook(ctx, u, r, r.WebhookID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete the webhook: %w", err)
	}

	r, err = i.store.Deactivate(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to deactivate the webhook: %w", err)
	}

	return r, nil
}
