package store

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/comment"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *Store) ListCommentsOfDeployment(ctx context.Context, d *ent.Deployment) ([]*ent.Comment, error) {
	cmts, err := s.c.Comment.
		Query().
		Where(
			comment.DeploymentIDEQ(d.ID),
		).
		WithDeployment().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return cmts, nil
}

func (s *Store) FindCommentByID(ctx context.Context, id int) (*ent.Comment, error) {
	cmt, err := s.c.Comment.
		Query().
		Where(
			comment.IDEQ(id),
		).
		WithDeployment().
		WithUser().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeNotFound, "The comment is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return cmt, nil
}

func (s *Store) CreateComment(ctx context.Context, cmt *ent.Comment) (*ent.Comment, error) {
	cmt, err := s.c.Comment.
		Create().
		SetStatus(cmt.Status).
		SetComment(cmt.Comment).
		SetUserID(cmt.UserID).
		SetDeploymentID(cmt.DeploymentID).
		Save(ctx)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return cmt, nil
}
