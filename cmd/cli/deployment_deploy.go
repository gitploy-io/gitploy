package main

import (
	"github.com/urfave/cli/v2"

	"github.com/gitploy-io/gitploy/pkg/api"
)

var deploymentDeployCommand = &cli.Command{
	Name:      "deploy",
	Usage:     "Deploy a specific ref(branch, SHA, tag) to the environment.",
	ArgsUsage: "<owner>/<repo>",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "type",
			Usage:    "The type of the ref: 'commit', 'branch', or 'tag'.",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "env",
			Usage:    "The name of the environment.",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "ref",
			Usage:    "The specific ref. It can be any named branch, tag, or SHA.",
			Required: true,
		},
	},
	Action: func(cli *cli.Context) error {
		ns, n, err := splitFullName(cli.Args().First())
		if err != nil {
			return err
		}

		c := buildClient(cli)
		d, err := c.Deployment.Create(cli.Context, ns, n, api.DeploymentCreateRequest{
			Type: cli.String("type"),
			Ref:  cli.String("ref"),
			Env:  cli.String("env"),
		})
		if err != nil {
			return err
		}

		return printJson(d, cli.String("query"))
	},
}
