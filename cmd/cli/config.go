package main

import "github.com/urfave/cli/v2"

var configCommand = &cli.Command{
	Name:  "config",
	Usage: "Manage config.",
	Subcommands: []*cli.Command{
		configGetCommand,
	},
}
