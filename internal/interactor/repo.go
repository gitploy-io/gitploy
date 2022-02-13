package interactor

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"go.uber.org/zap"
)

type RepoInteractor service

func (i *RepoInteractor) ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo, c *extent.WebhookConfig) (*ent.Repo, error) {
	hid, err := i.scm.CreateWebhook(ctx, u, r, c)
	if err != nil {
		return nil, fmt.Errorf("failed to create a webhook: %s", err)
	}

	r.WebhookID = hid
	r.OwnerID = u.ID

	r, err = i.store.Activate(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to activate the webhook: %w", err)
	}

	return r, nil
}

func (i *RepoInteractor) DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error) {
	err := i.scm.DeleteWebhook(ctx, u, r, r.WebhookID)
	if e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
		i.log.Info("The webhook is not found, skip to delete the webhook.", zap.Int64("id", r.WebhookID))
	} else if err != nil {
		return nil, err
	}

	r, err = i.store.Deactivate(ctx, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
