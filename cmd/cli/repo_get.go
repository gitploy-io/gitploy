package main

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
)

var repoGetCommand = &cli.Command{
	Name:      "get",
	Usage:     "Show the repository.",
	ArgsUsage: "<owner>/<repo>",
	Action: func(cli *cli.Context) error {
		ns, n, err := splitFullName(cli.Args().First())
		if err != nil {
			return err
		}

		c := buildClient(cli)
		repo, err := c.Repo.Get(cli.Context, ns, n)
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
