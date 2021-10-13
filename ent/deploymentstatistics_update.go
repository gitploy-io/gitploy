// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/ent/deploymentstatistics"
	"github.com/gitploy-io/gitploy/ent/predicate"
)

// DeploymentStatisticsUpdate is the builder for updating DeploymentStatistics entities.
type DeploymentStatisticsUpdate struct {
	config
	hooks    []Hook
	mutation *DeploymentStatisticsMutation
}

// Where appends a list predicates to the DeploymentStatisticsUpdate builder.
func (dsu *DeploymentStatisticsUpdate) Where(ps ...predicate.DeploymentStatistics) *DeploymentStatisticsUpdate {
	dsu.mutation.Where(ps...)
	return dsu
}

// SetNamespace sets the "namespace" field.
func (dsu *DeploymentStatisticsUpdate) SetNamespace(s string) *DeploymentStatisticsUpdate {
	dsu.mutation.SetNamespace(s)
	return dsu
}

// SetName sets the "name" field.
func (dsu *DeploymentStatisticsUpdate) SetName(s string) *DeploymentStatisticsUpdate {
	dsu.mutation.SetName(s)
	return dsu
}

// SetEnv sets the "env" field.
func (dsu *DeploymentStatisticsUpdate) SetEnv(s string) *DeploymentStatisticsUpdate {
	dsu.mutation.SetEnv(s)
	return dsu
}

// SetCount sets the "count" field.
func (dsu *DeploymentStatisticsUpdate) SetCount(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.ResetCount()
	dsu.mutation.SetCount(i)
	return dsu
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (dsu *DeploymentStatisticsUpdate) SetNillableCount(i *int) *DeploymentStatisticsUpdate {
	if i != nil {
		dsu.SetCount(*i)
	}
	return dsu
}

// AddCount adds i to the "count" field.
func (dsu *DeploymentStatisticsUpdate) AddCount(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.AddCount(i)
	return dsu
}

// SetCreatedAt sets the "created_at" field.
func (dsu *DeploymentStatisticsUpdate) SetCreatedAt(t time.Time) *DeploymentStatisticsUpdate {
	dsu.mutation.SetCreatedAt(t)
	return dsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dsu *DeploymentStatisticsUpdate) SetNillableCreatedAt(t *time.Time) *DeploymentStatisticsUpdate {
	if t != nil {
		dsu.SetCreatedAt(*t)
	}
	return dsu
}

// SetUpdatedAt sets the "updated_at" field.
func (dsu *DeploymentStatisticsUpdate) SetUpdatedAt(t time.Time) *DeploymentStatisticsUpdate {
	dsu.mutation.SetUpdatedAt(t)
	return dsu
}

// Mutation returns the DeploymentStatisticsMutation object of the builder.
func (dsu *DeploymentStatisticsUpdate) Mutation() *DeploymentStatisticsMutation {
	return dsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dsu *DeploymentStatisticsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	dsu.defaults()
	if len(dsu.hooks) == 0 {
		affected, err = dsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentStatisticsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dsu.mutation = mutation
			affected, err = dsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dsu.hooks) - 1; i >= 0; i-- {
			if dsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dsu *DeploymentStatisticsUpdate) SaveX(ctx context.Context) int {
	affected, err := dsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dsu *DeploymentStatisticsUpdate) Exec(ctx context.Context) error {
	_, err := dsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsu *DeploymentStatisticsUpdate) ExecX(ctx context.Context) {
	if err := dsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dsu *DeploymentStatisticsUpdate) defaults() {
	if _, ok := dsu.mutation.UpdatedAt(); !ok {
		v := deploymentstatistics.UpdateDefaultUpdatedAt()
		dsu.mutation.SetUpdatedAt(v)
	}
}

func (dsu *DeploymentStatisticsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deploymentstatistics.Table,
			Columns: deploymentstatistics.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deploymentstatistics.FieldID,
			},
		},
	}
	if ps := dsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dsu.mutation.Namespace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatistics.FieldNamespace,
		})
	}
	if value, ok := dsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatistics.FieldName,
		})
	}
	if value, ok := dsu.mutation.Env(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatistics.FieldEnv,
		})
	}
	if value, ok := dsu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldCount,
		})
	}
	if value, ok := dsu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldCount,
		})
	}
	if value, ok := dsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deploymentstatistics.FieldCreatedAt,
		})
	}
	if value, ok := dsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deploymentstatistics.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deploymentstatistics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DeploymentStatisticsUpdateOne is the builder for updating a single DeploymentStatistics entity.
type DeploymentStatisticsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeploymentStatisticsMutation
}

// SetNamespace sets the "namespace" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetNamespace(s string) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.SetNamespace(s)
	return dsuo
}

// SetName sets the "name" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetName(s string) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.SetName(s)
	return dsuo
}

// SetEnv sets the "env" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetEnv(s string) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.SetEnv(s)
	return dsuo
}

// SetCount sets the "count" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetCount(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.ResetCount()
	dsuo.mutation.SetCount(i)
	return dsuo
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (dsuo *DeploymentStatisticsUpdateOne) SetNillableCount(i *int) *DeploymentStatisticsUpdateOne {
	if i != nil {
		dsuo.SetCount(*i)
	}
	return dsuo
}

// AddCount adds i to the "count" field.
func (dsuo *DeploymentStatisticsUpdateOne) AddCount(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.AddCount(i)
	return dsuo
}

// SetCreatedAt sets the "created_at" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetCreatedAt(t time.Time) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.SetCreatedAt(t)
	return dsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dsuo *DeploymentStatisticsUpdateOne) SetNillableCreatedAt(t *time.Time) *DeploymentStatisticsUpdateOne {
	if t != nil {
		dsuo.SetCreatedAt(*t)
	}
	return dsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetUpdatedAt(t time.Time) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.SetUpdatedAt(t)
	return dsuo
}

// Mutation returns the DeploymentStatisticsMutation object of the builder.
func (dsuo *DeploymentStatisticsUpdateOne) Mutation() *DeploymentStatisticsMutation {
	return dsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dsuo *DeploymentStatisticsUpdateOne) Select(field string, fields ...string) *DeploymentStatisticsUpdateOne {
	dsuo.fields = append([]string{field}, fields...)
	return dsuo
}

// Save executes the query and returns the updated DeploymentStatistics entity.
func (dsuo *DeploymentStatisticsUpdateOne) Save(ctx context.Context) (*DeploymentStatistics, error) {
	var (
		err  error
		node *DeploymentStatistics
	)
	dsuo.defaults()
	if len(dsuo.hooks) == 0 {
		node, err = dsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentStatisticsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dsuo.mutation = mutation
			node, err = dsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dsuo.hooks) - 1; i >= 0; i-- {
			if dsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dsuo *DeploymentStatisticsUpdateOne) SaveX(ctx context.Context) *DeploymentStatistics {
	node, err := dsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dsuo *DeploymentStatisticsUpdateOne) Exec(ctx context.Context) error {
	_, err := dsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsuo *DeploymentStatisticsUpdateOne) ExecX(ctx context.Context) {
	if err := dsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dsuo *DeploymentStatisticsUpdateOne) defaults() {
	if _, ok := dsuo.mutation.UpdatedAt(); !ok {
		v := deploymentstatistics.UpdateDefaultUpdatedAt()
		dsuo.mutation.SetUpdatedAt(v)
	}
}

func (dsuo *DeploymentStatisticsUpdateOne) sqlSave(ctx context.Context) (_node *DeploymentStatistics, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deploymentstatistics.Table,
			Columns: deploymentstatistics.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deploymentstatistics.FieldID,
			},
		},
	}
	id, ok := dsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DeploymentStatistics.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := dsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deploymentstatistics.FieldID)
		for _, f := range fields {
			if !deploymentstatistics.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deploymentstatistics.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dsuo.mutation.Namespace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatistics.FieldNamespace,
		})
	}
	if value, ok := dsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatistics.FieldName,
		})
	}
	if value, ok := dsuo.mutation.Env(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatistics.FieldEnv,
		})
	}
	if value, ok := dsuo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldCount,
		})
	}
	if value, ok := dsuo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldCount,
		})
	}
	if value, ok := dsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deploymentstatistics.FieldCreatedAt,
		})
	}
	if value, ok := dsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deploymentstatistics.FieldUpdatedAt,
		})
	}
	_node = &DeploymentStatistics{config: dsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deploymentstatistics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}