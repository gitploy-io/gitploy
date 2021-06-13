package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (s *Store) FindUser() (*ent.User, error) {
	return s.c.User.Get(context.Background(), "17633736")
}
