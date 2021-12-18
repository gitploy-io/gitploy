// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/model/ent/perm"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
	"github.com/gitploy-io/gitploy/model/ent/repo"
	"github.com/gitploy-io/gitploy/model/ent/user"
)

// PermUpdate is the builder for updating Perm entities.
type PermUpdate struct {
	config
	hooks    []Hook
	mutation *PermMutation
}

// Where appends a list predicates to the PermUpdate builder.
func (pu *PermUpdate) Where(ps ...predicate.Perm) *PermUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetRepoPerm sets the "repo_perm" field.
func (pu *PermUpdate) SetRepoPerm(pp perm.RepoPerm) *PermUpdate {
	pu.mutation.SetRepoPerm(pp)
	return pu
}

// SetNillableRepoPerm sets the "repo_perm" field if the given value is not nil.
func (pu *PermUpdate) SetNillableRepoPerm(pp *perm.RepoPerm) *PermUpdate {
	if pp != nil {
		pu.SetRepoPerm(*pp)
	}
	return pu
}

// SetSyncedAt sets the "synced_at" field.
func (pu *PermUpdate) SetSyncedAt(t time.Time) *PermUpdate {
	pu.mutation.SetSyncedAt(t)
	return pu
}

// SetNillableSyncedAt sets the "synced_at" field if the given value is not nil.
func (pu *PermUpdate) SetNillableSyncedAt(t *time.Time) *PermUpdate {
	if t != nil {
		pu.SetSyncedAt(*t)
	}
	return pu
}

// ClearSyncedAt clears the value of the "synced_at" field.
func (pu *PermUpdate) ClearSyncedAt() *PermUpdate {
	pu.mutation.ClearSyncedAt()
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PermUpdate) SetCreatedAt(t time.Time) *PermUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PermUpdate) SetNillableCreatedAt(t *time.Time) *PermUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermUpdate) SetUpdatedAt(t time.Time) *PermUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PermUpdate) SetUserID(i int64) *PermUpdate {
	pu.mutation.SetUserID(i)
	return pu
}

// SetRepoID sets the "repo_id" field.
func (pu *PermUpdate) SetRepoID(i int64) *PermUpdate {
	pu.mutation.SetRepoID(i)
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PermUpdate) SetUser(u *User) *PermUpdate {
	return pu.SetUserID(u.ID)
}

// SetRepo sets the "repo" edge to the Repo entity.
func (pu *PermUpdate) SetRepo(r *Repo) *PermUpdate {
	return pu.SetRepoID(r.ID)
}

// Mutation returns the PermMutation object of the builder.
func (pu *PermUpdate) Mutation() *PermMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PermUpdate) ClearUser() *PermUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearRepo clears the "repo" edge to the Repo entity.
func (pu *PermUpdate) ClearRepo() *PermUpdate {
	pu.mutation.ClearRepo()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := perm.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermUpdate) check() error {
	if v, ok := pu.mutation.RepoPerm(); ok {
		if err := perm.RepoPermValidator(v); err != nil {
			return &ValidationError{Name: "repo_perm", err: fmt.Errorf("ent: validator failed for field \"repo_perm\": %w", err)}
		}
	}
	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	if _, ok := pu.mutation.RepoID(); pu.mutation.RepoCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"repo\"")
	}
	return nil
}

func (pu *PermUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   perm.Table,
			Columns: perm.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: perm.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.RepoPerm(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: perm.FieldRepoPerm,
		})
	}
	if value, ok := pu.mutation.SyncedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldSyncedAt,
		})
	}
	if pu.mutation.SyncedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: perm.FieldSyncedAt,
		})
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldUpdatedAt,
		})
	}
	if pu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.RepoCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RepoIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{perm.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PermUpdateOne is the builder for updating a single Perm entity.
type PermUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermMutation
}

// SetRepoPerm sets the "repo_perm" field.
func (puo *PermUpdateOne) SetRepoPerm(pp perm.RepoPerm) *PermUpdateOne {
	puo.mutation.SetRepoPerm(pp)
	return puo
}

// SetNillableRepoPerm sets the "repo_perm" field if the given value is not nil.
func (puo *PermUpdateOne) SetNillableRepoPerm(pp *perm.RepoPerm) *PermUpdateOne {
	if pp != nil {
		puo.SetRepoPerm(*pp)
	}
	return puo
}

// SetSyncedAt sets the "synced_at" field.
func (puo *PermUpdateOne) SetSyncedAt(t time.Time) *PermUpdateOne {
	puo.mutation.SetSyncedAt(t)
	return puo
}

// SetNillableSyncedAt sets the "synced_at" field if the given value is not nil.
func (puo *PermUpdateOne) SetNillableSyncedAt(t *time.Time) *PermUpdateOne {
	if t != nil {
		puo.SetSyncedAt(*t)
	}
	return puo
}

// ClearSyncedAt clears the value of the "synced_at" field.
func (puo *PermUpdateOne) ClearSyncedAt() *PermUpdateOne {
	puo.mutation.ClearSyncedAt()
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PermUpdateOne) SetCreatedAt(t time.Time) *PermUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PermUpdateOne) SetNillableCreatedAt(t *time.Time) *PermUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermUpdateOne) SetUpdatedAt(t time.Time) *PermUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *PermUpdateOne) SetUserID(i int64) *PermUpdateOne {
	puo.mutation.SetUserID(i)
	return puo
}

// SetRepoID sets the "repo_id" field.
func (puo *PermUpdateOne) SetRepoID(i int64) *PermUpdateOne {
	puo.mutation.SetRepoID(i)
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PermUpdateOne) SetUser(u *User) *PermUpdateOne {
	return puo.SetUserID(u.ID)
}

// SetRepo sets the "repo" edge to the Repo entity.
func (puo *PermUpdateOne) SetRepo(r *Repo) *PermUpdateOne {
	return puo.SetRepoID(r.ID)
}

// Mutation returns the PermMutation object of the builder.
func (puo *PermUpdateOne) Mutation() *PermMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PermUpdateOne) ClearUser() *PermUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearRepo clears the "repo" edge to the Repo entity.
func (puo *PermUpdateOne) ClearRepo() *PermUpdateOne {
	puo.mutation.ClearRepo()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermUpdateOne) Select(field string, fields ...string) *PermUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Perm entity.
func (puo *PermUpdateOne) Save(ctx context.Context) (*Perm, error) {
	var (
		err  error
		node *Perm
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermUpdateOne) SaveX(ctx context.Context) *Perm {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := perm.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermUpdateOne) check() error {
	if v, ok := puo.mutation.RepoPerm(); ok {
		if err := perm.RepoPermValidator(v); err != nil {
			return &ValidationError{Name: "repo_perm", err: fmt.Errorf("ent: validator failed for field \"repo_perm\": %w", err)}
		}
	}
	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	if _, ok := puo.mutation.RepoID(); puo.mutation.RepoCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"repo\"")
	}
	return nil
}

func (puo *PermUpdateOne) sqlSave(ctx context.Context) (_node *Perm, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   perm.Table,
			Columns: perm.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: perm.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Perm.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, perm.FieldID)
		for _, f := range fields {
			if !perm.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != perm.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.RepoPerm(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: perm.FieldRepoPerm,
		})
	}
	if value, ok := puo.mutation.SyncedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldSyncedAt,
		})
	}
	if puo.mutation.SyncedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: perm.FieldSyncedAt,
		})
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: perm.FieldUpdatedAt,
		})
	}
	if puo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.RepoCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RepoIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Perm{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{perm.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}