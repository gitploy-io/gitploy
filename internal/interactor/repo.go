package interactor

import (
	"context"
	"fmt"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) ListReposOfUser(ctx context.Context, u *ent.User, sorted bool, q string, page, perPage int) (repos []*ent.Repo, err error) {
	if sorted {
		repos, err = i.Store.ListSortedReposOfUser(ctx, u, q, page, perPage)
	} else {
		repos, err = i.Store.ListReposOfUser(ctx, u, q, page, perPage)
	}

	return repos, err
}

func (i *Interactor) ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (*ent.Repo, error) {
	hid, err := i.CreateWebhook(ctx, u, r, c)
	if err != nil {
		return nil, fmt.Errorf("failed to create a webhook: %s", err)
	}

	r.WebhookID = hid
	r, err = i.Activate(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to activate the webhook: %w", err)
	}

	return r, nil
}

func (i *Interactor) DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error) {
	err := i.DeleteWebhook(ctx, u, r, r.WebhookID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete the webhook: %w", err)
	}

	r, err = i.Deactivate(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to deactivate the webhook: %w", err)
	}

	return r, nil
}
