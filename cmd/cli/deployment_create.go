package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/api"
)

var deploymentCreateCommand = &cli.Command{
	Name:      "create",
	Usage:     "Deploy a specific ref(branch, SHA, tag) to the environment.",
	ArgsUsage: "<owner>/<repo>",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "type",
			Usage:    "The type of the ref: 'commit', 'branch', or 'tag'.",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "env",
			Usage:    "The name of the environment.",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "ref",
			Usage:    "The specific ref. It can be any named branch, tag, or SHA.",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:    "field",
			Aliases: []string{"f"},
			Usage:   "The pair of key and value to add to the payload. The format must be <key>=<value>.",
		},
	},
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

		// If the 'dynamic_payload' field is enabled,
		// it creates a payload and pass it as a parameter.
		var payload map[string]interface{}
		if env := config.GetEnv(cli.String("env")); env.IsDynamicPayloadEnabled() {
			if payload, err = buildDyanmicPayload(cli.StringSlice("field"), env); err != nil {
				return err
			}
		}

		d, err := c.Deployment.Create(cli.Context, ns, n, &api.DeploymentCreateRequest{
			Type:           cli.String("type"),
			Ref:            cli.String("ref"),
			Env:            cli.String("env"),
			DynamicPayload: payload,
		})
		if err != nil {
			return err
		}

		return printJson(cli, d)
	},
}

func buildDyanmicPayload(fields []string, env *extent.Env) (map[string]interface{}, error) {
	values := make(map[string]string)

	for _, f := range fields {
		keyAndValue := strings.SplitN(f, "=", 2)
		if len(keyAndValue) != 2 {
			return nil, fmt.Errorf("The field must be <key>=<value> format")
		}

		values[keyAndValue[0]] = keyAndValue[1]
	}

	payload := make(map[string]interface{})

	for key, input := range env.DynamicPayload.Inputs {
		val, ok := values[key]
		// Set the default value if the value doesn't exist.
		if !ok {
			if input.Default != nil {
				payload[key] = *input.Default
				continue
			}
		}

		parsed, err := parseValue(input, val)
		if err != nil {
			return nil, fmt.Errorf("The value of the '%s' field is not %s", key, input.Type)
		}

		payload[key] = parsed
	}

	return payload, nil
}

func parseValue(input extent.Input, s string) (interface{}, error) {
	switch input.Type {
	case extent.InputTypeSelect:
		return s, nil

	case extent.InputTypeString:
		return s, nil

	case extent.InputTypeNumber:
		if val, err := strconv.ParseFloat(s, 64); err != nil {
			return nil, err
		} else {
			return val, nil
		}

	case extent.InputTypeBoolean:
		if val, err := strconv.ParseBool(s); err != nil {
			return nil, err
		} else {
			return val, nil
		}

	default:
		return nil, fmt.Errorf("%s is unsupported type.", input.Type)
	}
}
