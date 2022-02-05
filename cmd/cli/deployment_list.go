package main

import (
	"github.com/urfave/cli/v2"

	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/pkg/api"
)

var deploymentListCommand = &cli.Command{
	Name:      "list",
	Aliases:   []string{"ls"},
	Usage:     "Show the deployments under the repository.",
	ArgsUsage: "<owner>/<repo>",
	Flags: []cli.Flag{
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
		&cli.StringFlag{
			Name:  "env",
			Usage: "The name of environment. It only shows deployments for the environment.",
		},
		&cli.StringFlag{
			Name:  "status",
			Usage: "The deployment status: 'waiting', 'created', 'queued', 'running', 'success', or 'failure'. It only shows deployments the status is matched. ",
		},
	},
	Action: func(cli *cli.Context) error {
		c := buildClient(cli)

		ns, n, err := splitFullName(cli.Args().First())
		if err != nil {
			return err
		}

		ds, err := c.Deployment.List(cli.Context, ns, n, api.DeploymentListOptions{
			ListOptions: api.ListOptions{Page: cli.Int("page"), PerPage: cli.Int("per-page")},
			Env:         cli.String("env"),
			Status:      deployment.Status(cli.String("status")),
		})
		if err != nil {
			return err
		}

		return printJson(ds, cli.String("query"))
	},
}
