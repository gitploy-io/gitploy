// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
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

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RepoCreate) SetNillableDescription(s *string) *RepoCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetSyncedAt sets the "synced_at" field.
func (rc *RepoCreate) SetSyncedAt(t time.Time) *RepoCreate {
	rc.mutation.SetSyncedAt(t)
	return rc
}

// SetNillableSyncedAt sets the "synced_at" field if the given value is not nil.
func (rc *RepoCreate) SetNillableSyncedAt(t *time.Time) *RepoCreate {
	if t != nil {
		rc.SetSyncedAt(*t)
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

// SetID sets the "id" field.
func (rc *RepoCreate) SetID(s string) *RepoCreate {
	rc.mutation.SetID(s)
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
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
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

// defaults sets the default values of the builder before save.
func (rc *RepoCreate) defaults() {
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
		return &ValidationError{Name: "namespace", err: errors.New("ent: missing required field \"namespace\"")}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (rc *RepoCreate) sqlSave(ctx context.Context) (*Repo, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (rc *RepoCreate) createSpec() (*Repo, *sqlgraph.CreateSpec) {
	var (
		_node = &Repo{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: repo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
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
	if value, ok := rc.mutation.SyncedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldSyncedAt,
		})
		_node.SyncedAt = value
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
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
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
