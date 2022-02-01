package main

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"

	"github.com/gitploy-io/gitploy/pkg/api"
)

func buildClient(cli *cli.Context) *api.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cli.String("token")},
	)
	tc := oauth2.NewClient(cli.Context, ts)

	return api.NewClient(cli.String("host"), tc)
}
