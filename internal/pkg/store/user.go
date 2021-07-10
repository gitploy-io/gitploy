package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/user"
)

// FindUserByID return the user with chat_user,
// but it returns nil if chat_user is not found.
//
// Note that add new indexes if you need to more edges.
func (s *Store) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	return s.c.User.
		Query().
		Where(
			user.IDEQ(id),
		).
		WithChatUser().
		First(ctx)
}

func (s *Store) FindUserByHash(ctx context.Context, hash string) (*ent.User, error) {
	return s.c.User.
		Query().
		Where(
			user.HashEQ(hash),
		).
		WithChatUser().
		Only(ctx)
}

func (s *Store) FindUserByLogin(ctx context.Context, login string) (*ent.User, error) {
	return s.c.User.
		Query().
		Where(
			user.LoginEQ(login),
		).
		WithChatUser().
		Only(ctx)
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
	return s.c.User.UpdateOneID(u.ID).
		SetLogin(u.Login).
		SetAvatar(u.Avatar).
		SetAdmin(u.Admin).
		SetToken(u.Token).
		SetRefresh(u.Refresh).
		SetExpiry(u.Expiry).
		Save(ctx)
}
