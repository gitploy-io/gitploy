package store

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/perm"
	"github.com/gitploy-io/gitploy/ent/repo"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/vo"
)

func (s *Store) CountActiveRepos(ctx context.Context) (int, error) {
	cnt, err := s.c.Repo.
		Query().
		Where(repo.ActiveEQ(true)).
		Count(ctx)
	if err != nil {
		return 0, e.NewError(e.ErrorCodeInternalError, err)
	}

	return cnt, nil
}

func (s *Store) CountRepos(ctx context.Context) (int, error) {
	cnt, err := s.c.Repo.
		Query().
		Count(ctx)
	if err != nil {
		return 0, e.NewError(e.ErrorCodeInternalError, err)
	}

	return cnt, nil
}

func (s *Store) ListReposOfUser(ctx context.Context, u *ent.User, q, namespace, name string, sorted bool, page, perPage int) ([]*ent.Repo, error) {
	// Build the query with parameters.
	qry := s.c.Repo.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(perm.Table)
			s.
				Join(t).
				On(t.C(perm.FieldRepoID), s.C(repo.FieldID)).
				Where(sql.EQ(t.C(perm.FieldUserID), u.ID))
		}).
		WithOwner().
		Limit(perPage).
		Offset(offset(page, perPage))

	if q != "" {
		qry = qry.Where(
			repo.Or(
				repo.NamespaceContains(q),
				repo.NameContains(q),
			),
		)
	}

	if namespace != "" {
		qry = qry.Where(repo.NamespaceEQ(namespace))
	}

	if name != "" {
		qry = qry.Where(repo.NameEQ(name))
	}

	if sorted {
		qry = qry.Order(
			ent.Desc(repo.FieldLatestDeployedAt),
		)
	}

	repos, err := qry.All(ctx)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	for _, r := range repos {
		deployments, err := r.
			QueryDeployments().
			Order(ent.Desc(deployment.FieldID)).
			Limit(3).
			WithUser().
			All(ctx)
		if err != nil {
			return nil, e.NewError(e.ErrorCodeInternalError, err)
		}

		r.Edges.Deployments = deployments
	}
	return repos, nil
}

func (s *Store) FindRepoByID(ctx context.Context, id int64) (*ent.Repo, error) {
	r, err := s.c.Repo.
		Query().
		Where(
			repo.IDEQ(id),
		).
		WithOwner().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeNotFound, "The repository is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return r, nil
}

func (s *Store) FindRepoOfUserByID(ctx context.Context, u *ent.User, id int64) (*ent.Repo, error) {
	r, err := s.c.Repo.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(perm.Table)
			s.
				Join(t).
				On(t.C(perm.FieldRepoID), s.C(repo.FieldID)).
				Where(sql.EQ(t.C(perm.FieldUserID), u.ID))
		}).
		Where(
			repo.IDEQ(id),
		).
		WithOwner().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeNotFound, "The repository is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return r, nil
}

func (s *Store) FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	r, err := s.c.Repo.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(perm.Table)
			s.
				Join(t).
				On(t.C(perm.FieldRepoID), s.C(repo.FieldID)).
				Where(sql.EQ(t.C(perm.FieldUserID), u.ID))
		}).
		Where(
			repo.And(
				repo.NamespaceEQ(namespace),
				repo.NameEQ(name),
			),
		).
		WithOwner().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeNotFound, "The repository is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return r, nil
}

func (s *Store) SyncRepo(ctx context.Context, r *vo.RemoteRepo) (*ent.Repo, error) {
	ret, err := s.c.Repo.
		Create().
		SetID(r.ID).
		SetNamespace(r.Namespace).
		SetName(r.Name).
		SetDescription(r.Description).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeUnprocessableEntity,
			fmt.Sprintf("The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return ret, nil
}

func (s *Store) UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	ret, err := s.c.Repo.
		UpdateOne(r).
		SetConfigPath(r.ConfigPath).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeUnprocessableEntity,
			fmt.Sprintf("The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return ret, nil
}

func (s *Store) Activate(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	ret, err := s.c.Repo.
		UpdateOne(r).
		SetActive(true).
		SetWebhookID(r.WebhookID).
		SetOwnerID(r.OwnerID).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeUnprocessableEntity,
			fmt.Sprintf("The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return ret, nil
}

func (s *Store) Deactivate(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	ret, err := s.c.Repo.
		UpdateOne(r).
		SetActive(false).
		SetWebhookID(0).
		SetOwnerID(0).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeUnprocessableEntity,
			fmt.Sprintf("The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return ret, nil
}
