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
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/deploymentstatus"
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
	"github.com/gitploy-io/gitploy/model/ent/repo"
)

// DeploymentStatusUpdate is the builder for updating DeploymentStatus entities.
type DeploymentStatusUpdate struct {
	config
	hooks    []Hook
	mutation *DeploymentStatusMutation
}

// Where appends a list predicates to the DeploymentStatusUpdate builder.
func (dsu *DeploymentStatusUpdate) Where(ps ...predicate.DeploymentStatus) *DeploymentStatusUpdate {
	dsu.mutation.Where(ps...)
	return dsu
}

// SetStatus sets the "status" field.
func (dsu *DeploymentStatusUpdate) SetStatus(s string) *DeploymentStatusUpdate {
	dsu.mutation.SetStatus(s)
	return dsu
}

// SetDescription sets the "description" field.
func (dsu *DeploymentStatusUpdate) SetDescription(s string) *DeploymentStatusUpdate {
	dsu.mutation.SetDescription(s)
	return dsu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dsu *DeploymentStatusUpdate) SetNillableDescription(s *string) *DeploymentStatusUpdate {
	if s != nil {
		dsu.SetDescription(*s)
	}
	return dsu
}

// ClearDescription clears the value of the "description" field.
func (dsu *DeploymentStatusUpdate) ClearDescription() *DeploymentStatusUpdate {
	dsu.mutation.ClearDescription()
	return dsu
}

// SetLogURL sets the "log_url" field.
func (dsu *DeploymentStatusUpdate) SetLogURL(s string) *DeploymentStatusUpdate {
	dsu.mutation.SetLogURL(s)
	return dsu
}

// SetNillableLogURL sets the "log_url" field if the given value is not nil.
func (dsu *DeploymentStatusUpdate) SetNillableLogURL(s *string) *DeploymentStatusUpdate {
	if s != nil {
		dsu.SetLogURL(*s)
	}
	return dsu
}

// ClearLogURL clears the value of the "log_url" field.
func (dsu *DeploymentStatusUpdate) ClearLogURL() *DeploymentStatusUpdate {
	dsu.mutation.ClearLogURL()
	return dsu
}

// SetCreatedAt sets the "created_at" field.
func (dsu *DeploymentStatusUpdate) SetCreatedAt(t time.Time) *DeploymentStatusUpdate {
	dsu.mutation.SetCreatedAt(t)
	return dsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dsu *DeploymentStatusUpdate) SetNillableCreatedAt(t *time.Time) *DeploymentStatusUpdate {
	if t != nil {
		dsu.SetCreatedAt(*t)
	}
	return dsu
}

// SetUpdatedAt sets the "updated_at" field.
func (dsu *DeploymentStatusUpdate) SetUpdatedAt(t time.Time) *DeploymentStatusUpdate {
	dsu.mutation.SetUpdatedAt(t)
	return dsu
}

// SetDeploymentID sets the "deployment_id" field.
func (dsu *DeploymentStatusUpdate) SetDeploymentID(i int) *DeploymentStatusUpdate {
	dsu.mutation.SetDeploymentID(i)
	return dsu
}

