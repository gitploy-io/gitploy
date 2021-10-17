package store

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/lock"
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
	return s.c.Lock.
		Query().
		Where(lock.RepoID(r.ID)).
		WithUser().
		WithRepo().
		All(ctx)
}

func (s *Store) FindLockOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.Lock, error) {
	return s.c.Lock.
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
		return false, err
	}

	return cnt > 0, nil
}

func (s *Store) FindLockByID(ctx context.Context, id int) (*ent.Lock, error) {
	return s.c.Lock.
		Query().
		Where(
			lock.IDEQ(id),
		).
		WithUser().
		WithRepo().
		Only(ctx)
}

func (s *Store) CreateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error) {
	return s.c.Lock.
		Create().
		SetEnv(l.Env).
		SetNillableExpiredAt(l.ExpiredAt).
		SetRepoID(l.RepoID).
		SetUserID(l.UserID).
		Save(ctx)
}

func (s *Store) UpdateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error) {
	return s.c.Lock.
		UpdateOne(l).
		SetNillableExpiredAt(l.ExpiredAt).
		Save(ctx)
}

func (s *Store) DeleteLock(ctx context.Context, l *ent.Lock) error {
	return s.c.Lock.
		DeleteOne(l).
		Exec(ctx)
}
