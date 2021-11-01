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
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/event"
	"github.com/gitploy-io/gitploy/ent/predicate"
	"github.com/gitploy-io/gitploy/ent/review"
	"github.com/gitploy-io/gitploy/ent/user"
)

// ReviewUpdate is the builder for updating Review entities.
type ReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ru *ReviewUpdate) Where(ps ...predicate.Review) *ReviewUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetStatus sets the "status" field.
func (ru *ReviewUpdate) SetStatus(r review.Status) *ReviewUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableStatus(r *review.Status) *ReviewUpdate {
	if r != nil {
		ru.SetStatus(*r)
	}
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *ReviewUpdate) SetCreatedAt(t time.Time) *ReviewUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableCreatedAt(t *time.Time) *ReviewUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ReviewUpdate) SetUpdatedAt(t time.Time) *ReviewUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *ReviewUpdate) SetUserID(i int64) *ReviewUpdate {
	ru.mutation.SetUserID(i)
	return ru
}

// SetDeploymentID sets the "deployment_id" field.
func (ru *ReviewUpdate) SetDeploymentID(i int) *ReviewUpdate {
	ru.mutation.SetDeploymentID(i)
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *ReviewUpdate) SetUser(u *User) *ReviewUpdate {
	return ru.SetUserID(u.ID)
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (ru *ReviewUpdate) SetDeployment(d *Deployment) *ReviewUpdate {
	return ru.SetDeploymentID(d.ID)
}

// AddEventIDs adds the "event" edge to the Event entity by IDs.
func (ru *ReviewUpdate) AddEventIDs(ids ...int) *ReviewUpdate {
	ru.mutation.AddEventIDs(ids...)
	return ru
}

// AddEvent adds the "event" edges to the Event entity.
func (ru *ReviewUpdate) AddEvent(e ...*Event) *ReviewUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.AddEventIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ru *ReviewUpdate) Mutation() *ReviewMutation {
	return ru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ru *ReviewUpdate) ClearUser() *ReviewUpdate {
	ru.mutation.ClearUser()
	return ru
}

// ClearDeployment clears the "deployment" edge to the Deployment entity.
func (ru *ReviewUpdate) ClearDeployment() *ReviewUpdate {
	ru.mutation.ClearDeployment()
	return ru
}

// ClearEvent clears all "event" edges to the Event entity.
func (ru *ReviewUpdate) ClearEvent() *ReviewUpdate {
	ru.mutation.ClearEvent()
	return ru
}

// RemoveEventIDs removes the "event" edge to Event entities by IDs.
func (ru *ReviewUpdate) RemoveEventIDs(ids ...int) *ReviewUpdate {
	ru.mutation.RemoveEventIDs(ids...)
	return ru
}

// RemoveEvent removes "event" edges to Event entities.
func (ru *ReviewUpdate) RemoveEvent(e ...*Event) *ReviewUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReviewUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := review.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReviewUpdate) check() error {
	if v, ok := ru.mutation.Status(); ok {
		if err := review.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	if _, ok := ru.mutation.DeploymentID(); ru.mutation.DeploymentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"deployment\"")
	}
	return nil
}

func (ru *ReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
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
	if value, ok := ru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: review.FieldStatus,
		})
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: review.FieldCreatedAt,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: review.FieldUpdatedAt,
		})
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
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
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
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
	if ru.mutation.DeploymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.DeploymentTable,
			Columns: []string{review.DeploymentColumn},
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
	if nodes := ru.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.DeploymentTable,
			Columns: []string{review.DeploymentColumn},
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
	if ru.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.EventTable,
			Columns: []string{review.EventColumn},
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
	if nodes := ru.mutation.RemovedEventIDs(); len(nodes) > 0 && !ru.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.EventTable,
			Columns: []string{review.EventColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.EventTable,
			Columns: []string{review.EventColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReviewUpdateOne is the builder for updating a single Review entity.
type ReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewMutation
}

// SetStatus sets the "status" field.
func (ruo *ReviewUpdateOne) SetStatus(r review.Status) *ReviewUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableStatus(r *review.Status) *ReviewUpdateOne {
	if r != nil {
		ruo.SetStatus(*r)
	}
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *ReviewUpdateOne) SetCreatedAt(t time.Time) *ReviewUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableCreatedAt(t *time.Time) *ReviewUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ReviewUpdateOne) SetUpdatedAt(t time.Time) *ReviewUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *ReviewUpdateOne) SetUserID(i int64) *ReviewUpdateOne {
	ruo.mutation.SetUserID(i)
	return ruo
}

