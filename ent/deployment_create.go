// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/deploymentstatus"
	"github.com/hanjunlee/gitploy/ent/event"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/user"
)

// DeploymentCreate is the builder for creating a Deployment entity.
type DeploymentCreate struct {
	config
	mutation *DeploymentMutation
	hooks    []Hook
}

// SetNumber sets the "number" field.
func (dc *DeploymentCreate) SetNumber(i int) *DeploymentCreate {
	dc.mutation.SetNumber(i)
	return dc
}

// SetType sets the "type" field.
func (dc *DeploymentCreate) SetType(d deployment.Type) *DeploymentCreate {
	dc.mutation.SetType(d)
	return dc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableType(d *deployment.Type) *DeploymentCreate {
	if d != nil {
		dc.SetType(*d)
	}
	return dc
}

// SetEnv sets the "env" field.
func (dc *DeploymentCreate) SetEnv(s string) *DeploymentCreate {
	dc.mutation.SetEnv(s)
	return dc
}

// SetRef sets the "ref" field.
func (dc *DeploymentCreate) SetRef(s string) *DeploymentCreate {
	dc.mutation.SetRef(s)
	return dc
}

// SetStatus sets the "status" field.
func (dc *DeploymentCreate) SetStatus(d deployment.Status) *DeploymentCreate {
	dc.mutation.SetStatus(d)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableStatus(d *deployment.Status) *DeploymentCreate {
	if d != nil {
		dc.SetStatus(*d)
	}
	return dc
}

// SetUID sets the "uid" field.
func (dc *DeploymentCreate) SetUID(i int64) *DeploymentCreate {
	dc.mutation.SetUID(i)
	return dc
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableUID(i *int64) *DeploymentCreate {
	if i != nil {
		dc.SetUID(*i)
	}
	return dc
}

// SetSha sets the "sha" field.
func (dc *DeploymentCreate) SetSha(s string) *DeploymentCreate {
	dc.mutation.SetSha(s)
	return dc
}

// SetNillableSha sets the "sha" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableSha(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetSha(*s)
	}
	return dc
}

// SetHTMLURL sets the "html_url" field.
func (dc *DeploymentCreate) SetHTMLURL(s string) *DeploymentCreate {
	dc.mutation.SetHTMLURL(s)
	return dc
}

// SetNillableHTMLURL sets the "html_url" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableHTMLURL(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetHTMLURL(*s)
	}
	return dc
}

// SetIsRollback sets the "is_rollback" field.
func (dc *DeploymentCreate) SetIsRollback(b bool) *DeploymentCreate {
	dc.mutation.SetIsRollback(b)
	return dc
}

// SetNillableIsRollback sets the "is_rollback" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableIsRollback(b *bool) *DeploymentCreate {
	if b != nil {
		dc.SetIsRollback(*b)
	}
	return dc
}

// SetIsApprovalEnabled sets the "is_approval_enabled" field.
func (dc *DeploymentCreate) SetIsApprovalEnabled(b bool) *DeploymentCreate {
	dc.mutation.SetIsApprovalEnabled(b)
	return dc
}

// SetNillableIsApprovalEnabled sets the "is_approval_enabled" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableIsApprovalEnabled(b *bool) *DeploymentCreate {
	if b != nil {
		dc.SetIsApprovalEnabled(*b)
	}
	return dc
}

// SetRequiredApprovalCount sets the "required_approval_count" field.
func (dc *DeploymentCreate) SetRequiredApprovalCount(i int) *DeploymentCreate {
	dc.mutation.SetRequiredApprovalCount(i)
	return dc
}

// SetNillableRequiredApprovalCount sets the "required_approval_count" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableRequiredApprovalCount(i *int) *DeploymentCreate {
	if i != nil {
		dc.SetRequiredApprovalCount(*i)
	}
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeploymentCreate) SetCreatedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableCreatedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DeploymentCreate) SetUpdatedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableUpdatedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetUserID sets the "user_id" field.
func (dc *DeploymentCreate) SetUserID(s string) *DeploymentCreate {
	dc.mutation.SetUserID(s)
	return dc
}

// SetRepoID sets the "repo_id" field.
func (dc *DeploymentCreate) SetRepoID(s string) *DeploymentCreate {
	dc.mutation.SetRepoID(s)
	return dc
}

// SetUser sets the "user" edge to the User entity.
func (dc *DeploymentCreate) SetUser(u *User) *DeploymentCreate {
	return dc.SetUserID(u.ID)
}

// SetRepo sets the "repo" edge to the Repo entity.
func (dc *DeploymentCreate) SetRepo(r *Repo) *DeploymentCreate {
	return dc.SetRepoID(r.ID)
}

// AddApprovalIDs adds the "approvals" edge to the Approval entity by IDs.
func (dc *DeploymentCreate) AddApprovalIDs(ids ...int) *DeploymentCreate {
	dc.mutation.AddApprovalIDs(ids...)
	return dc
}

// AddApprovals adds the "approvals" edges to the Approval entity.
func (dc *DeploymentCreate) AddApprovals(a ...*Approval) *DeploymentCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return dc.AddApprovalIDs(ids...)
}

// AddDeploymentStatusIDs adds the "deployment_statuses" edge to the DeploymentStatus entity by IDs.
func (dc *DeploymentCreate) AddDeploymentStatusIDs(ids ...int) *DeploymentCreate {
	dc.mutation.AddDeploymentStatusIDs(ids...)
	return dc
}