// SetRepoID sets the "repo_id" field.
func (dsu *DeploymentStatusUpdate) SetRepoID(i int64) *DeploymentStatusUpdate {
	dsu.mutation.SetRepoID(i)
	return dsu
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (dsu *DeploymentStatusUpdate) SetNillableRepoID(i *int64) *DeploymentStatusUpdate {
	if i != nil {
		dsu.SetRepoID(*i)
	}
	return dsu
}

// ClearRepoID clears the value of the "repo_id" field.
func (dsu *DeploymentStatusUpdate) ClearRepoID() *DeploymentStatusUpdate {
	dsu.mutation.ClearRepoID()
	return dsu
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (dsu *DeploymentStatusUpdate) SetDeployment(d *Deployment) *DeploymentStatusUpdate {
	return dsu.SetDeploymentID(d.ID)
}

// SetRepo sets the "repo" edge to the Repo entity.
func (dsu *DeploymentStatusUpdate) SetRepo(r *Repo) *DeploymentStatusUpdate {
	return dsu.SetRepoID(r.ID)
}

// AddEventIDs adds the "event" edge to the Event entity by IDs.
func (dsu *DeploymentStatusUpdate) AddEventIDs(ids ...int) *DeploymentStatusUpdate {
	dsu.mutation.AddEventIDs(ids...)
	return dsu
}

// AddEvent adds the "event" edges to the Event entity.
func (dsu *DeploymentStatusUpdate) AddEvent(e ...*Event) *DeploymentStatusUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dsu.AddEventIDs(ids...)
}

// Mutation returns the DeploymentStatusMutation object of the builder.
func (dsu *DeploymentStatusUpdate) Mutation() *DeploymentStatusMutation {
	return dsu.mutation
}

// ClearDeployment clears the "deployment" edge to the Deployment entity.
func (dsu *DeploymentStatusUpdate) ClearDeployment() *DeploymentStatusUpdate {
	dsu.mutation.ClearDeployment()
	return dsu
}

// ClearRepo clears the "repo" edge to the Repo entity.
func (dsu *DeploymentStatusUpdate) ClearRepo() *DeploymentStatusUpdate {
	dsu.mutation.ClearRepo()
	return dsu
}

// ClearEvent clears all "event" edges to the Event entity.
func (dsu *DeploymentStatusUpdate) ClearEvent() *DeploymentStatusUpdate {
	dsu.mutation.ClearEvent()
	return dsu
}

// RemoveEventIDs removes the "event" edge to Event entities by IDs.
func (dsu *DeploymentStatusUpdate) RemoveEventIDs(ids ...int) *DeploymentStatusUpdate {
	dsu.mutation.RemoveEventIDs(ids...)
	return dsu
}

// RemoveEvent removes "event" edges to Event entities.
func (dsu *DeploymentStatusUpdate) RemoveEvent(e ...*Event) *DeploymentStatusUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dsu.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dsu *DeploymentStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	dsu.defaults()
	if len(dsu.hooks) == 0 {
		if err = dsu.check(); err != nil {
			return 0, err
		}
		affected, err = dsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dsu.check(); err != nil {
				return 0, err
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
func (dsu *DeploymentStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := dsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dsu *DeploymentStatusUpdate) Exec(ctx context.Context) error {
	_, err := dsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsu *DeploymentStatusUpdate) ExecX(ctx context.Context) {
	if err := dsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dsu *DeploymentStatusUpdate) defaults() {
	if _, ok := dsu.mutation.UpdatedAt(); !ok {
		v := deploymentstatus.UpdateDefaultUpdatedAt()
		dsu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dsu *DeploymentStatusUpdate) check() error {
	if _, ok := dsu.mutation.DeploymentID(); dsu.mutation.DeploymentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeploymentStatus.deployment"`)
	}
	return nil
}

func (dsu *DeploymentStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deploymentstatus.Table,
			Columns: deploymentstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deploymentstatus.FieldID,
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
	if value, ok := dsu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatus.FieldStatus,
		})
	}
	if value, ok := dsu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatus.FieldDescription,
		})
	}
	if dsu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deploymentstatus.FieldDescription,
		})
	}
	if value, ok := dsu.mutation.LogURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatus.FieldLogURL,
		})
	}
	if dsu.mutation.LogURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deploymentstatus.FieldLogURL,
		})
	}
	if value, ok := dsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deploymentstatus.FieldCreatedAt,
		})
	}
	if value, ok := dsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deploymentstatus.FieldUpdatedAt,
		})
	}
	if dsu.mutation.DeploymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatus.DeploymentTable,
			Columns: []string{deploymentstatus.DeploymentColumn},
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
	if nodes := dsu.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatus.DeploymentTable,
			Columns: []string{deploymentstatus.DeploymentColumn},
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
	if dsu.mutation.RepoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatus.RepoTable,
			Columns: []string{deploymentstatus.RepoColumn},
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
	if nodes := dsu.mutation.RepoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatus.RepoTable,
			Columns: []string{deploymentstatus.RepoColumn},
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
	if dsu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deploymentstatus.EventTable,
			Columns: []string{deploymentstatus.EventColumn},
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
	if nodes := dsu.mutation.RemovedEventIDs(); len(nodes) > 0 && !dsu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deploymentstatus.EventTable,
			Columns: []string{deploymentstatus.EventColumn},
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
	if nodes := dsu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deploymentstatus.EventTable,
			Columns: []string{deploymentstatus.EventColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, dsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deploymentstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DeploymentStatusUpdateOne is the builder for updating a single DeploymentStatus entity.
type DeploymentStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeploymentStatusMutation
}

// SetStatus sets the "status" field.
func (dsuo *DeploymentStatusUpdateOne) SetStatus(s string) *DeploymentStatusUpdateOne {
	dsuo.mutation.SetStatus(s)
	return dsuo
}

