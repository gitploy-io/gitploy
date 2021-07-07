package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (s *Store) CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
	return s.c.Approval.
		Create().
		SetUserID(a.UserID).
		SetDeploymentID(a.DeploymentID).
		Save(ctx)
}
