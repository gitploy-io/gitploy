package interactor

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"go.uber.org/zap"
)

type (
	// LockInteractor provides application logic for interacting with users.
	LockInteractor service

	// LockStore defines operations for working with locks.
	LockStore interface {
		ListExpiredLocksLessThanTime(ctx context.Context, t time.Time) ([]*ent.Lock, error)
		ListLocksOfRepo(ctx context.Context, r *ent.Repo) ([]*ent.Lock, error)
		FindLockOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.Lock, error)
		HasLockOfRepoForEnv(ctx context.Context, r *ent.Repo, env string) (bool, error)
		FindLockByID(ctx context.Context, id int) (*ent.Lock, error)
		CreateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		UpdateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		DeleteLock(ctx context.Context, l *ent.Lock) error
	}
)

func (i *LockInteractor) runAutoUnlock(stop <-chan struct{}) {
	ctx := context.Background()

	ticker := time.NewTicker(time.Minute)
L:
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				ticker.Stop()
				break L
			}
		case t := <-ticker.C:
			ls, err := i.store.ListExpiredLocksLessThanTime(ctx, t.UTC())
			if err != nil {
				i.log.Error("It has failed to read expired locks.", zap.Error(err))
				continue
			}

			for _, l := range ls {
				i.store.DeleteLock(ctx, l)
				i.log.Debug("Delete the expired lock.", zap.Int("id", l.ID), zap.Time("time", t))
			}
		}
	}
}
