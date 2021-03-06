// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/model/ent/deploymentstatistics"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
	"github.com/gitploy-io/gitploy/model/ent/repo"
)

// DeploymentStatisticsQuery is the builder for querying DeploymentStatistics entities.
type DeploymentStatisticsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeploymentStatistics
	// eager-loading edges.
	withRepo  *RepoQuery
	modifiers []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeploymentStatisticsQuery builder.
func (dsq *DeploymentStatisticsQuery) Where(ps ...predicate.DeploymentStatistics) *DeploymentStatisticsQuery {
	dsq.predicates = append(dsq.predicates, ps...)
	return dsq
}

// Limit adds a limit step to the query.
func (dsq *DeploymentStatisticsQuery) Limit(limit int) *DeploymentStatisticsQuery {
	dsq.limit = &limit
	return dsq
}

// Offset adds an offset step to the query.
func (dsq *DeploymentStatisticsQuery) Offset(offset int) *DeploymentStatisticsQuery {
	dsq.offset = &offset
	return dsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dsq *DeploymentStatisticsQuery) Unique(unique bool) *DeploymentStatisticsQuery {
	dsq.unique = &unique
	return dsq
}

// Order adds an order step to the query.
func (dsq *DeploymentStatisticsQuery) Order(o ...OrderFunc) *DeploymentStatisticsQuery {
	dsq.order = append(dsq.order, o...)
	return dsq
}

