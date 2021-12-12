package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/google/go-github/v32/github"
)

func (g *Github) CreateRemoteDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *extent.Env) (*extent.RemoteDeployment, error) {
	gd, res, err := g.Client(ctx, u.Token).
		Repositories.
		CreateDeployment(ctx, r.Namespace, r.Name, &github.DeploymentRequest{
			Ref:                   github.String(d.Ref),
			Environment:           github.String(env.Name),
			Task:                  env.Task,
			Description:           env.Description,
			AutoMerge:             env.AutoMerge,
			RequiredContexts:      env.RequiredContexts,
			Payload:               env.Payload,
			ProductionEnvironment: env.ProductionEnvironment,
		})
	if res.StatusCode == http.StatusConflict {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			"There is merge conflict or a commit status check failed.",
			err,
		)
	} else if res.StatusCode == http.StatusUnprocessableEntity {
		return nil, e.NewError(
			e.ErrorCodeDeploymentInvalid,
			err,
		)
	}
	if err != nil {
		return nil, e.NewError(
			e.ErrorCodeInternalError,
			err,
		)
	}

	var url string
	commit, _, err := g.Client(ctx, u.Token).
		Repositories.
		GetCommit(ctx, r.Namespace, r.Name, *gd.SHA)
	if err == nil {
		url = *commit.HTMLURL
	}

	return &extent.RemoteDeployment{
		UID:     *gd.ID,
		SHA:     *gd.SHA,
		HTLMURL: url,
	}, nil
}

func (g *Github) CancelDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, s *ent.DeploymentStatus) error {
	_, _, err := g.Client(ctx, u.Token).
		Repositories.
		CreateDeploymentStatus(ctx, r.Namespace, r.Name, d.UID, &github.DeploymentStatusRequest{
			State:       github.String("inactive"),
			Environment: github.String(d.Env),
			Description: github.String(s.Description),
			LogURL:      github.String(s.LogURL),
		})
	return err
}

func (g *Github) GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*extent.Config, error) {
	file, _, res, err := g.Client(ctx, u.Token).
		Repositories.
		GetContents(ctx, r.Namespace, r.Name, r.ConfigPath, &github.RepositoryContentGetOptions{})
	if res.StatusCode == http.StatusNotFound {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityNotFound,
			"The configuration file is not found.",
			err,
		)
	} else if err != nil {
		return nil, err
	}

	content, err := base64.StdEncoding.DecodeString(*file.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the config file: %w", err)
	}

	c := &extent.Config{}
	if err := extent.UnmarshalYAML([]byte(content), c); err != nil {
		return nil, e.NewError(
			e.ErrorCodeConfigParseError,
			err,
		)
	}

	return c, nil
}
