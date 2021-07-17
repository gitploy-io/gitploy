package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/approval"
)

func (s *Store) ListApprovals(ctx context.Context, d *ent.Deployment) ([]*ent.Approval, error) {
	return s.c.Approval.
		Query().
		Where(
			approval.DeploymentIDEQ(d.ID),
		).
		WithUser().
		WithDeployment().
		All(ctx)
}

func (s *Store) FindApprovalByID(ctx context.Context, id int) (*ent.Approval, error) {
	return s.c.Approval.
		Query().
		Where(
			approval.IDEQ(id),
		).
		WithUser().
		WithDeployment().
		First(ctx)
}

func (s *Store) FindApprovalOfUser(ctx context.Context, d *ent.Deployment, u *ent.User) (*ent.Approval, error) {
	return s.c.Approval.
		Query().
		Where(
			approval.And(
				approval.UserIDEQ(u.ID),
				approval.DeploymentIDEQ(d.ID),
			),
		).
		WithUser().
		WithDeployment().
		First(ctx)
}

func (s *Store) CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
	return s.c.Approval.
		Create().
		SetUserID(a.UserID).
		SetDeploymentID(a.DeploymentID).
		Save(ctx)
}

func (s *Store) UpdateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
	return s.c.Approval.
		UpdateOne(a).
		SetIsApproved(a.IsApproved).
		Save(ctx)
}

func (s *Store) DeleteApproval(ctx context.Context, a *ent.Approval) error {
	return s.c.Approval.
		DeleteOne(a).
		Exec(ctx)
}