// QueryRepo chains the current query on the "repo" edge.
func (dsq *DeploymentStatisticsQuery) QueryRepo() *RepoQuery {
	query := &RepoQuery{config: dsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deploymentstatistics.Table, deploymentstatistics.FieldID, selector),
			sqlgraph.To(repo.Table, repo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deploymentstatistics.RepoTable, deploymentstatistics.RepoColumn),
		)
		fromU = sqlgraph.SetNeighbors(dsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeploymentStatistics entity from the query.
// Returns a *NotFoundError when no DeploymentStatistics was found.
func (dsq *DeploymentStatisticsQuery) First(ctx context.Context) (*DeploymentStatistics, error) {
	nodes, err := dsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deploymentstatistics.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dsq *DeploymentStatisticsQuery) FirstX(ctx context.Context) *DeploymentStatistics {
	node, err := dsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeploymentStatistics ID from the query.
// Returns a *NotFoundError when no DeploymentStatistics ID was found.
func (dsq *DeploymentStatisticsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deploymentstatistics.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dsq *DeploymentStatisticsQuery) FirstIDX(ctx context.Context) int {
	id, err := dsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeploymentStatistics entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeploymentStatistics entity is found.
// Returns a *NotFoundError when no DeploymentStatistics entities are found.
func (dsq *DeploymentStatisticsQuery) Only(ctx context.Context) (*DeploymentStatistics, error) {
	nodes, err := dsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deploymentstatistics.Label}
	default:
		return nil, &NotSingularError{deploymentstatistics.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dsq *DeploymentStatisticsQuery) OnlyX(ctx context.Context) *DeploymentStatistics {
	node, err := dsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeploymentStatistics ID in the query.
// Returns a *NotSingularError when more than one DeploymentStatistics ID is found.
// Returns a *NotFoundError when no entities are found.
func (dsq *DeploymentStatisticsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = &NotSingularError{deploymentstatistics.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dsq *DeploymentStatisticsQuery) OnlyIDX(ctx context.Context) int {
	id, err := dsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeploymentStatisticsSlice.
func (dsq *DeploymentStatisticsQuery) All(ctx context.Context) ([]*DeploymentStatistics, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dsq *DeploymentStatisticsQuery) AllX(ctx context.Context) []*DeploymentStatistics {
	nodes, err := dsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeploymentStatistics IDs.
func (dsq *DeploymentStatisticsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dsq.Select(deploymentstatistics.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dsq *DeploymentStatisticsQuery) IDsX(ctx context.Context) []int {
	ids, err := dsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dsq *DeploymentStatisticsQuery) Count(ctx context.Context) (int, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dsq *DeploymentStatisticsQuery) CountX(ctx context.Context) int {
	count, err := dsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dsq *DeploymentStatisticsQuery) Exist(ctx context.Context) (bool, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dsq *DeploymentStatisticsQuery) ExistX(ctx context.Context) bool {
	exist, err := dsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeploymentStatisticsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dsq *DeploymentStatisticsQuery) Clone() *DeploymentStatisticsQuery {
	if dsq == nil {
		return nil
	}
	return &DeploymentStatisticsQuery{
		config:     dsq.config,
		limit:      dsq.limit,
		offset:     dsq.offset,
		order:      append([]OrderFunc{}, dsq.order...),
		predicates: append([]predicate.DeploymentStatistics{}, dsq.predicates...),
		withRepo:   dsq.withRepo.Clone(),
		// clone intermediate query.
		sql:    dsq.sql.Clone(),
		path:   dsq.path,
		unique: dsq.unique,
	}
}

// WithRepo tells the query-builder to eager-load the nodes that are connected to
// the "repo" edge. The optional arguments are used to configure the query builder of the edge.
func (dsq *DeploymentStatisticsQuery) WithRepo(opts ...func(*RepoQuery)) *DeploymentStatisticsQuery {
	query := &RepoQuery{config: dsq.config}
	for _, opt := range opts {
		opt(query)
	}
	dsq.withRepo = query
	return dsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Env string `json:"env"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeploymentStatistics.Query().
//		GroupBy(deploymentstatistics.FieldEnv).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dsq *DeploymentStatisticsQuery) GroupBy(field string, fields ...string) *DeploymentStatisticsGroupBy {
	group := &DeploymentStatisticsGroupBy{config: dsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Env string `json:"env"`
//	}
//
//	client.DeploymentStatistics.Query().
//		Select(deploymentstatistics.FieldEnv).
//		Scan(ctx, &v)
//
func (dsq *DeploymentStatisticsQuery) Select(fields ...string) *DeploymentStatisticsSelect {
	dsq.fields = append(dsq.fields, fields...)
	return &DeploymentStatisticsSelect{DeploymentStatisticsQuery: dsq}
}

func (dsq *DeploymentStatisticsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dsq.fields {
		if !deploymentstatistics.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dsq.path != nil {
		prev, err := dsq.path(ctx)
		if err != nil {
			return err
		}
		dsq.sql = prev
	}
	return nil
}

func (dsq *DeploymentStatisticsQuery) sqlAll(ctx context.Context) ([]*DeploymentStatistics, error) {
	var (
		nodes       = []*DeploymentStatistics{}
		_spec       = dsq.querySpec()
		loadedTypes = [1]bool{
			dsq.withRepo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DeploymentStatistics{config: dsq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dsq.modifiers) > 0 {
		_spec.Modifiers = dsq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, dsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dsq.withRepo; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*DeploymentStatistics)
		for i := range nodes {
			fk := nodes[i].RepoID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(repo.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "repo_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Repo = n
			}
		}
	}

	return nodes, nil
}

func (dsq *DeploymentStatisticsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dsq.querySpec()
	if len(dsq.modifiers) > 0 {
		_spec.Modifiers = dsq.modifiers
	}
	_spec.Node.Columns = dsq.fields
	if len(dsq.fields) > 0 {
		_spec.Unique = dsq.unique != nil && *dsq.unique
	}
	return sqlgraph.CountNodes(ctx, dsq.driver, _spec)
}

func (dsq *DeploymentStatisticsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dsq *DeploymentStatisticsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deploymentstatistics.Table,
			Columns: deploymentstatistics.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deploymentstatistics.FieldID,
			},
		},
		From:   dsq.sql,
		Unique: true,
	}
	if unique := dsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deploymentstatistics.FieldID)
		for i := range fields {
			if fields[i] != deploymentstatistics.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dsq *DeploymentStatisticsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dsq.driver.Dialect())
	t1 := builder.Table(deploymentstatistics.Table)
	columns := dsq.fields
	if len(columns) == 0 {
		columns = deploymentstatistics.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dsq.sql != nil {
		selector = dsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dsq.unique != nil && *dsq.unique {
		selector.Distinct()
	}
	for _, m := range dsq.modifiers {
		m(selector)
	}
	for _, p := range dsq.predicates {
		p(selector)
	}
	for _, p := range dsq.order {
		p(selector)
	}
	if offset := dsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dsq *DeploymentStatisticsQuery) ForUpdate(opts ...sql.LockOption) *DeploymentStatisticsQuery {
	if dsq.driver.Dialect() == dialect.Postgres {
		dsq.Unique(false)
	}
	dsq.modifiers = append(dsq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dsq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dsq *DeploymentStatisticsQuery) ForShare(opts ...sql.LockOption) *DeploymentStatisticsQuery {
	if dsq.driver.Dialect() == dialect.Postgres {
		dsq.Unique(false)
	}
	dsq.modifiers = append(dsq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dsq
}

// DeploymentStatisticsGroupBy is the group-by builder for DeploymentStatistics entities.
type DeploymentStatisticsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dsgb *DeploymentStatisticsGroupBy) Aggregate(fns ...AggregateFunc) *DeploymentStatisticsGroupBy {
	dsgb.fns = append(dsgb.fns, fns...)
	return dsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dsgb *DeploymentStatisticsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dsgb.path(ctx)
	if err != nil {
		return err
	}
	dsgb.sql = query
	return dsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatisticsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatisticsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) StringsX(ctx context.Context) []string {
	v, err := dsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatisticsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatisticsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) StringX(ctx context.Context) string {
	v, err := dsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatisticsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatisticsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) IntsX(ctx context.Context) []int {
	v, err := dsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatisticsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatisticsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) IntX(ctx context.Context) int {
	v, err := dsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatisticsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatisticsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatisticsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatisticsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatisticsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatisticsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatisticsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatisticsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dsgb *DeploymentStatisticsGroupBy) BoolX(ctx context.Context) bool {
	v, err := dsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dsgb *DeploymentStatisticsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dsgb.fields {
		if !deploymentstatistics.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dsgb *DeploymentStatisticsGroupBy) sqlQuery() *sql.Selector {
	selector := dsgb.sql.Select()
	aggregation := make([]string, 0, len(dsgb.fns))
	for _, fn := range dsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dsgb.fields)+len(dsgb.fns))
		for _, f := range dsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dsgb.fields...)...)
}

// DeploymentStatisticsSelect is the builder for selecting fields of DeploymentStatistics entities.
type DeploymentStatisticsSelect struct {
	*DeploymentStatisticsQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dss *DeploymentStatisticsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dss.prepareQuery(ctx); err != nil {
		return err
	}
	dss.sql = dss.DeploymentStatisticsQuery.sqlQuery(ctx)
	return dss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatisticsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatisticsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) StringsX(ctx context.Context) []string {
	v, err := dss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatisticsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatisticsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) StringX(ctx context.Context) string {
	v, err := dss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatisticsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatisticsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) IntsX(ctx context.Context) []int {
	v, err := dss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatisticsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatisticsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) IntX(ctx context.Context) int {
	v, err := dss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatisticsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatisticsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatisticsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatisticsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) Float64X(ctx context.Context) float64 {
	v, err := dss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatisticsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatisticsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) BoolsX(ctx context.Context) []bool {
	v, err := dss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatisticsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatistics.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatisticsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dss *DeploymentStatisticsSelect) BoolX(ctx context.Context) bool {
	v, err := dss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dss *DeploymentStatisticsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dss.sql.Query()
	if err := dss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
