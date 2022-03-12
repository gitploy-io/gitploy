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
	"github.com/gitploy-io/gitploy/model/ent/deploymentstatistics"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
	"github.com/gitploy-io/gitploy/model/ent/repo"
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

// SetRollbackCount sets the "rollback_count" field.
func (dsu *DeploymentStatisticsUpdate) SetRollbackCount(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.ResetRollbackCount()
	dsu.mutation.SetRollbackCount(i)
	return dsu
}

// SetNillableRollbackCount sets the "rollback_count" field if the given value is not nil.
func (dsu *DeploymentStatisticsUpdate) SetNillableRollbackCount(i *int) *DeploymentStatisticsUpdate {
	if i != nil {
		dsu.SetRollbackCount(*i)
	}
	return dsu
}

// AddRollbackCount adds i to the "rollback_count" field.
func (dsu *DeploymentStatisticsUpdate) AddRollbackCount(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.AddRollbackCount(i)
	return dsu
}

// SetAdditions sets the "additions" field.
func (dsu *DeploymentStatisticsUpdate) SetAdditions(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.ResetAdditions()
	dsu.mutation.SetAdditions(i)
	return dsu
}

// SetNillableAdditions sets the "additions" field if the given value is not nil.
func (dsu *DeploymentStatisticsUpdate) SetNillableAdditions(i *int) *DeploymentStatisticsUpdate {
	if i != nil {
		dsu.SetAdditions(*i)
	}
	return dsu
}

// AddAdditions adds i to the "additions" field.
func (dsu *DeploymentStatisticsUpdate) AddAdditions(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.AddAdditions(i)
	return dsu
}

// SetDeletions sets the "deletions" field.
func (dsu *DeploymentStatisticsUpdate) SetDeletions(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.ResetDeletions()
	dsu.mutation.SetDeletions(i)
	return dsu
}

// SetNillableDeletions sets the "deletions" field if the given value is not nil.
func (dsu *DeploymentStatisticsUpdate) SetNillableDeletions(i *int) *DeploymentStatisticsUpdate {
	if i != nil {
		dsu.SetDeletions(*i)
	}
	return dsu
}

// AddDeletions adds i to the "deletions" field.
func (dsu *DeploymentStatisticsUpdate) AddDeletions(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.AddDeletions(i)
	return dsu
}

// SetChanges sets the "changes" field.
func (dsu *DeploymentStatisticsUpdate) SetChanges(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.ResetChanges()
	dsu.mutation.SetChanges(i)
	return dsu
}

// SetNillableChanges sets the "changes" field if the given value is not nil.
func (dsu *DeploymentStatisticsUpdate) SetNillableChanges(i *int) *DeploymentStatisticsUpdate {
	if i != nil {
		dsu.SetChanges(*i)
	}
	return dsu
}

// AddChanges adds i to the "changes" field.
func (dsu *DeploymentStatisticsUpdate) AddChanges(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.AddChanges(i)
	return dsu
}

// SetLeadTimeSeconds sets the "lead_time_seconds" field.
func (dsu *DeploymentStatisticsUpdate) SetLeadTimeSeconds(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.ResetLeadTimeSeconds()
	dsu.mutation.SetLeadTimeSeconds(i)
	return dsu
}

// SetNillableLeadTimeSeconds sets the "lead_time_seconds" field if the given value is not nil.
func (dsu *DeploymentStatisticsUpdate) SetNillableLeadTimeSeconds(i *int) *DeploymentStatisticsUpdate {
	if i != nil {
		dsu.SetLeadTimeSeconds(*i)
	}
	return dsu
}

// AddLeadTimeSeconds adds i to the "lead_time_seconds" field.
func (dsu *DeploymentStatisticsUpdate) AddLeadTimeSeconds(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.AddLeadTimeSeconds(i)
	return dsu
}

// SetCommitCount sets the "commit_count" field.
func (dsu *DeploymentStatisticsUpdate) SetCommitCount(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.ResetCommitCount()
	dsu.mutation.SetCommitCount(i)
	return dsu
}

// SetNillableCommitCount sets the "commit_count" field if the given value is not nil.
func (dsu *DeploymentStatisticsUpdate) SetNillableCommitCount(i *int) *DeploymentStatisticsUpdate {
	if i != nil {
		dsu.SetCommitCount(*i)
	}
	return dsu
}

