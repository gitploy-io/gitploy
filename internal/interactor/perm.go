package interactor

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"go.uber.org/zap"
)

// PermStore defines operations for working with perms.
type PermStore interface {
	ListPerms(ctx context.Context, opt *ListOptions) ([]*ent.Perm, error)
	ListPermsOfRepo(ctx context.Context, r *ent.Repo, opt *ListPermsOfRepoOptions) ([]*ent.Perm, error)
	FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)
	CreatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error)
	UpdatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error)
	DeletePermsOfUserLessThanSyncedAt(ctx context.Context, u *ent.User, t time.Time) (int, error)
	DeletePerm(ctx context.Context, p *ent.Perm) error
}

type ListPermsOfRepoOptions struct {
	ListOptions

	// Query search the 'login' contains the query.
	Query string
}

type PermInteractor struct {
	*service

	orgEntries []string
}

// ResyncPerms delete all permissions not included in the organization entries.
func (i *PermInteractor) ResyncPerms(ctx context.Context) error {
	const perPage = 100

	page := 1
	for {
		perms, err := i.store.ListPerms(ctx, &ListOptions{
			Page:    page,
			PerPage: perPage,
		})
		if err != nil {
			return err
		}

		for _, p := range perms {
			if p.Edges.Repo == nil {
				i.log.Warn("Failed to eager loading for the perm.", zap.Int("perm_id", p.ID))
				continue
			}

			if i.matchOrg(p.Edges.Repo.Namespace) {
				continue
			}

			i.log.Debug("Delete the perm.", zap.String("repo_fullname", p.Edges.Repo.GetFullName()))
			if err := i.store.DeletePerm(ctx, p); err != nil {
				i.log.Error("Failed to delete the perm.", zap.Error(err))
			}
		}

		// Stop the loop if it is the last page.
		if len(perms) < perPage {
			break
		}

		page += 1
	}

	return nil
}

func (i *PermInteractor) matchOrg(namespace string) bool {
	for _, org := range i.orgEntries {
		if namespace == org {
			return true
		}
	}

	return false
}
