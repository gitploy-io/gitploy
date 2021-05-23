package github

import (
	"context"

	"github.com/google/go-github/v32/github"
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
