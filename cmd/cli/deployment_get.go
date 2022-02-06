package main

import (
	"strconv"

	"github.com/urfave/cli/v2"
)

var deploymentGetCommand = &cli.Command{
	Name:      "get",
	Usage:     "Show the deployment",
	ArgsUsage: "<owner>/<repo> <number>",
	Action: func(cli *cli.Context) error {
		// Validate arguments.
		ns, n, err := splitFullName(cli.Args().First())
		if err != nil {
			return err
		}

		number, err := strconv.Atoi(cli.Args().Get(1))
		if err != nil {
			return err
		}

		c := buildClient(cli)
		d, err := c.Deployments.Get(cli.Context, ns, n, number)
		if err != nil {
			return err
		}

		return printJson(cli, d)
	},
}