// AddCommitCount adds i to the "commit_count" field.
func (dsu *DeploymentStatisticsUpdate) AddCommitCount(i int) *DeploymentStatisticsUpdate {
	dsu.mutation.AddCommitCount(i)
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

// SetRepoID sets the "repo_id" field.
func (dsu *DeploymentStatisticsUpdate) SetRepoID(i int64) *DeploymentStatisticsUpdate {
	dsu.mutation.SetRepoID(i)
	return dsu
}

// SetRepo sets the "repo" edge to the Repo entity.
func (dsu *DeploymentStatisticsUpdate) SetRepo(r *Repo) *DeploymentStatisticsUpdate {
	return dsu.SetRepoID(r.ID)
}

// Mutation returns the DeploymentStatisticsMutation object of the builder.
func (dsu *DeploymentStatisticsUpdate) Mutation() *DeploymentStatisticsMutation {
	return dsu.mutation
}

// ClearRepo clears the "repo" edge to the Repo entity.
func (dsu *DeploymentStatisticsUpdate) ClearRepo() *DeploymentStatisticsUpdate {
	dsu.mutation.ClearRepo()
	return dsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dsu *DeploymentStatisticsUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*DeploymentStatisticsMutation)
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

// check runs all checks and user-defined validators on the builder.
func (dsu *DeploymentStatisticsUpdate) check() error {
	if _, ok := dsu.mutation.RepoID(); dsu.mutation.RepoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeploymentStatistics.repo"`)
	}
	return nil
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
	if value, ok := dsu.mutation.RollbackCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldRollbackCount,
		})
	}
	if value, ok := dsu.mutation.AddedRollbackCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldRollbackCount,
		})
	}
	if value, ok := dsu.mutation.Additions(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldAdditions,
		})
	}
	if value, ok := dsu.mutation.AddedAdditions(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldAdditions,
		})
	}
	if value, ok := dsu.mutation.Deletions(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldDeletions,
		})
	}
	if value, ok := dsu.mutation.AddedDeletions(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldDeletions,
		})
	}
	if value, ok := dsu.mutation.Changes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldChanges,
		})
	}
	if value, ok := dsu.mutation.AddedChanges(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldChanges,
		})
	}
	if value, ok := dsu.mutation.LeadTimeSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldLeadTimeSeconds,
		})
	}
	if value, ok := dsu.mutation.AddedLeadTimeSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldLeadTimeSeconds,
		})
	}
	if value, ok := dsu.mutation.CommitCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldCommitCount,
		})
	}
	if value, ok := dsu.mutation.AddedCommitCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldCommitCount,
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
	if dsu.mutation.RepoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatistics.RepoTable,
			Columns: []string{deploymentstatistics.RepoColumn},
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
			Table:   deploymentstatistics.RepoTable,
			Columns: []string{deploymentstatistics.RepoColumn},
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

// SetRollbackCount sets the "rollback_count" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetRollbackCount(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.ResetRollbackCount()
	dsuo.mutation.SetRollbackCount(i)
	return dsuo
}

// SetNillableRollbackCount sets the "rollback_count" field if the given value is not nil.
func (dsuo *DeploymentStatisticsUpdateOne) SetNillableRollbackCount(i *int) *DeploymentStatisticsUpdateOne {
	if i != nil {
		dsuo.SetRollbackCount(*i)
	}
	return dsuo
}

// AddRollbackCount adds i to the "rollback_count" field.
func (dsuo *DeploymentStatisticsUpdateOne) AddRollbackCount(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.AddRollbackCount(i)
	return dsuo
}

// SetAdditions sets the "additions" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetAdditions(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.ResetAdditions()
	dsuo.mutation.SetAdditions(i)
	return dsuo
}

// SetNillableAdditions sets the "additions" field if the given value is not nil.
func (dsuo *DeploymentStatisticsUpdateOne) SetNillableAdditions(i *int) *DeploymentStatisticsUpdateOne {
	if i != nil {
		dsuo.SetAdditions(*i)
	}
	return dsuo
}

// AddAdditions adds i to the "additions" field.
func (dsuo *DeploymentStatisticsUpdateOne) AddAdditions(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.AddAdditions(i)
	return dsuo
}

// SetDeletions sets the "deletions" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetDeletions(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.ResetDeletions()
	dsuo.mutation.SetDeletions(i)
	return dsuo
}

// SetNillableDeletions sets the "deletions" field if the given value is not nil.
func (dsuo *DeploymentStatisticsUpdateOne) SetNillableDeletions(i *int) *DeploymentStatisticsUpdateOne {
	if i != nil {
		dsuo.SetDeletions(*i)
	}
	return dsuo
}

// AddDeletions adds i to the "deletions" field.
func (dsuo *DeploymentStatisticsUpdateOne) AddDeletions(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.AddDeletions(i)
	return dsuo
}

// SetChanges sets the "changes" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetChanges(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.ResetChanges()
	dsuo.mutation.SetChanges(i)
	return dsuo
}

