package store

import (
	"context"
	"fmt"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/lock"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *Store) ListExpiredLocksLessThanTime(ctx context.Context, t time.Time) ([]*ent.Lock, error) {
	return s.c.Lock.
		Query().
		Where(
			lock.ExpiredAtLT(t),
		).
		WithUser().
		WithRepo().
		All(ctx)
}

func (s *Store) ListLocksOfRepo(ctx context.Context, r *ent.Repo) ([]*ent.Lock, error) {
	ls, err := s.c.Lock.
		Query().
		Where(lock.RepoID(r.ID)).
		WithUser().
		WithRepo().
		All(ctx)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return ls, nil
}

func (s *Store) FindLockOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.Lock, error) {
	l, err := s.c.Lock.
		Query().
		Where(
			lock.And(
				lock.RepoID(r.ID),
				lock.EnvEQ(env),
			),
		).
		WithUser().
		WithRepo().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The lock is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return l, nil
}

func (s *Store) HasLockOfRepoForEnv(ctx context.Context, r *ent.Repo, env string) (bool, error) {
	cnt, err := s.c.Lock.
		Query().
		Where(
			lock.And(
				lock.RepoID(r.ID),
				lock.EnvEQ(env),
			),
		).
		WithUser().
		WithRepo().
		Count(ctx)
	if err != nil {
		return false, e.NewError(e.ErrorCodeInternalError, err)
	}

	return cnt > 0, nil
}

func (s *Store) FindLockByID(ctx context.Context, id int) (*ent.Lock, error) {
	l, err := s.c.Lock.
		Query().
		Where(
			lock.IDEQ(id),
		).
		WithUser().
		WithRepo().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The lock is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return l, nil
}

func (s *Store) CreateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error) {
	l, err := s.c.Lock.
		Create().
		SetEnv(l.Env).
		SetNillableExpiredAt(l.ExpiredAt).
		SetRepoID(l.RepoID).
		SetUserID(l.UserID).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("Failed to create a lock. The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return l, nil
}

func (s *Store) UpdateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error) {
	l, err := s.c.Lock.
		UpdateOne(l).
		SetNillableExpiredAt(l.ExpiredAt).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("Failed to update the lock. The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return l, nil
}

func (s *Store) DeleteLock(ctx context.Context, l *ent.Lock) error {
	return s.c.Lock.
		DeleteOne(l).
		Exec(ctx)
}
