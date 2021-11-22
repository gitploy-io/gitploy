// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/ent/callback"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/deploymentstatistics"
	"github.com/gitploy-io/gitploy/ent/lock"
	"github.com/gitploy-io/gitploy/ent/perm"
	"github.com/gitploy-io/gitploy/ent/repo"
	"github.com/gitploy-io/gitploy/ent/user"
)

// RepoCreate is the builder for creating a Repo entity.
type RepoCreate struct {
	config
	mutation *RepoMutation
	hooks    []Hook
}

// SetNamespace sets the "namespace" field.
func (rc *RepoCreate) SetNamespace(s string) *RepoCreate {
	rc.mutation.SetNamespace(s)
	return rc
}

// SetName sets the "name" field.
func (rc *RepoCreate) SetName(s string) *RepoCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetDescription sets the "description" field.
func (rc *RepoCreate) SetDescription(s string) *RepoCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetConfigPath sets the "config_path" field.
func (rc *RepoCreate) SetConfigPath(s string) *RepoCreate {
	rc.mutation.SetConfigPath(s)
	return rc
}

// SetNillableConfigPath sets the "config_path" field if the given value is not nil.
func (rc *RepoCreate) SetNillableConfigPath(s *string) *RepoCreate {
	if s != nil {
		rc.SetConfigPath(*s)
	}
	return rc
}

// SetActive sets the "active" field.
func (rc *RepoCreate) SetActive(b bool) *RepoCreate {
	rc.mutation.SetActive(b)
	return rc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (rc *RepoCreate) SetNillableActive(b *bool) *RepoCreate {
	if b != nil {
		rc.SetActive(*b)
	}
	return rc
}

// SetWebhookID sets the "webhook_id" field.
func (rc *RepoCreate) SetWebhookID(i int64) *RepoCreate {
	rc.mutation.SetWebhookID(i)
	return rc
}

// SetNillableWebhookID sets the "webhook_id" field if the given value is not nil.
func (rc *RepoCreate) SetNillableWebhookID(i *int64) *RepoCreate {
	if i != nil {
		rc.SetWebhookID(*i)
	}
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RepoCreate) SetCreatedAt(t time.Time) *RepoCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RepoCreate) SetNillableCreatedAt(t *time.Time) *RepoCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RepoCreate) SetUpdatedAt(t time.Time) *RepoCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RepoCreate) SetNillableUpdatedAt(t *time.Time) *RepoCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetLatestDeployedAt sets the "latest_deployed_at" field.
func (rc *RepoCreate) SetLatestDeployedAt(t time.Time) *RepoCreate {
	rc.mutation.SetLatestDeployedAt(t)
	return rc
}

// SetNillableLatestDeployedAt sets the "latest_deployed_at" field if the given value is not nil.
func (rc *RepoCreate) SetNillableLatestDeployedAt(t *time.Time) *RepoCreate {
	if t != nil {
		rc.SetLatestDeployedAt(*t)
	}
	return rc
}

// SetOwnerID sets the "owner_id" field.
func (rc *RepoCreate) SetOwnerID(i int64) *RepoCreate {
	rc.mutation.SetOwnerID(i)
	return rc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (rc *RepoCreate) SetNillableOwnerID(i *int64) *RepoCreate {
	if i != nil {
		rc.SetOwnerID(*i)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RepoCreate) SetID(i int64) *RepoCreate {
	rc.mutation.SetID(i)
	return rc
}

// AddPermIDs adds the "perms" edge to the Perm entity by IDs.
func (rc *RepoCreate) AddPermIDs(ids ...int) *RepoCreate {
	rc.mutation.AddPermIDs(ids...)
	return rc
}

// AddPerms adds the "perms" edges to the Perm entity.
func (rc *RepoCreate) AddPerms(p ...*Perm) *RepoCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (rc *RepoCreate) AddDeploymentIDs(ids ...int) *RepoCreate {
	rc.mutation.AddDeploymentIDs(ids...)
	return rc
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (rc *RepoCreate) AddDeployments(d ...*Deployment) *RepoCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return rc.AddDeploymentIDs(ids...)
}

// AddCallbackIDs adds the "callback" edge to the Callback entity by IDs.
func (rc *RepoCreate) AddCallbackIDs(ids ...int) *RepoCreate {
	rc.mutation.AddCallbackIDs(ids...)
	return rc
}

// AddCallback adds the "callback" edges to the Callback entity.
func (rc *RepoCreate) AddCallback(c ...*Callback) *RepoCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rc.AddCallbackIDs(ids...)
}

// AddLockIDs adds the "locks" edge to the Lock entity by IDs.
func (rc *RepoCreate) AddLockIDs(ids ...int) *RepoCreate {
	rc.mutation.AddLockIDs(ids...)
	return rc
}

// AddLocks adds the "locks" edges to the Lock entity.
func (rc *RepoCreate) AddLocks(l ...*Lock) *RepoCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return rc.AddLockIDs(ids...)
}

