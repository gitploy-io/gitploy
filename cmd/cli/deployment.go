package main

import "github.com/urfave/cli/v2"

var deploymentCommand = &cli.Command{
	Name:    "deployment",
	Aliases: []string{"d"},
	Usage:   "Manage deployments.",
	Subcommands: []*cli.Command{
		deploymentListCommand,
		deploymentGetCommand,
	},
}
