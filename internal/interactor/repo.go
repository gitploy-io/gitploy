package interactor

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"go.uber.org/zap"
)

type (
	// RepoInteractor provides application logic for interacting with repos.
	RepoInteractor struct {
		*service

		WebhookURL    string
		WebhookSSL    bool
		WebhookSecret string
	}

	// RepoStore defines operations for working with repos.
	RepoStore interface {
		CountActiveRepos(ctx context.Context) (int, error)
		CountRepos(ctx context.Context) (int, error)
		ListReposOfUser(ctx context.Context, u *ent.User, opt *ListReposOfUserOptions) ([]*ent.Repo, error)
		FindRepoOfUserByID(ctx context.Context, u *ent.User, id int64) (*ent.Repo, error)
		FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, opt *FindRepoOfUserByNamespaceNameOptions) (*ent.Repo, error)
		FindRepoByID(ctx context.Context, id int64) (*ent.Repo, error)
		SyncRepo(ctx context.Context, r *extent.RemoteRepo) (*ent.Repo, error)
		UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		Activate(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		Deactivate(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
	}

	// ListReposOfUser specifies the optional parameters that
	// search repos.
	ListReposOfUserOptions struct {
		ListOptions

		// Query search the repos contains the query in the namespace or name.
		Query string
		// Sorted instructs the system to sort by the 'deployed_at' field.
		Sorted bool
	}

	// FindRepoOfUserByNamespaceName specifies the parameters to get the repository.
	FindRepoOfUserByNamespaceNameOptions struct {
		Namespace, Name string
	}

	// RepoSCM defines operations for working with remote repos.
	RepoSCM interface {
		ListRemoteRepos(ctx context.Context, u *ent.User) ([]*extent.RemoteRepo, error)
		CreateWebhook(ctx context.Context, u *ent.User, r *ent.Repo, c *extent.WebhookConfig) (int64, error)
		DeleteWebhook(ctx context.Context, u *ent.User, r *ent.Repo, id int64) error
	}
)

// ActivateRepo create a new hook to listen events, and saves
// the hook ID.
func (i *RepoInteractor) ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error) {
	hid, err := i.scm.CreateWebhook(ctx, u, r, &extent.WebhookConfig{
		URL:         i.WebhookURL,
		InsecureSSL: !i.WebhookSSL,
		Secret:      i.WebhookSecret,
	})
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

// DeactivateRepo removes the webhook.
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
