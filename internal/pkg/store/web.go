package store

import (
	"context"

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

func (s *Store) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	return s.c.User.Get(ctx, id)
}

func (s *Store) CreateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	return s.c.User.Create().
		SetID(u.ID).
		SetLogin(u.Login).
		SetAvatar(u.Avatar).
		SetAdmin(u.Admin).
		SetToken(u.Token).
		SetRefresh(u.Refresh).
		SetExpiry(u.Expiry).
		Save(ctx)
}

func (s *Store) UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	return s.c.User.UpdateOne(u).
		SetLogin(u.Login).
		SetAvatar(u.Avatar).
		SetAdmin(u.Admin).
		SetToken(u.Token).
		SetRefresh(u.Refresh).
		SetExpiry(u.Expiry).
		Save(ctx)
}
