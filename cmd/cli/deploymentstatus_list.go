package main

import (
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/gitploy-io/gitploy/pkg/api"
)

var deploymentStatusListCommand = &cli.Command{
	Name:      "list",
	Aliases:   []string{"ls"},
	Usage:     "Show the deployment status under the deployment.",
	ArgsUsage: "<owner>/<repo> <number>",
	Action: func(cli *cli.Context) error {
		ns, n, err := splitFullName(cli.Args().First())
		if err != nil {
			return err
		}

		number, err := strconv.Atoi(cli.Args().Get(1))
		if err != nil {
			return err
		}

		c := buildClient(cli)
		dss, err := c.DeploymentStatus.List(cli.Context, ns, n, number, api.ListOptions{})
		if err != nil {
			return err
		}

		return printJson(cli, dss)
	},
}
