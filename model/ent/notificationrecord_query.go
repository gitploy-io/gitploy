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
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/model/ent/notificationrecord"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
)

// NotificationRecordQuery is the builder for querying NotificationRecord entities.
type NotificationRecordQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NotificationRecord
	// eager-loading edges.
	withEvent *EventQuery
	modifiers []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotificationRecordQuery builder.
func (nrq *NotificationRecordQuery) Where(ps ...predicate.NotificationRecord) *NotificationRecordQuery {
	nrq.predicates = append(nrq.predicates, ps...)
	return nrq
}

// Limit adds a limit step to the query.
func (nrq *NotificationRecordQuery) Limit(limit int) *NotificationRecordQuery {
	nrq.limit = &limit
	return nrq
}

// Offset adds an offset step to the query.
func (nrq *NotificationRecordQuery) Offset(offset int) *NotificationRecordQuery {
	nrq.offset = &offset
	return nrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nrq *NotificationRecordQuery) Unique(unique bool) *NotificationRecordQuery {
	nrq.unique = &unique
	return nrq
}

// Order adds an order step to the query.
func (nrq *NotificationRecordQuery) Order(o ...OrderFunc) *NotificationRecordQuery {
	nrq.order = append(nrq.order, o...)
	return nrq
}

// QueryEvent chains the current query on the "event" edge.
func (nrq *NotificationRecordQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: nrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notificationrecord.Table, notificationrecord.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, notificationrecord.EventTable, notificationrecord.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(nrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NotificationRecord entity from the query.
// Returns a *NotFoundError when no NotificationRecord was found.
func (nrq *NotificationRecordQuery) First(ctx context.Context) (*NotificationRecord, error) {
	nodes, err := nrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notificationrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nrq *NotificationRecordQuery) FirstX(ctx context.Context) *NotificationRecord {
	node, err := nrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NotificationRecord ID from the query.
// Returns a *NotFoundError when no NotificationRecord ID was found.
func (nrq *NotificationRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notificationrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nrq *NotificationRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := nrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NotificationRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NotificationRecord entity is not found.
// Returns a *NotFoundError when no NotificationRecord entities are found.
func (nrq *NotificationRecordQuery) Only(ctx context.Context) (*NotificationRecord, error) {
	nodes, err := nrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notificationrecord.Label}
	default:
		return nil, &NotSingularError{notificationrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nrq *NotificationRecordQuery) OnlyX(ctx context.Context) *NotificationRecord {
	node, err := nrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NotificationRecord ID in the query.
// Returns a *NotSingularError when exactly one NotificationRecord ID is not found.
// Returns a *NotFoundError when no entities are found.
func (nrq *NotificationRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = &NotSingularError{notificationrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nrq *NotificationRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := nrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NotificationRecords.
func (nrq *NotificationRecordQuery) All(ctx context.Context) ([]*NotificationRecord, error) {
	if err := nrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nrq *NotificationRecordQuery) AllX(ctx context.Context) []*NotificationRecord {
	nodes, err := nrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NotificationRecord IDs.
func (nrq *NotificationRecordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := nrq.Select(notificationrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nrq *NotificationRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := nrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nrq *NotificationRecordQuery) Count(ctx context.Context) (int, error) {
	if err := nrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nrq *NotificationRecordQuery) CountX(ctx context.Context) int {
	count, err := nrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nrq *NotificationRecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := nrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nrq *NotificationRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := nrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotificationRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nrq *NotificationRecordQuery) Clone() *NotificationRecordQuery {
	if nrq == nil {
		return nil
	}
	return &NotificationRecordQuery{
		config:     nrq.config,
		limit:      nrq.limit,
		offset:     nrq.offset,
		order:      append([]OrderFunc{}, nrq.order...),
		predicates: append([]predicate.NotificationRecord{}, nrq.predicates...),
		withEvent:  nrq.withEvent.Clone(),
		// clone intermediate query.
		sql:  nrq.sql.Clone(),
		path: nrq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (nrq *NotificationRecordQuery) WithEvent(opts ...func(*EventQuery)) *NotificationRecordQuery {
	query := &EventQuery{config: nrq.config}
	for _, opt := range opts {
		opt(query)
	}
	nrq.withEvent = query
	return nrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventID int `json:"event_id"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NotificationRecord.Query().
//		GroupBy(notificationrecord.FieldEventID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nrq *NotificationRecordQuery) GroupBy(field string, fields ...string) *NotificationRecordGroupBy {
	group := &NotificationRecordGroupBy{config: nrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nrq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventID int `json:"event_id"`
//	}
//
//	client.NotificationRecord.Query().
//		Select(notificationrecord.FieldEventID).
//		Scan(ctx, &v)
//
func (nrq *NotificationRecordQuery) Select(fields ...string) *NotificationRecordSelect {
	nrq.fields = append(nrq.fields, fields...)
	return &NotificationRecordSelect{NotificationRecordQuery: nrq}
}

func (nrq *NotificationRecordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nrq.fields {
		if !notificationrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nrq.path != nil {
		prev, err := nrq.path(ctx)
		if err != nil {
			return err
		}
		nrq.sql = prev
	}
	return nil
}

func (nrq *NotificationRecordQuery) sqlAll(ctx context.Context) ([]*NotificationRecord, error) {
	var (
		nodes       = []*NotificationRecord{}
		_spec       = nrq.querySpec()
		loadedTypes = [1]bool{
			nrq.withEvent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NotificationRecord{config: nrq.config}
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
	if len(nrq.modifiers) > 0 {
		_spec.Modifiers = nrq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, nrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := nrq.withEvent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*NotificationRecord)
		for i := range nodes {
			fk := nodes[i].EventID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(event.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "event_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Event = n
			}
		}
	}

	return nodes, nil
}

func (nrq *NotificationRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nrq.querySpec()
	if len(nrq.modifiers) > 0 {
		_spec.Modifiers = nrq.modifiers
	}
	return sqlgraph.CountNodes(ctx, nrq.driver, _spec)
}

func (nrq *NotificationRecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nrq *NotificationRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notificationrecord.Table,
			Columns: notificationrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notificationrecord.FieldID,
			},
		},
		From:   nrq.sql,
		Unique: true,
	}
	if unique := nrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationrecord.FieldID)
		for i := range fields {
			if fields[i] != notificationrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nrq *NotificationRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nrq.driver.Dialect())
	t1 := builder.Table(notificationrecord.Table)
	columns := nrq.fields
	if len(columns) == 0 {
		columns = notificationrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nrq.sql != nil {
		selector = nrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range nrq.modifiers {
		m(selector)
	}
	for _, p := range nrq.predicates {
		p(selector)
	}
	for _, p := range nrq.order {
		p(selector)
	}
	if offset := nrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (nrq *NotificationRecordQuery) ForUpdate(opts ...sql.LockOption) *NotificationRecordQuery {
	if nrq.driver.Dialect() == dialect.Postgres {
		nrq.Unique(false)
	}
	nrq.modifiers = append(nrq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return nrq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (nrq *NotificationRecordQuery) ForShare(opts ...sql.LockOption) *NotificationRecordQuery {
	if nrq.driver.Dialect() == dialect.Postgres {
		nrq.Unique(false)
	}
	nrq.modifiers = append(nrq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return nrq
}

// NotificationRecordGroupBy is the group-by builder for NotificationRecord entities.
type NotificationRecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nrgb *NotificationRecordGroupBy) Aggregate(fns ...AggregateFunc) *NotificationRecordGroupBy {
	nrgb.fns = append(nrgb.fns, fns...)
	return nrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (nrgb *NotificationRecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := nrgb.path(ctx)
	if err != nil {
		return err
	}
	nrgb.sql = query
	return nrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := nrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (nrgb *NotificationRecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(nrgb.fields) > 1 {
		return nil, errors.New("ent: NotificationRecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := nrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := nrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nrgb *NotificationRecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = fmt.Errorf("ent: NotificationRecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) StringX(ctx context.Context) string {
	v, err := nrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (nrgb *NotificationRecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(nrgb.fields) > 1 {
		return nil, errors.New("ent: NotificationRecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := nrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := nrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nrgb *NotificationRecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = fmt.Errorf("ent: NotificationRecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) IntX(ctx context.Context) int {
	v, err := nrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (nrgb *NotificationRecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(nrgb.fields) > 1 {
		return nil, errors.New("ent: NotificationRecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := nrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := nrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nrgb *NotificationRecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = fmt.Errorf("ent: NotificationRecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := nrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (nrgb *NotificationRecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(nrgb.fields) > 1 {
		return nil, errors.New("ent: NotificationRecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := nrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := nrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nrgb *NotificationRecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = fmt.Errorf("ent: NotificationRecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nrgb *NotificationRecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := nrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nrgb *NotificationRecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range nrgb.fields {
		if !notificationrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := nrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nrgb *NotificationRecordGroupBy) sqlQuery() *sql.Selector {
	selector := nrgb.sql.Select()
	aggregation := make([]string, 0, len(nrgb.fns))
	for _, fn := range nrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(nrgb.fields)+len(nrgb.fns))
		for _, f := range nrgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(nrgb.fields...)...)
}

// NotificationRecordSelect is the builder for selecting fields of NotificationRecord entities.
type NotificationRecordSelect struct {
	*NotificationRecordQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nrs *NotificationRecordSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nrs.prepareQuery(ctx); err != nil {
		return err
	}
	nrs.sql = nrs.NotificationRecordQuery.sqlQuery(ctx)
	return nrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nrs *NotificationRecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := nrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (nrs *NotificationRecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(nrs.fields) > 1 {
		return nil, errors.New("ent: NotificationRecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := nrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nrs *NotificationRecordSelect) StringsX(ctx context.Context) []string {
	v, err := nrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (nrs *NotificationRecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = fmt.Errorf("ent: NotificationRecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nrs *NotificationRecordSelect) StringX(ctx context.Context) string {
	v, err := nrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (nrs *NotificationRecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(nrs.fields) > 1 {
		return nil, errors.New("ent: NotificationRecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := nrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nrs *NotificationRecordSelect) IntsX(ctx context.Context) []int {
	v, err := nrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (nrs *NotificationRecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = fmt.Errorf("ent: NotificationRecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nrs *NotificationRecordSelect) IntX(ctx context.Context) int {
	v, err := nrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (nrs *NotificationRecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(nrs.fields) > 1 {
		return nil, errors.New("ent: NotificationRecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := nrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nrs *NotificationRecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := nrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (nrs *NotificationRecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = fmt.Errorf("ent: NotificationRecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nrs *NotificationRecordSelect) Float64X(ctx context.Context) float64 {
	v, err := nrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (nrs *NotificationRecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(nrs.fields) > 1 {
		return nil, errors.New("ent: NotificationRecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := nrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nrs *NotificationRecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := nrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (nrs *NotificationRecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{notificationrecord.Label}
	default:
		err = fmt.Errorf("ent: NotificationRecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nrs *NotificationRecordSelect) BoolX(ctx context.Context) bool {
	v, err := nrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nrs *NotificationRecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nrs.sql.Query()
	if err := nrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
