package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"

	"github.com/gitploy-io/gitploy/pkg/api"
)

// buildClient returns a client to interact with a server.
func buildClient(cli *cli.Context) *api.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cli.String("token")},
	)
	tc := oauth2.NewClient(cli.Context, ts)

	return api.NewClient(cli.String("host"), tc)
}

func splitFullName(name string) (string, string, error) {
	ss := strings.Split(name, "/")
	if len(ss) != 2 {
		return "", "", fmt.Errorf("'%s' is invalid format", name)
	}

	return ss[0], ss[1], nil
}