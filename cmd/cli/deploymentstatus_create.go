package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/gitploy-io/gitploy/pkg/api"
)

var deploymentStatusCreateCommand = &cli.Command{
	Name:      "create",
	Usage:     "Create the remote deployment status under the deployment.",
	ArgsUsage: "<owner>/<repo> <number>",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "status",
			Usage:    "The remote deployment status. For GitHub, Can be one of error, failure, in_progress, queued, or success.",
			Required: true,
		},
		&cli.StringFlag{
			Name:        "description",
			Usage:       "A short description of the status.",
			DefaultText: "USER_LOGIN updated the status manually.",
		},
	},
	Action: func(cli *cli.Context) error {
		ns, n, err := splitFullName(cli.Args().First())
		if err != nil {
			return err
		}

		number, err := strconv.Atoi(cli.Args().Get(1))
		if err != nil {
			return err
		}

		// Build the request.
		c := buildClient(cli)
		req := &api.DeploymentStatusCreateRemoteRequest{
			Status:      cli.String("status"),
			Description: cli.String("description"),
		}
		if cli.String("description") == "" {
			req.Description = buildDescription(cli)
		}

		dss, err := c.DeploymentStatus.CreateRemote(cli.Context, ns, n, number, req)
		if err != nil {
			return err
		}

		return printJson(cli, dss)
	},
}

func buildDescription(cli *cli.Context) string {
	c := buildClient(cli)

	me, err := c.User.GetMe(cli.Context)
	if err != nil {
		return "Updated the status manually."
	}

	return fmt.Sprintf("%s updated the status manually.", me.Login)
}
