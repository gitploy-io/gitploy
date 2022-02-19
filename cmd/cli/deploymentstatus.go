package main

import "github.com/urfave/cli/v2"

var deploymentStatusCommand = &cli.Command{
	Name:    "deployment-status",
	Aliases: []string{"ds"},
	Usage:   "Manage deployment statuses.",
	Subcommands: []*cli.Command{
		deploymentStatusListCommand,
		deploymentStatusCreateCommand,
	},
}
