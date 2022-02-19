package api

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	// UserService communicates with the server for users.
	UserService service
)

// GetMe returns the user information.
func (s *UserService) GetMe(ctx context.Context) (*ent.User, error) {
	req, err := s.client.NewRequest(
		"GET",
		"/api/v1/user",
		nil,
	)
	if err != nil {
		return nil, err
	}

	var u *ent.User
	if err := s.client.Do(ctx, req, &u); err != nil {
		return nil, err
	}

	return u, nil
}