// SetDescription sets the "description" field.
func (dsuo *DeploymentStatusUpdateOne) SetDescription(s string) *DeploymentStatusUpdateOne {
	dsuo.mutation.SetDescription(s)
	return dsuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dsuo *DeploymentStatusUpdateOne) SetNillableDescription(s *string) *DeploymentStatusUpdateOne {
	if s != nil {
		dsuo.SetDescription(*s)
	}
	return dsuo
}

// ClearDescription clears the value of the "description" field.
func (dsuo *DeploymentStatusUpdateOne) ClearDescription() *DeploymentStatusUpdateOne {
	dsuo.mutation.ClearDescription()
	return dsuo
}

// SetLogURL sets the "log_url" field.
func (dsuo *DeploymentStatusUpdateOne) SetLogURL(s string) *DeploymentStatusUpdateOne {
	dsuo.mutation.SetLogURL(s)
	return dsuo
}

// SetNillableLogURL sets the "log_url" field if the given value is not nil.
func (dsuo *DeploymentStatusUpdateOne) SetNillableLogURL(s *string) *DeploymentStatusUpdateOne {
	if s != nil {
		dsuo.SetLogURL(*s)
	}
	return dsuo
}

// ClearLogURL clears the value of the "log_url" field.
func (dsuo *DeploymentStatusUpdateOne) ClearLogURL() *DeploymentStatusUpdateOne {
	dsuo.mutation.ClearLogURL()
	return dsuo
}

