// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/predicate"
	"github.com/hanjunlee/gitploy/ent/repo"
)

// RepoUpdate is the builder for updating Repo entities.
type RepoUpdate struct {
	config
	hooks    []Hook
	mutation *RepoMutation
}

// Where adds a new predicate for the RepoUpdate builder.
func (ru *RepoUpdate) Where(ps ...predicate.Repo) *RepoUpdate {
	ru.mutation.predicates = append(ru.mutation.predicates, ps...)
	return ru
}

// SetNamespace sets the "namespace" field.
func (ru *RepoUpdate) SetNamespace(s string) *RepoUpdate {
	ru.mutation.SetNamespace(s)
	return ru
}

// SetName sets the "name" field.
func (ru *RepoUpdate) SetName(s string) *RepoUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetDescription sets the "description" field.
func (ru *RepoUpdate) SetDescription(s string) *RepoUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RepoUpdate) SetNillableDescription(s *string) *RepoUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// ClearDescription clears the value of the "description" field.
func (ru *RepoUpdate) ClearDescription() *RepoUpdate {
	ru.mutation.ClearDescription()
	return ru
}

// SetConfigPath sets the "config_path" field.
func (ru *RepoUpdate) SetConfigPath(s string) *RepoUpdate {
	ru.mutation.SetConfigPath(s)
	return ru
}

// SetNillableConfigPath sets the "config_path" field if the given value is not nil.
func (ru *RepoUpdate) SetNillableConfigPath(s *string) *RepoUpdate {
	if s != nil {
		ru.SetConfigPath(*s)
	}
	return ru
}

// SetSyncedAt sets the "synced_at" field.
func (ru *RepoUpdate) SetSyncedAt(t time.Time) *RepoUpdate {
	ru.mutation.SetSyncedAt(t)
	return ru
}

// SetNillableSyncedAt sets the "synced_at" field if the given value is not nil.
func (ru *RepoUpdate) SetNillableSyncedAt(t *time.Time) *RepoUpdate {
	if t != nil {
		ru.SetSyncedAt(*t)
	}
	return ru
}

// ClearSyncedAt clears the value of the "synced_at" field.
func (ru *RepoUpdate) ClearSyncedAt() *RepoUpdate {
	ru.mutation.ClearSyncedAt()
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *RepoUpdate) SetCreatedAt(t time.Time) *RepoUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *RepoUpdate) SetNillableCreatedAt(t *time.Time) *RepoUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RepoUpdate) SetUpdatedAt(t time.Time) *RepoUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// AddPermIDs adds the "perms" edge to the Perm entity by IDs.
func (ru *RepoUpdate) AddPermIDs(ids ...int) *RepoUpdate {
	ru.mutation.AddPermIDs(ids...)
	return ru
}

// AddPerms adds the "perms" edges to the Perm entity.
func (ru *RepoUpdate) AddPerms(p ...*Perm) *RepoUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPermIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (ru *RepoUpdate) AddDeploymentIDs(ids ...int) *RepoUpdate {
	ru.mutation.AddDeploymentIDs(ids...)
	return ru
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (ru *RepoUpdate) AddDeployments(d ...*Deployment) *RepoUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.AddDeploymentIDs(ids...)
}

// Mutation returns the RepoMutation object of the builder.
func (ru *RepoUpdate) Mutation() *RepoMutation {
	return ru.mutation
}

// ClearPerms clears all "perms" edges to the Perm entity.
func (ru *RepoUpdate) ClearPerms() *RepoUpdate {
	ru.mutation.ClearPerms()
	return ru
}

// RemovePermIDs removes the "perms" edge to Perm entities by IDs.
func (ru *RepoUpdate) RemovePermIDs(ids ...int) *RepoUpdate {
	ru.mutation.RemovePermIDs(ids...)
	return ru
}

// RemovePerms removes "perms" edges to Perm entities.
func (ru *RepoUpdate) RemovePerms(p ...*Perm) *RepoUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePermIDs(ids...)
}

// ClearDeployments clears all "deployments" edges to the Deployment entity.
func (ru *RepoUpdate) ClearDeployments() *RepoUpdate {
	ru.mutation.ClearDeployments()
	return ru
}

// RemoveDeploymentIDs removes the "deployments" edge to Deployment entities by IDs.
func (ru *RepoUpdate) RemoveDeploymentIDs(ids ...int) *RepoUpdate {
	ru.mutation.RemoveDeploymentIDs(ids...)
	return ru
}

// RemoveDeployments removes "deployments" edges to Deployment entities.
func (ru *RepoUpdate) RemoveDeployments(d ...*Deployment) *RepoUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.RemoveDeploymentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RepoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RepoUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RepoUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RepoUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RepoUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := repo.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

func (ru *RepoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repo.Table,
			Columns: repo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: repo.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Namespace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldNamespace,
		})
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldName,
		})
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldDescription,
		})
	}
	if ru.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: repo.FieldDescription,
		})
	}
	if value, ok := ru.mutation.ConfigPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldConfigPath,
		})
	}
	if value, ok := ru.mutation.SyncedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldSyncedAt,
		})
	}
	if ru.mutation.SyncedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: repo.FieldSyncedAt,
		})
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldCreatedAt,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldUpdatedAt,
		})
	}
	if ru.mutation.PermsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPermsIDs(); len(nodes) > 0 && !ru.mutation.PermsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PermsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.DeploymentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedDeploymentsIDs(); len(nodes) > 0 && !ru.mutation.DeploymentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.DeploymentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RepoUpdateOne is the builder for updating a single Repo entity.
type RepoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepoMutation
}

