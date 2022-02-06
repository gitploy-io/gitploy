package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
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

// splitFullName splits the full name into namespace, and name.
func splitFullName(name string) (string, string, error) {
	ss := strings.Split(name, "/")
	if len(ss) != 2 {
		return "", "", fmt.Errorf("'%s' is invalid format", name)
	}

	return ss[0], ss[1], nil
}

// printJson prints the object as JSON-format.
func printJson(cli *cli.Context, v interface{}) error {
	output, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to marshal: %w", err)
	}

	if query := cli.String("query"); query != "" {
		fmt.Println(gjson.GetBytes(output, query))
		return nil
	}

	fmt.Println(string(output))
	return nil
}
