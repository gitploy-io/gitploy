package main

import (
	"encoding/json"
	"fmt"

	"github.com/AlekSi/pointer"
	"github.com/gitploy-io/gitploy/pkg/api"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
)

var repoUpdateCommand = &cli.Command{
	Name:      "update",
	Usage:     "Update the repository.",
	ArgsUsage: "<owner>/<repo>",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"C"},
			Usage:   "The path of configuration file.",
		},
	},
	Action: func(cli *cli.Context) error {
		ns, n, err := splitFullName(cli.Args().First())
		if err != nil {
			return err
		}

		// Build the request body.
		req := api.RepoUpdateRequest{}
		if config := cli.String("config"); config != "" {
			req.ConfigPath = pointer.ToString(config)
		}

		c := buildClient(cli)
		repo, err := c.Repo.Update(cli.Context, ns, n, req)
		if err != nil {
			return err
		}

		output, err := json.MarshalIndent(repo, "", "  ")
		if err != nil {
			return fmt.Errorf("Failed to marshal: %w", err)
		}

		if q := cli.String("query"); q != "" {
			fmt.Println(gjson.GetBytes(output, q))
			return nil
		}

		fmt.Println(string(output))
		return nil
	},
}
