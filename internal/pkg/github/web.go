package github

import (
	"context"

	"github.com/gitploy-io/gitploy/vo"
	"github.com/google/go-github/v32/github"
)

func (g *Github) GetRemoteUserByToken(ctx context.Context, token string) (*vo.RemoteUser, error) {
	c := g.Client(ctx, token)

	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return mapGithubUserToUser(u), err
}

func (g *Github) ListRemoteOrgsByToken(ctx context.Context, token string) ([]string, error) {
	// TODO: List all orgs.
	orgs, _, err := g.Client(ctx, token).
		Organizations.
		List(ctx, "", &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, o := range orgs {
		ret = append(ret, *o.Login)
	}

	return ret, nil
}
