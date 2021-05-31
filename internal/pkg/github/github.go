package github

import (
	"context"

	"github.com/google/go-github/v32/github"
	graphql "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type (
	Github struct{}
)

func NewGithub() *Github {
	return &Github{}
}

func (g *Github) Client(c context.Context, token string) *github.Client {
	tc := oauth2.NewClient(c, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	return github.NewClient(tc)
}

func (g *Github) GraphQLClient(c context.Context, token string) *graphql.Client {
	tc := oauth2.NewClient(c, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	return graphql.NewClient(tc)
}
