package github

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/google/go-github/v32/github"
)

func (g *Github) CreateRemoteDeploymentStatus(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, ds *extent.RemoteDeploymentStatus) (*extent.RemoteDeploymentStatus, error) {
	gds, _, err := g.Client(ctx, u.Token).
		Repositories.
		CreateDeploymentStatus(ctx, r.Namespace, r.Name, d.UID, &github.DeploymentStatusRequest{
			State:       github.String(ds.Status),
			Description: github.String(ds.Description),
			LogURL:      github.String(ds.LogURL),
		})
	if err != nil {
		return nil, e.NewError(e.ErrorCodeEntityUnprocessable, err)
	}

	return mapGithubDeploymentStatusToRemoteDeploymentStatus(gds), nil
}

func mapGithubDeploymentStatusToRemoteDeploymentStatus(gds *github.DeploymentStatus) *extent.RemoteDeploymentStatus {
	return &extent.RemoteDeploymentStatus{
		ID:          *gds.ID,
		Status:      *gds.State,
		Description: *gds.Description,
		LogURL:      *gds.LogURL,
	}
}