// SetCreatedAt sets the "created_at" field.
func (dsuo *DeploymentStatusUpdateOne) SetCreatedAt(t time.Time) *DeploymentStatusUpdateOne {
	dsuo.mutation.SetCreatedAt(t)
	return dsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dsuo *DeploymentStatusUpdateOne) SetNillableCreatedAt(t *time.Time) *DeploymentStatusUpdateOne {
	if t != nil {
		dsuo.SetCreatedAt(*t)
	}
	return dsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (dsuo *DeploymentStatusUpdateOne) SetUpdatedAt(t time.Time) *DeploymentStatusUpdateOne {
	dsuo.mutation.SetUpdatedAt(t)
	return dsuo
}

// SetDeploymentID sets the "deployment_id" field.
func (dsuo *DeploymentStatusUpdateOne) SetDeploymentID(i int) *DeploymentStatusUpdateOne {
	dsuo.mutation.SetDeploymentID(i)
	return dsuo
}

// SetRepoID sets the "repo_id" field.
func (dsuo *DeploymentStatusUpdateOne) SetRepoID(i int64) *DeploymentStatusUpdateOne {
	dsuo.mutation.SetRepoID(i)
	return dsuo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (dsuo *DeploymentStatusUpdateOne) SetNillableRepoID(i *int64) *DeploymentStatusUpdateOne {
	if i != nil {
		dsuo.SetRepoID(*i)
	}
	return dsuo
}

// ClearRepoID clears the value of the "repo_id" field.
func (dsuo *DeploymentStatusUpdateOne) ClearRepoID() *DeploymentStatusUpdateOne {
	dsuo.mutation.ClearRepoID()
	return dsuo
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (dsuo *DeploymentStatusUpdateOne) SetDeployment(d *Deployment) *DeploymentStatusUpdateOne {
	return dsuo.SetDeploymentID(d.ID)
}

// SetRepo sets the "repo" edge to the Repo entity.
func (dsuo *DeploymentStatusUpdateOne) SetRepo(r *Repo) *DeploymentStatusUpdateOne {
	return dsuo.SetRepoID(r.ID)
}

// AddEventIDs adds the "event" edge to the Event entity by IDs.
func (dsuo *DeploymentStatusUpdateOne) AddEventIDs(ids ...int) *DeploymentStatusUpdateOne {
	dsuo.mutation.AddEventIDs(ids...)
	return dsuo
}

// AddEvent adds the "event" edges to the Event entity.
func (dsuo *DeploymentStatusUpdateOne) AddEvent(e ...*Event) *DeploymentStatusUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dsuo.AddEventIDs(ids...)
}

// Mutation returns the DeploymentStatusMutation object of the builder.
func (dsuo *DeploymentStatusUpdateOne) Mutation() *DeploymentStatusMutation {
	return dsuo.mutation
}

// ClearDeployment clears the "deployment" edge to the Deployment entity.
func (dsuo *DeploymentStatusUpdateOne) ClearDeployment() *DeploymentStatusUpdateOne {
	dsuo.mutation.ClearDeployment()
	return dsuo
}

// ClearRepo clears the "repo" edge to the Repo entity.
func (dsuo *DeploymentStatusUpdateOne) ClearRepo() *DeploymentStatusUpdateOne {
	dsuo.mutation.ClearRepo()
	return dsuo
}

// ClearEvent clears all "event" edges to the Event entity.
func (dsuo *DeploymentStatusUpdateOne) ClearEvent() *DeploymentStatusUpdateOne {
	dsuo.mutation.ClearEvent()
	return dsuo
}

// RemoveEventIDs removes the "event" edge to Event entities by IDs.
func (dsuo *DeploymentStatusUpdateOne) RemoveEventIDs(ids ...int) *DeploymentStatusUpdateOne {
	dsuo.mutation.RemoveEventIDs(ids...)
	return dsuo
}

// RemoveEvent removes "event" edges to Event entities.
func (dsuo *DeploymentStatusUpdateOne) RemoveEvent(e ...*Event) *DeploymentStatusUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dsuo.RemoveEventIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dsuo *DeploymentStatusUpdateOne) Select(field string, fields ...string) *DeploymentStatusUpdateOne {
	dsuo.fields = append([]string{field}, fields...)
	return dsuo
}

// Save executes the query and returns the updated DeploymentStatus entity.
func (dsuo *DeploymentStatusUpdateOne) Save(ctx context.Context) (*DeploymentStatus, error) {
	var (
		err  error
		node *DeploymentStatus
	)
	dsuo.defaults()
	if len(dsuo.hooks) == 0 {
		if err = dsuo.check(); err != nil {
			return nil, err
		}
		node, err = dsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dsuo.check(); err != nil {
				return nil, err
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
func (dsuo *DeploymentStatusUpdateOne) SaveX(ctx context.Context) *DeploymentStatus {
	node, err := dsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dsuo *DeploymentStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := dsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsuo *DeploymentStatusUpdateOne) ExecX(ctx context.Context) {
	if err := dsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dsuo *DeploymentStatusUpdateOne) defaults() {
	if _, ok := dsuo.mutation.UpdatedAt(); !ok {
		v := deploymentstatus.UpdateDefaultUpdatedAt()
		dsuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dsuo *DeploymentStatusUpdateOne) check() error {
	if _, ok := dsuo.mutation.DeploymentID(); dsuo.mutation.DeploymentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeploymentStatus.deployment"`)
	}
	return nil
}

func (dsuo *DeploymentStatusUpdateOne) sqlSave(ctx context.Context) (_node *DeploymentStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deploymentstatus.Table,
			Columns: deploymentstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deploymentstatus.FieldID,
			},
		},
	}
	id, ok := dsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeploymentStatus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deploymentstatus.FieldID)
		for _, f := range fields {
			if !deploymentstatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deploymentstatus.FieldID {
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
	if value, ok := dsuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatus.FieldStatus,
		})
	}
	if value, ok := dsuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatus.FieldDescription,
		})
	}
	if dsuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deploymentstatus.FieldDescription,
		})
	}
	if value, ok := dsuo.mutation.LogURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentstatus.FieldLogURL,
		})
	}
	if dsuo.mutation.LogURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deploymentstatus.FieldLogURL,
		})
	}
	if value, ok := dsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deploymentstatus.FieldCreatedAt,
		})
	}
	if value, ok := dsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deploymentstatus.FieldUpdatedAt,
		})
	}
	if dsuo.mutation.DeploymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatus.DeploymentTable,
			Columns: []string{deploymentstatus.DeploymentColumn},
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
	if nodes := dsuo.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatus.DeploymentTable,
			Columns: []string{deploymentstatus.DeploymentColumn},
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
	if dsuo.mutation.RepoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatus.RepoTable,
			Columns: []string{deploymentstatus.RepoColumn},
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
	if nodes := dsuo.mutation.RepoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatus.RepoTable,
			Columns: []string{deploymentstatus.RepoColumn},
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
	if dsuo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deploymentstatus.EventTable,
			Columns: []string{deploymentstatus.EventColumn},
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
	if nodes := dsuo.mutation.RemovedEventIDs(); len(nodes) > 0 && !dsuo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deploymentstatus.EventTable,
			Columns: []string{deploymentstatus.EventColumn},
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
	if nodes := dsuo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deploymentstatus.EventTable,
			Columns: []string{deploymentstatus.EventColumn},
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
	_node = &DeploymentStatus{config: dsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deploymentstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
