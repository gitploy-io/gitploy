package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/vo"
	"github.com/google/go-github/v32/github"
)

func (g *Github) CreateRemoteDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *vo.Env) (*vo.RemoteDeployment, error) {
	gd, res, err := g.Client(ctx, u.Token).
		Repositories.
		CreateDeployment(ctx, r.Namespace, r.Name, &github.DeploymentRequest{
			Ref:                   github.String(d.Ref),
			Environment:           github.String(e.Name),
			Task:                  github.String(e.Task),
			Description:           github.String(e.Description),
			AutoMerge:             github.Bool(e.AutoMerge),
			RequiredContexts:      &e.RequiredContexts,
			Payload:               github.String(e.Payload),
			ProductionEnvironment: github.Bool(e.ProductionEnvironment),
		})
	if res.StatusCode == http.StatusConflict {
		return nil, &vo.UnprocessibleDeploymentError{
			Err: err,
		}
	} else if res.StatusCode == http.StatusUnprocessableEntity {
		return nil, &vo.UnprocessibleDeploymentError{
			Err: err,
		}
	}
	if err != nil {
		return nil, err
	}

	var url string
	commit, _, err := g.Client(ctx, u.Token).
		Repositories.
		GetCommit(ctx, r.Namespace, r.Name, *gd.SHA)
	if err == nil {
		url = *commit.HTMLURL
	}

	return &vo.RemoteDeployment{
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

func (g *Github) GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error) {
	file, _, res, err := g.Client(ctx, u.Token).
		Repositories.
		GetContents(ctx, r.Namespace, r.Name, r.ConfigPath, &github.RepositoryContentGetOptions{})
	if res.StatusCode == http.StatusNotFound {
		return nil, &vo.ConfigNotFoundError{
			RepoName: r.Name,
		}
	} else if err != nil {
		return nil, err
	}

	content, err := base64.StdEncoding.DecodeString(*file.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the config file: %w", err)
	}

	c := &vo.Config{}
	if err := vo.UnmarshalYAML([]byte(content), c); err != nil {
		return nil, &vo.ConfigParseError{
			RepoName: r.Name,
			Err:      err,
		}
	}

	return c, nil
}
