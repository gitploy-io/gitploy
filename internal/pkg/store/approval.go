package store

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/predicate"
)

func (s *Store) SearchApprovals(ctx context.Context, u *ent.User, ss []approval.Status, from time.Time, to time.Time, page, perPage int) ([]*ent.Approval, error) {
	statusIn := func(ss []approval.Status) predicate.Approval {
		if len(ss) == 0 {
			// if not status were provided,
			// it always make this predicate truly.
			return func(s *sql.Selector) {}
		}

		return approval.StatusIn(ss...)
	}

	return s.c.Approval.
		Query().
		Where(
			approval.And(
				approval.UserID(u.ID),
				statusIn(ss),
				approval.CreatedAtGTE(from),
				approval.CreatedAtLT(to),
			),
		).
		WithUser().
		WithDeployment(func(dq *ent.DeploymentQuery) {
			dq.
				WithRepo().
				WithUser()
		}).
		Offset(offset(page, perPage)).
		Limit(perPage).
		All(ctx)
}

func (s *Store) ListApprovals(ctx context.Context, d *ent.Deployment) ([]*ent.Approval, error) {
	return s.c.Approval.
		Query().
		Where(
			approval.DeploymentIDEQ(d.ID),
		).
		WithUser().
		WithDeployment(func(dq *ent.DeploymentQuery) {
			dq.
				WithRepo().
				WithUser()
		}).
		All(ctx)
}

func (s *Store) FindApprovalByID(ctx context.Context, id int) (*ent.Approval, error) {
	return s.c.Approval.
		Query().
		Where(
			approval.IDEQ(id),
		).
		WithUser().
		WithDeployment(func(dq *ent.DeploymentQuery) {
			dq.
				WithRepo().
				WithUser()
		}).
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
		WithDeployment(func(dq *ent.DeploymentQuery) {
			dq.
				WithRepo().
				WithUser()
		}).
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
		SetStatus(a.Status).
		Save(ctx)
}

func (s *Store) DeleteApproval(ctx context.Context, a *ent.Approval) error {
	return s.c.Approval.
		DeleteOne(a).
		Exec(ctx)
}
