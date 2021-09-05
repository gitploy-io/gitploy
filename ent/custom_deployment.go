package ent

import (
	"github.com/gitploy-io/gitploy/ent/deployment"
)

// CheckEagerLoading checks the deployment has required edges; Repo and User.
func (d *Deployment) CheckEagerLoading() error {
	if d.Edges.Repo == nil {
		return &EagerLoadingError{
			Edge: "repo",
		}
	}

	if d.Edges.User == nil {
		return &EagerLoadingError{
			Edge: "user",
		}
	}

	return nil
}

func (d *Deployment) GetShortRef() string {
	const maxlen = 7

	if d.Type == deployment.TypeCommit &&
		len(d.Ref) > maxlen {
		return d.Ref[:7]
	}

	return d.Ref
}
