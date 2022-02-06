package main

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/api"
)

var repoListCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "Show own repositories.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Show all repositories.",
		},
		&cli.IntFlag{
			Name:  "page",
			Value: 1,
			Usage: "The page of list.",
		},
		&cli.IntFlag{
			Name:  "per-page",
			Value: 30,
			Usage: "The item count per page.",
		},
	},
	Action: func(cli *cli.Context) error {
		c := buildClient(cli)

		var (
			repos []*ent.Repo
			err   error
		)

		if cli.Bool("all") {
			if repos, err = c.Repos.ListAll(cli.Context); err != nil {
				return err
			}
		} else {
			if repos, err = c.Repos.List(cli.Context, api.RepoListOptions{
				ListOptions: api.ListOptions{Page: cli.Int("page"), PerPage: cli.Int("per-page")},
			}); err != nil {
				return err
			}
		}

		output, err := json.MarshalIndent(repos, "", "  ")
		if err != nil {
			return fmt.Errorf("Failed to marshal: %w", err)
		}

		if q := cli.String("query"); q != "" {
			fmt.Println(gjson.GetBytes(output, q))
			return nil
		}

		fmt.Println(string(output))
		return nil
	},
}
