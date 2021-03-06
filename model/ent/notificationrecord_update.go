// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/model/ent/notificationrecord"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
)

// NotificationRecordUpdate is the builder for updating NotificationRecord entities.
type NotificationRecordUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationRecordMutation
}

// Where appends a list predicates to the NotificationRecordUpdate builder.
func (nru *NotificationRecordUpdate) Where(ps ...predicate.NotificationRecord) *NotificationRecordUpdate {
	nru.mutation.Where(ps...)
	return nru
}

// SetEventID sets the "event_id" field.
func (nru *NotificationRecordUpdate) SetEventID(i int) *NotificationRecordUpdate {
	nru.mutation.SetEventID(i)
	return nru
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (nru *NotificationRecordUpdate) SetNillableEventID(i *int) *NotificationRecordUpdate {
	if i != nil {
		nru.SetEventID(*i)
	}
	return nru
}

// ClearEventID clears the value of the "event_id" field.
func (nru *NotificationRecordUpdate) ClearEventID() *NotificationRecordUpdate {
	nru.mutation.ClearEventID()
	return nru
}

// SetEvent sets the "event" edge to the Event entity.
func (nru *NotificationRecordUpdate) SetEvent(e *Event) *NotificationRecordUpdate {
	return nru.SetEventID(e.ID)
}

// Mutation returns the NotificationRecordMutation object of the builder.
func (nru *NotificationRecordUpdate) Mutation() *NotificationRecordMutation {
	return nru.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (nru *NotificationRecordUpdate) ClearEvent() *NotificationRecordUpdate {
	nru.mutation.ClearEvent()
	return nru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nru *NotificationRecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nru.hooks) == 0 {
		affected, err = nru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nru.mutation = mutation
			affected, err = nru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nru.hooks) - 1; i >= 0; i-- {
			if nru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nru *NotificationRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := nru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nru *NotificationRecordUpdate) Exec(ctx context.Context) error {
	_, err := nru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nru *NotificationRecordUpdate) ExecX(ctx context.Context) {
	if err := nru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nru *NotificationRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notificationrecord.Table,
			Columns: notificationrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notificationrecord.FieldID,
			},
		},
	}
	if ps := nru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nru.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationrecord.EventTable,
			Columns: []string{notificationrecord.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nru.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationrecord.EventTable,
			Columns: []string{notificationrecord.EventColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NotificationRecordUpdateOne is the builder for updating a single NotificationRecord entity.
type NotificationRecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationRecordMutation
}

// SetEventID sets the "event_id" field.
func (nruo *NotificationRecordUpdateOne) SetEventID(i int) *NotificationRecordUpdateOne {
	nruo.mutation.SetEventID(i)
	return nruo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (nruo *NotificationRecordUpdateOne) SetNillableEventID(i *int) *NotificationRecordUpdateOne {
	if i != nil {
		nruo.SetEventID(*i)
	}
	return nruo
}

// ClearEventID clears the value of the "event_id" field.
func (nruo *NotificationRecordUpdateOne) ClearEventID() *NotificationRecordUpdateOne {
	nruo.mutation.ClearEventID()
	return nruo
}

// SetEvent sets the "event" edge to the Event entity.
func (nruo *NotificationRecordUpdateOne) SetEvent(e *Event) *NotificationRecordUpdateOne {
	return nruo.SetEventID(e.ID)
}

// Mutation returns the NotificationRecordMutation object of the builder.
func (nruo *NotificationRecordUpdateOne) Mutation() *NotificationRecordMutation {
	return nruo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (nruo *NotificationRecordUpdateOne) ClearEvent() *NotificationRecordUpdateOne {
	nruo.mutation.ClearEvent()
	return nruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nruo *NotificationRecordUpdateOne) Select(field string, fields ...string) *NotificationRecordUpdateOne {
	nruo.fields = append([]string{field}, fields...)
	return nruo
}

// Save executes the query and returns the updated NotificationRecord entity.
func (nruo *NotificationRecordUpdateOne) Save(ctx context.Context) (*NotificationRecord, error) {
	var (
		err  error
		node *NotificationRecord
	)
	if len(nruo.hooks) == 0 {
		node, err = nruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nruo.mutation = mutation
			node, err = nruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nruo.hooks) - 1; i >= 0; i-- {
			if nruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nruo *NotificationRecordUpdateOne) SaveX(ctx context.Context) *NotificationRecord {
	node, err := nruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nruo *NotificationRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := nruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nruo *NotificationRecordUpdateOne) ExecX(ctx context.Context) {
	if err := nruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nruo *NotificationRecordUpdateOne) sqlSave(ctx context.Context) (_node *NotificationRecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notificationrecord.Table,
			Columns: notificationrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notificationrecord.FieldID,
			},
		},
	}
	id, ok := nruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotificationRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationrecord.FieldID)
		for _, f := range fields {
			if !notificationrecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notificationrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nruo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationrecord.EventTable,
			Columns: []string{notificationrecord.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nruo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationrecord.EventTable,
			Columns: []string{notificationrecord.EventColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NotificationRecord{config: nruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