// AddDeploymentStatuses adds the "deployment_statuses" edges to the DeploymentStatus entity.
func (dc *DeploymentCreate) AddDeploymentStatuses(d ...*DeploymentStatus) *DeploymentCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDeploymentStatusIDs(ids...)
}

// AddEventIDs adds the "event" edge to the Event entity by IDs.
func (dc *DeploymentCreate) AddEventIDs(ids ...int) *DeploymentCreate {
	dc.mutation.AddEventIDs(ids...)
	return dc
}

// AddEvent adds the "event" edges to the Event entity.
func (dc *DeploymentCreate) AddEvent(e ...*Event) *DeploymentCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dc.AddEventIDs(ids...)
}

// Mutation returns the DeploymentMutation object of the builder.
func (dc *DeploymentCreate) Mutation() *DeploymentMutation {
	return dc.mutation
}

// Save creates the Deployment in the database.
func (dc *DeploymentCreate) Save(ctx context.Context) (*Deployment, error) {
	var (
		err  error
		node *Deployment
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeploymentCreate) SaveX(ctx context.Context) *Deployment {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeploymentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeploymentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeploymentCreate) defaults() {
	if _, ok := dc.mutation.GetType(); !ok {
		v := deployment.DefaultType
		dc.mutation.SetType(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := deployment.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.IsRollback(); !ok {
		v := deployment.DefaultIsRollback
		dc.mutation.SetIsRollback(v)
	}
	if _, ok := dc.mutation.IsApprovalEnabled(); !ok {
		v := deployment.DefaultIsApprovalEnabled
		dc.mutation.SetIsApprovalEnabled(v)
	}
	if _, ok := dc.mutation.RequiredApprovalCount(); !ok {
		v := deployment.DefaultRequiredApprovalCount
		dc.mutation.SetRequiredApprovalCount(v)
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := deployment.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := deployment.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeploymentCreate) check() error {
	if _, ok := dc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "number"`)}
	}
	if _, ok := dc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := dc.mutation.GetType(); ok {
		if err := deployment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Env(); !ok {
		return &ValidationError{Name: "env", err: errors.New(`ent: missing required field "env"`)}
	}
	if _, ok := dc.mutation.Ref(); !ok {
		return &ValidationError{Name: "ref", err: errors.New(`ent: missing required field "ref"`)}
	}
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := dc.mutation.Status(); ok {
		if err := deployment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if v, ok := dc.mutation.HTMLURL(); ok {
		if err := deployment.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "html_url": %w`, err)}
		}
	}
	if _, ok := dc.mutation.IsRollback(); !ok {
		return &ValidationError{Name: "is_rollback", err: errors.New(`ent: missing required field "is_rollback"`)}
	}
	if _, ok := dc.mutation.IsApprovalEnabled(); !ok {
		return &ValidationError{Name: "is_approval_enabled", err: errors.New(`ent: missing required field "is_approval_enabled"`)}
	}
	if _, ok := dc.mutation.RequiredApprovalCount(); !ok {
		return &ValidationError{Name: "required_approval_count", err: errors.New(`ent: missing required field "required_approval_count"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := dc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := dc.mutation.RepoID(); !ok {
		return &ValidationError{Name: "repo_id", err: errors.New(`ent: missing required field "repo_id"`)}
	}
	if _, ok := dc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	if _, ok := dc.mutation.RepoID(); !ok {
		return &ValidationError{Name: "repo", err: errors.New("ent: missing required edge \"repo\"")}
	}
	return nil
}

func (dc *DeploymentCreate) sqlSave(ctx context.Context) (*Deployment, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DeploymentCreate) createSpec() (*Deployment, *sqlgraph.CreateSpec) {
	var (
		_node = &Deployment{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deployment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deployment.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Number(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deployment.FieldNumber,
		})
		_node.Number = value
	}
	if value, ok := dc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: deployment.FieldType,
		})
		_node.Type = value
	}
	if value, ok := dc.mutation.Env(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldEnv,
		})
		_node.Env = value
	}
	if value, ok := dc.mutation.Ref(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldRef,
		})
		_node.Ref = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: deployment.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := dc.mutation.UID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: deployment.FieldUID,
		})
		_node.UID = value
	}
	if value, ok := dc.mutation.Sha(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldSha,
		})
		_node.Sha = value
	}
	if value, ok := dc.mutation.HTMLURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldHTMLURL,
		})
		_node.HTMLURL = value
	}
	if value, ok := dc.mutation.IsRollback(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: deployment.FieldIsRollback,
		})
		_node.IsRollback = value
	}
	if value, ok := dc.mutation.IsApprovalEnabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: deployment.FieldIsApprovalEnabled,
		})
		_node.IsApprovalEnabled = value
	}
	if value, ok := dc.mutation.RequiredApprovalCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deployment.FieldRequiredApprovalCount,
		})
		_node.RequiredApprovalCount = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := dc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deployment.UserTable,
			Columns: []string{deployment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.RepoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deployment.RepoTable,
			Columns: []string{deployment.RepoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RepoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ApprovalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deployment.ApprovalsTable,
			Columns: []string{deployment.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: approval.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DeploymentStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deployment.DeploymentStatusesTable,
			Columns: []string{deployment.DeploymentStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deploymentstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deployment.EventTable,
			Columns: []string{deployment.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeploymentCreateBulk is the builder for creating many Deployment entities in bulk.
type DeploymentCreateBulk struct {
	config
	builders []*DeploymentCreate
}

// Save creates the Deployment entities in the database.
func (dcb *DeploymentCreateBulk) Save(ctx context.Context) ([]*Deployment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Deployment, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeploymentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) SaveX(ctx context.Context) []*Deployment {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeploymentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
