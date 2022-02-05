package main

import (
	"github.com/urfave/cli/v2"
)

var repoCommand *cli.Command = &cli.Command{
	Name:  "repo",
	Usage: "Manage repositories.",
	Subcommands: []*cli.Command{
		repoListCommand,
		repoGetCommand,
		repoUpdateCommand,
	},
}
