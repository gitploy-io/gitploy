package store

import (
	"github.com/hanjunlee/gitploy/ent"
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
