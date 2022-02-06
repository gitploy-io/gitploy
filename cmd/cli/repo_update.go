package main

import (
	"encoding/json"
	"fmt"
	"strconv"

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
		&cli.StringFlag{
			Name:    "active",
			Aliases: []string{"A"},
			Usage:   "Activate or deactivate the repository. Ex 'true', 'false'",
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

		if active := cli.String("active"); active != "" {
			b, err := strconv.ParseBool(active)
			if err != nil {
				return fmt.Errorf("'%s' is invalid format: %w", active, err)
			}

			req.Active = pointer.ToBool(b)
		}

		c := buildClient(cli)
		repo, err := c.Repos.Update(cli.Context, ns, n, req)
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
