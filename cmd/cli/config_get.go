package main

import (
	"github.com/urfave/cli/v2"
)

var configGetCommand = &cli.Command{
	Name:      "get",
	Usage:     "Show the pipeline configurations.",
	ArgsUsage: "<owner>/<repo>",
	Action: func(cli *cli.Context) error {
		ns, n, err := splitFullName(cli.Args().First())
		if err != nil {
			return err
		}

		c := buildClient(cli)
		config, err := c.Config.Get(cli.Context, ns, n)
		if err != nil {
			return err
		}

		return printJson(cli, config)
	},
}
