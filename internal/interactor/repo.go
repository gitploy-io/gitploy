package interactor

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"go.uber.org/zap"
)

func (i *Interactor) ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo, c *extent.WebhookConfig) (*ent.Repo, error) {
	hid, err := i.SCM.CreateWebhook(ctx, u, r, c)
	if err != nil {
		return nil, fmt.Errorf("failed to create a webhook: %s", err)
	}

	r.WebhookID = hid
	r.OwnerID = u.ID

	r, err = i.Store.Activate(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to activate the webhook: %w", err)
	}

	return r, nil
}

func (i *Interactor) DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error) {
	err := i.SCM.DeleteWebhook(ctx, u, r, r.WebhookID)
	if e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
		i.log.Info("The webhook is not found, skip to delete the webhook.", zap.Int64("id", r.WebhookID))
	} else if err != nil {
		return nil, err
	}

	r, err = i.Store.Deactivate(ctx, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
