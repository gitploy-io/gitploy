// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/ent/perm"
	"github.com/gitploy-io/gitploy/ent/repo"
	"github.com/gitploy-io/gitploy/ent/user"
)

// PermCreate is the builder for creating a Perm entity.
type PermCreate struct {
	config
	mutation *PermMutation
	hooks    []Hook
}

// SetRepoPerm sets the "repo_perm" field.
func (pc *PermCreate) SetRepoPerm(pp perm.RepoPerm) *PermCreate {
	pc.mutation.SetRepoPerm(pp)
	return pc
}

// SetNillableRepoPerm sets the "repo_perm" field if the given value is not nil.
func (pc *PermCreate) SetNillableRepoPerm(pp *perm.RepoPerm) *PermCreate {
	if pp != nil {
		pc.SetRepoPerm(*pp)
	}
	return pc
}

// SetSyncedAt sets the "synced_at" field.
func (pc *PermCreate) SetSyncedAt(t time.Time) *PermCreate {
	pc.mutation.SetSyncedAt(t)
	return pc
}

// SetNillableSyncedAt sets the "synced_at" field if the given value is not nil.
func (pc *PermCreate) SetNillableSyncedAt(t *time.Time) *PermCreate {
	if t != nil {
		pc.SetSyncedAt(*t)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PermCreate) SetCreatedAt(t time.Time) *PermCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PermCreate) SetNillableCreatedAt(t *time.Time) *PermCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PermCreate) SetUpdatedAt(t time.Time) *PermCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PermCreate) SetNillableUpdatedAt(t *time.Time) *PermCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *PermCreate) SetUserID(i int64) *PermCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetRepoID sets the "repo_id" field.
func (pc *PermCreate) SetRepoID(i int64) *PermCreate {
	pc.mutation.SetRepoID(i)
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PermCreate) SetUser(u *User) *PermCreate {
	return pc.SetUserID(u.ID)
}

// SetRepo sets the "repo" edge to the Repo entity.
func (pc *PermCreate) SetRepo(r *Repo) *PermCreate {
	return pc.SetRepoID(r.ID)
}

// Mutation returns the PermMutation object of the builder.
func (pc *PermCreate) Mutation() *PermMutation {
	return pc.mutation
}

// Save creates the Perm in the database.
func (pc *PermCreate) Save(ctx context.Context) (*Perm, error) {
	var (
		err  error
		node *Perm
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermCreate) SaveX(ctx context.Context) *Perm {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PermCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PermCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PermCreate) defaults() {
	if _, ok := pc.mutation.RepoPerm(); !ok {
		v := perm.DefaultRepoPerm
		pc.mutation.SetRepoPerm(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := perm.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := perm.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermCreate) check() error {
	if _, ok := pc.mutation.RepoPerm(); !ok {
		return &ValidationError{Name: "repo_perm", err: errors.New(`ent: missing required field "repo_perm"`)}
	}
	if v, ok := pc.mutation.RepoPerm(); ok {
		if err := perm.RepoPermValidator(v); err != nil {
			return &ValidationError{Name: "repo_perm", err: fmt.Errorf(`ent: validator failed for field "repo_perm": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := pc.mutation.RepoID(); !ok {
		return &ValidationError{Name: "repo_id", err: errors.New(`ent: missing required field "repo_id"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	if _, ok := pc.mutation.RepoID(); !ok {
		return &ValidationError{Name: "repo", err: errors.New("ent: missing required edge \"repo\"")}
	}
	return nil
}

func (pc *PermCreate) sqlSave(ctx context.Context) (*Perm, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PermCreate) createSpec() (*Perm, *sqlgraph.CreateSpec) {
	var (
		_node = &Perm{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: perm.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: perm.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.RepoPerm(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: perm.FieldRepoPerm,
		})
		_node.RepoPerm = value
	}
	if value, ok := pc.mutation.SyncedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldSyncedAt,
		})
		_node.SyncedAt = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   perm.UserTable,
			Columns: []string{perm.UserColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.RepoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   perm.RepoTable,
			Columns: []string{perm.RepoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
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
	return _node, _spec
}

// PermCreateBulk is the builder for creating many Perm entities in bulk.
type PermCreateBulk struct {
	config
	builders []*PermCreate
}

// Save creates the Perm entities in the database.
func (pcb *PermCreateBulk) Save(ctx context.Context) ([]*Perm, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Perm, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermCreateBulk) SaveX(ctx context.Context) []*Perm {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PermCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PermCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
