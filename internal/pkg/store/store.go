package store

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/pkg/errors"
)

type (
	Store struct {
		c *ent.Client
	}
)

func NewStore(c *ent.Client) *Store {
	return &Store{
		c: c,
	}
}

// WithTx runs callbacks in a transaction.
func (s *Store) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := s.c.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

func offset(page, perPage int) int {
	return (page - 1) * perPage
}