// SetNamespace sets the "namespace" field.
func (ruo *RepoUpdateOne) SetNamespace(s string) *RepoUpdateOne {
	ruo.mutation.SetNamespace(s)
	return ruo
}

// SetName sets the "name" field.
func (ruo *RepoUpdateOne) SetName(s string) *RepoUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RepoUpdateOne) SetDescription(s string) *RepoUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RepoUpdateOne) SetNillableDescription(s *string) *RepoUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// ClearDescription clears the value of the "description" field.
func (ruo *RepoUpdateOne) ClearDescription() *RepoUpdateOne {
	ruo.mutation.ClearDescription()
	return ruo
}

// SetConfigPath sets the "config_path" field.
func (ruo *RepoUpdateOne) SetConfigPath(s string) *RepoUpdateOne {
	ruo.mutation.SetConfigPath(s)
	return ruo
}

// SetNillableConfigPath sets the "config_path" field if the given value is not nil.
func (ruo *RepoUpdateOne) SetNillableConfigPath(s *string) *RepoUpdateOne {
	if s != nil {
		ruo.SetConfigPath(*s)
	}
	return ruo
}

// SetSyncedAt sets the "synced_at" field.
func (ruo *RepoUpdateOne) SetSyncedAt(t time.Time) *RepoUpdateOne {
	ruo.mutation.SetSyncedAt(t)
	return ruo
}

// SetNillableSyncedAt sets the "synced_at" field if the given value is not nil.
func (ruo *RepoUpdateOne) SetNillableSyncedAt(t *time.Time) *RepoUpdateOne {
	if t != nil {
		ruo.SetSyncedAt(*t)
	}
	return ruo
}

// ClearSyncedAt clears the value of the "synced_at" field.
func (ruo *RepoUpdateOne) ClearSyncedAt() *RepoUpdateOne {
	ruo.mutation.ClearSyncedAt()
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *RepoUpdateOne) SetCreatedAt(t time.Time) *RepoUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *RepoUpdateOne) SetNillableCreatedAt(t *time.Time) *RepoUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RepoUpdateOne) SetUpdatedAt(t time.Time) *RepoUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// AddPermIDs adds the "perms" edge to the Perm entity by IDs.
func (ruo *RepoUpdateOne) AddPermIDs(ids ...int) *RepoUpdateOne {
	ruo.mutation.AddPermIDs(ids...)
	return ruo
}

// AddPerms adds the "perms" edges to the Perm entity.
func (ruo *RepoUpdateOne) AddPerms(p ...*Perm) *RepoUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPermIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (ruo *RepoUpdateOne) AddDeploymentIDs(ids ...int) *RepoUpdateOne {
	ruo.mutation.AddDeploymentIDs(ids...)
	return ruo
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (ruo *RepoUpdateOne) AddDeployments(d ...*Deployment) *RepoUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.AddDeploymentIDs(ids...)
}

// Mutation returns the RepoMutation object of the builder.
func (ruo *RepoUpdateOne) Mutation() *RepoMutation {
	return ruo.mutation
}

// ClearPerms clears all "perms" edges to the Perm entity.
func (ruo *RepoUpdateOne) ClearPerms() *RepoUpdateOne {
	ruo.mutation.ClearPerms()
	return ruo
}

// RemovePermIDs removes the "perms" edge to Perm entities by IDs.
func (ruo *RepoUpdateOne) RemovePermIDs(ids ...int) *RepoUpdateOne {
	ruo.mutation.RemovePermIDs(ids...)
	return ruo
}

// RemovePerms removes "perms" edges to Perm entities.
func (ruo *RepoUpdateOne) RemovePerms(p ...*Perm) *RepoUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePermIDs(ids...)
}

// ClearDeployments clears all "deployments" edges to the Deployment entity.
func (ruo *RepoUpdateOne) ClearDeployments() *RepoUpdateOne {
	ruo.mutation.ClearDeployments()
	return ruo
}

// RemoveDeploymentIDs removes the "deployments" edge to Deployment entities by IDs.
func (ruo *RepoUpdateOne) RemoveDeploymentIDs(ids ...int) *RepoUpdateOne {
	ruo.mutation.RemoveDeploymentIDs(ids...)
	return ruo
}

// RemoveDeployments removes "deployments" edges to Deployment entities.
func (ruo *RepoUpdateOne) RemoveDeployments(d ...*Deployment) *RepoUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.RemoveDeploymentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RepoUpdateOne) Select(field string, fields ...string) *RepoUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Repo entity.
func (ruo *RepoUpdateOne) Save(ctx context.Context) (*Repo, error) {
	var (
		err  error
		node *Repo
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RepoUpdateOne) SaveX(ctx context.Context) *Repo {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RepoUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RepoUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RepoUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := repo.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

func (ruo *RepoUpdateOne) sqlSave(ctx context.Context) (_node *Repo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repo.Table,
			Columns: repo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: repo.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Repo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repo.FieldID)
		for _, f := range fields {
			if !repo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != repo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Namespace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldNamespace,
		})
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldName,
		})
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldDescription,
		})
	}
	if ruo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: repo.FieldDescription,
		})
	}
	if value, ok := ruo.mutation.ConfigPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldConfigPath,
		})
	}
	if value, ok := ruo.mutation.SyncedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldSyncedAt,
		})
	}
	if ruo.mutation.SyncedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: repo.FieldSyncedAt,
		})
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldCreatedAt,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldUpdatedAt,
		})
	}
	if ruo.mutation.PermsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPermsIDs(); len(nodes) > 0 && !ruo.mutation.PermsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PermsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.DeploymentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedDeploymentsIDs(); len(nodes) > 0 && !ruo.mutation.DeploymentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.DeploymentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Repo{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
