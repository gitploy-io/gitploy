package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gitploy",
		Usage: "Command line utility.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "host",
				Aliases:  []string{"H"},
				Required: true,
				Usage:    "The host of server. It must have a trailing slash (i.e., '/').",
				EnvVars:  []string{"GITPLOY_SERVER_HOST"},
			},
			&cli.StringFlag{
				Name:     "token",
				Aliases:  []string{"T"},
				Required: true,
				Usage:    "The authorization token.",
				EnvVars:  []string{"GITPLOY_TOKEN"},
			},
			&cli.StringFlag{
				Name:  "query",
				Usage: "A GJSON query to use in filtering the response data",
			},
		},
		Commands: []*cli.Command{
			repoCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
