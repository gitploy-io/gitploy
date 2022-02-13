package main

import (
	"strconv"

	"github.com/urfave/cli/v2"
)

var deploymentUpdateCommand = &cli.Command{
	Name:      "update",
	Usage:     "Trigger the deployment which has approved by reviews.",
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
		d, err := c.Deployment.Update(cli.Context, ns, n, number)
		if err != nil {
			return err
		}

		return printJson(cli, d)
	},
}
