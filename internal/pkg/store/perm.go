package store

import (
	"context"
	"fmt"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/perm"
	"github.com/gitploy-io/gitploy/model/ent/repo"
	"github.com/gitploy-io/gitploy/model/ent/user"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *Store) ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error) {
	perms, err := s.c.Perm.
		Query().
		Where(
			perm.And(
				perm.HasRepoWith(repo.IDEQ(r.ID)),
				perm.HasUserWith(user.LoginContains(q)),
			),
		).
		WithRepo().
		WithUser().
		Limit(perPage).
		Offset(offset(page, perPage)).
		All(ctx)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return perms, nil
}

func (s *Store) FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error) {
	p, err := s.c.Perm.
		Query().
		Where(
			perm.And(
				perm.HasUserWith(user.IDEQ(u.ID)),
				perm.HasRepoWith(repo.IDEQ(r.ID)),
			),
		).
		WithRepo().
		WithUser().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The user has no permission for the repository.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return p, nil
}

func (s *Store) CreatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
	perm, err := s.c.Perm.
		Create().
		SetRepoPerm(p.RepoPerm).
		SetSyncedAt(p.SyncedAt).
		SetUserID(p.UserID).
		SetRepoID(p.RepoID).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return perm, nil
}

func (s *Store) UpdatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
	perm, err := s.c.Perm.
		UpdateOne(p).
		SetRepoPerm(p.RepoPerm).
		SetSyncedAt(p.SyncedAt).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return perm, nil
}

func (s *Store) DeletePermsOfUserLessThanSyncedAt(ctx context.Context, u *ent.User, t time.Time) (int, error) {
	var (
		cnt int
		err error
	)

	if err = s.WithTx(ctx, func(tx *ent.Tx) error {
		cnt, err = tx.Perm.
			Delete().
			Where(
				perm.Or(
					perm.And(
						perm.UserIDEQ(u.ID),
						perm.SyncedAtLT(t),
					),
					perm.And(
						perm.UserIDEQ(u.ID),
						perm.SyncedAtIsNil(),
					),
				),
			).
			Exec(ctx)
		return err
	}); err != nil {
		return 0, e.NewError(e.ErrorCodeInternalError, err)
	}

	return cnt, nil
}
