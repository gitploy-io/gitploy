package store

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/user"
)

func (s *Store) ListUsers(ctx context.Context, login string, page, perPage int) ([]*ent.User, error) {
	return s.c.User.
		Query().
		Where(
			user.LoginContains(login),
		).
		WithChatUser().
		Order(ent.Asc(user.FieldLogin)).
		Offset(offset(page, perPage)).
		Limit(perPage).
		All(ctx)
}

// FindUserByID return the user with chat_user,
// but it returns nil if chat_user is not found.
//
// Note that add new indexes if you need to more edges.
func (s *Store) FindUserByID(ctx context.Context, id int64) (*ent.User, error) {
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

func (s *Store) CountUsers(ctx context.Context) (int, error) {
	return s.c.User.
		Query().
		Count(ctx)
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

func (s *Store) DeleteUser(ctx context.Context, u *ent.User) error {
	return s.c.User.
		DeleteOneID(u.ID).
		Exec(ctx)
}
