package github

import (
	"context"

	"github.com/google/go-github/v42/github"
	graphql "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type (
	Github struct {
		baseURL string
	}

	GithubConfig struct {
		BaseURL string
	}
)

func NewGithub(c *GithubConfig) *Github {
	return &Github{
		baseURL: c.BaseURL,
	}
}

func (g *Github) Client(c context.Context, token string) *github.Client {
	tc := oauth2.NewClient(c, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	var client *github.Client
	if g.baseURL != "" {
		client, _ = github.NewEnterpriseClient(g.baseURL, g.baseURL, tc)
	} else {
		client = github.NewClient(tc)
	}

	return client
}

func (g *Github) GraphQLClient(c context.Context, token string) *graphql.Client {
	tc := oauth2.NewClient(c, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	var client *graphql.Client
	if g.baseURL != "" {
		client = graphql.NewEnterpriseClient(g.baseURL, tc)
	} else {
		client = graphql.NewClient(tc)
	}

	return client
}