// SetDeploymentID sets the "deployment_id" field.
func (ruo *ReviewUpdateOne) SetDeploymentID(i int) *ReviewUpdateOne {
	ruo.mutation.SetDeploymentID(i)
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *ReviewUpdateOne) SetUser(u *User) *ReviewUpdateOne {
	return ruo.SetUserID(u.ID)
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (ruo *ReviewUpdateOne) SetDeployment(d *Deployment) *ReviewUpdateOne {
	return ruo.SetDeploymentID(d.ID)
}

// AddEventIDs adds the "event" edge to the Event entity by IDs.
func (ruo *ReviewUpdateOne) AddEventIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.AddEventIDs(ids...)
	return ruo
}

// AddEvent adds the "event" edges to the Event entity.
func (ruo *ReviewUpdateOne) AddEvent(e ...*Event) *ReviewUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.AddEventIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ruo *ReviewUpdateOne) Mutation() *ReviewMutation {
	return ruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *ReviewUpdateOne) ClearUser() *ReviewUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// ClearDeployment clears the "deployment" edge to the Deployment entity.
func (ruo *ReviewUpdateOne) ClearDeployment() *ReviewUpdateOne {
	ruo.mutation.ClearDeployment()
	return ruo
}

// ClearEvent clears all "event" edges to the Event entity.
func (ruo *ReviewUpdateOne) ClearEvent() *ReviewUpdateOne {
	ruo.mutation.ClearEvent()
	return ruo
}

// RemoveEventIDs removes the "event" edge to Event entities by IDs.
func (ruo *ReviewUpdateOne) RemoveEventIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.RemoveEventIDs(ids...)
	return ruo
}

// RemoveEvent removes "event" edges to Event entities.
func (ruo *ReviewUpdateOne) RemoveEvent(e ...*Event) *ReviewUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.RemoveEventIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewUpdateOne) Select(field string, fields ...string) *ReviewUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Review entity.
func (ruo *ReviewUpdateOne) Save(ctx context.Context) (*Review, error) {
	var (
		err  error
		node *Review
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewUpdateOne) SaveX(ctx context.Context) *Review {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReviewUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := review.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReviewUpdateOne) check() error {
	if v, ok := ruo.mutation.Status(); ok {
		if err := review.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	if _, ok := ruo.mutation.DeploymentID(); ruo.mutation.DeploymentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"deployment\"")
	}
	return nil
}

func (ruo *ReviewUpdateOne) sqlSave(ctx context.Context) (_node *Review, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Review.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for _, f := range fields {
			if !review.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != review.FieldID {
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
	if value, ok := ruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: review.FieldStatus,
		})
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: review.FieldCreatedAt,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: review.FieldUpdatedAt,
		})
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
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
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
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
	if ruo.mutation.DeploymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.DeploymentTable,
			Columns: []string{review.DeploymentColumn},
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
	if nodes := ruo.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.DeploymentTable,
			Columns: []string{review.DeploymentColumn},
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
	if ruo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.EventTable,
			Columns: []string{review.EventColumn},
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
	if nodes := ruo.mutation.RemovedEventIDs(); len(nodes) > 0 && !ruo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.EventTable,
			Columns: []string{review.EventColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.EventTable,
			Columns: []string{review.EventColumn},
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
	_node = &Review{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