// SetNillableChanges sets the "changes" field if the given value is not nil.
func (dsuo *DeploymentStatisticsUpdateOne) SetNillableChanges(i *int) *DeploymentStatisticsUpdateOne {
	if i != nil {
		dsuo.SetChanges(*i)
	}
	return dsuo
}

// AddChanges adds i to the "changes" field.
func (dsuo *DeploymentStatisticsUpdateOne) AddChanges(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.AddChanges(i)
	return dsuo
}

// SetLeadTimeSeconds sets the "lead_time_seconds" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetLeadTimeSeconds(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.ResetLeadTimeSeconds()
	dsuo.mutation.SetLeadTimeSeconds(i)
	return dsuo
}

// SetNillableLeadTimeSeconds sets the "lead_time_seconds" field if the given value is not nil.
func (dsuo *DeploymentStatisticsUpdateOne) SetNillableLeadTimeSeconds(i *int) *DeploymentStatisticsUpdateOne {
	if i != nil {
		dsuo.SetLeadTimeSeconds(*i)
	}
	return dsuo
}

// AddLeadTimeSeconds adds i to the "lead_time_seconds" field.
func (dsuo *DeploymentStatisticsUpdateOne) AddLeadTimeSeconds(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.AddLeadTimeSeconds(i)
	return dsuo
}

// SetCommitCount sets the "commit_count" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetCommitCount(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.ResetCommitCount()
	dsuo.mutation.SetCommitCount(i)
	return dsuo
}

// SetNillableCommitCount sets the "commit_count" field if the given value is not nil.
func (dsuo *DeploymentStatisticsUpdateOne) SetNillableCommitCount(i *int) *DeploymentStatisticsUpdateOne {
	if i != nil {
		dsuo.SetCommitCount(*i)
	}
	return dsuo
}

// AddCommitCount adds i to the "commit_count" field.
func (dsuo *DeploymentStatisticsUpdateOne) AddCommitCount(i int) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.AddCommitCount(i)
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

// SetRepoID sets the "repo_id" field.
func (dsuo *DeploymentStatisticsUpdateOne) SetRepoID(i int64) *DeploymentStatisticsUpdateOne {
	dsuo.mutation.SetRepoID(i)
	return dsuo
}

// SetRepo sets the "repo" edge to the Repo entity.
func (dsuo *DeploymentStatisticsUpdateOne) SetRepo(r *Repo) *DeploymentStatisticsUpdateOne {
	return dsuo.SetRepoID(r.ID)
}

// Mutation returns the DeploymentStatisticsMutation object of the builder.
func (dsuo *DeploymentStatisticsUpdateOne) Mutation() *DeploymentStatisticsMutation {
	return dsuo.mutation
}

// ClearRepo clears the "repo" edge to the Repo entity.
func (dsuo *DeploymentStatisticsUpdateOne) ClearRepo() *DeploymentStatisticsUpdateOne {
	dsuo.mutation.ClearRepo()
	return dsuo
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
		if err = dsuo.check(); err != nil {
			return nil, err
		}
		node, err = dsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentStatisticsMutation)
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

// check runs all checks and user-defined validators on the builder.
func (dsuo *DeploymentStatisticsUpdateOne) check() error {
	if _, ok := dsuo.mutation.RepoID(); dsuo.mutation.RepoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeploymentStatistics.repo"`)
	}
	return nil
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
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeploymentStatistics.id" for update`)}
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
	if value, ok := dsuo.mutation.RollbackCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldRollbackCount,
		})
	}
	if value, ok := dsuo.mutation.AddedRollbackCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldRollbackCount,
		})
	}
	if value, ok := dsuo.mutation.Additions(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldAdditions,
		})
	}
	if value, ok := dsuo.mutation.AddedAdditions(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldAdditions,
		})
	}
	if value, ok := dsuo.mutation.Deletions(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldDeletions,
		})
	}
	if value, ok := dsuo.mutation.AddedDeletions(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldDeletions,
		})
	}
	if value, ok := dsuo.mutation.Changes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldChanges,
		})
	}
	if value, ok := dsuo.mutation.AddedChanges(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldChanges,
		})
	}
	if value, ok := dsuo.mutation.LeadTimeSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldLeadTimeSeconds,
		})
	}
	if value, ok := dsuo.mutation.AddedLeadTimeSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldLeadTimeSeconds,
		})
	}
	if value, ok := dsuo.mutation.CommitCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldCommitCount,
		})
	}
	if value, ok := dsuo.mutation.AddedCommitCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deploymentstatistics.FieldCommitCount,
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
	if dsuo.mutation.RepoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deploymentstatistics.RepoTable,
			Columns: []string{deploymentstatistics.RepoColumn},
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
			Table:   deploymentstatistics.RepoTable,
			Columns: []string{deploymentstatistics.RepoColumn},
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
