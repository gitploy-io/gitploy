package github

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (g *Github) GetUser(ctx context.Context, token string) (*ent.User, error) {
	c := g.Client(ctx, token)

	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return mapGithubUserToUser(u), err
}