// AddDeploymentStatisticIDs adds the "deployment_statistics" edge to the DeploymentStatistics entity by IDs.
func (rc *RepoCreate) AddDeploymentStatisticIDs(ids ...int) *RepoCreate {
	rc.mutation.AddDeploymentStatisticIDs(ids...)
	return rc
}

// AddDeploymentStatistics adds the "deployment_statistics" edges to the DeploymentStatistics entity.
func (rc *RepoCreate) AddDeploymentStatistics(d ...*DeploymentStatistics) *RepoCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return rc.AddDeploymentStatisticIDs(ids...)
}

// SetOwner sets the "owner" edge to the User entity.
func (rc *RepoCreate) SetOwner(u *User) *RepoCreate {
	return rc.SetOwnerID(u.ID)
}

// Mutation returns the RepoMutation object of the builder.
func (rc *RepoCreate) Mutation() *RepoMutation {
	return rc.mutation
}

// Save creates the Repo in the database.
func (rc *RepoCreate) Save(ctx context.Context) (*Repo, error) {
	var (
		err  error
		node *Repo
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RepoCreate) SaveX(ctx context.Context) *Repo {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RepoCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RepoCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RepoCreate) defaults() {
	if _, ok := rc.mutation.ConfigPath(); !ok {
		v := repo.DefaultConfigPath
		rc.mutation.SetConfigPath(v)
	}
	if _, ok := rc.mutation.Active(); !ok {
		v := repo.DefaultActive
		rc.mutation.SetActive(v)
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := repo.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := repo.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RepoCreate) check() error {
	if _, ok := rc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required field "namespace"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := rc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := rc.mutation.ConfigPath(); !ok {
		return &ValidationError{Name: "config_path", err: errors.New(`ent: missing required field "config_path"`)}
	}
	if _, ok := rc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "active"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (rc *RepoCreate) sqlSave(ctx context.Context) (*Repo, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (rc *RepoCreate) createSpec() (*Repo, *sqlgraph.CreateSpec) {
	var (
		_node = &Repo{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: repo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: repo.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Namespace(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldNamespace,
		})
		_node.Namespace = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := rc.mutation.ConfigPath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldConfigPath,
		})
		_node.ConfigPath = value
	}
	if value, ok := rc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: repo.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := rc.mutation.WebhookID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: repo.FieldWebhookID,
		})
		_node.WebhookID = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.LatestDeployedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldLatestDeployedAt,
		})
		_node.LatestDeployedAt = value
	}
	if nodes := rc.mutation.PermsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repo.PermsTable,
			Columns: []string{repo.PermsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: perm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.DeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repo.DeploymentsTable,
			Columns: []string{repo.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.CallbackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repo.CallbackTable,
			Columns: []string{repo.CallbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: callback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.LocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repo.LocksTable,
			Columns: []string{repo.LocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lock.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.DeploymentStatisticsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repo.DeploymentStatisticsTable,
			Columns: []string{repo.DeploymentStatisticsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deploymentstatistics.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repo.OwnerTable,
			Columns: []string{repo.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RepoCreateBulk is the builder for creating many Repo entities in bulk.
type RepoCreateBulk struct {
	config
	builders []*RepoCreate
}

// Save creates the Repo entities in the database.
func (rcb *RepoCreateBulk) Save(ctx context.Context) ([]*Repo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Repo, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RepoMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RepoCreateBulk) SaveX(ctx context.Context) []*Repo {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RepoCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RepoCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
