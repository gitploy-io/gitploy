package main

import (
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/api"
	"github.com/urfave/cli/v2"
)

var (
	repoCommand *cli.Command = &cli.Command{
		Name:  "repo",
		Usage: "Manage repos.",
		Subcommands: []*cli.Command{
			repoListCommand,
		},
	}

	repoListCommand = &cli.Command{
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
				if repos, err = c.Repo.ListAll(cli.Context); err != nil {
					return err
				}
			} else {
				if repos, err = c.Repo.List(cli.Context, api.RepoListOptions{
					ListOptions: api.ListOptions{Page: cli.Int("page"), PerPage: cli.Int("per-page")},
				}); err != nil {
					return err
				}
			}

			// TODO: Enhance output format.
			for _, repo := range repos {
				fmt.Printf("%s/%s\n", repo.Namespace, repo.Name)
			}

			return nil
		},
	}
)
