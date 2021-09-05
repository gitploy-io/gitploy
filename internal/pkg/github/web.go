package github

import (
	"context"

	"github.com/gitploy-io/gitploy/vo"
)

func (g *Github) GetUser(ctx context.Context, token string) (*vo.RemoteUser, error) {
	c := g.Client(ctx, token)

	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return mapGithubUserToUser(u), err
}
