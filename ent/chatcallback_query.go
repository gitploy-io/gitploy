// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/chatuser"
	"github.com/hanjunlee/gitploy/ent/predicate"
	"github.com/hanjunlee/gitploy/ent/repo"
)

// ChatCallbackQuery is the builder for querying ChatCallback entities.
type ChatCallbackQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ChatCallback
	// eager-loading edges.
	withChatUser *ChatUserQuery
	withRepo     *RepoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChatCallbackQuery builder.
func (ccq *ChatCallbackQuery) Where(ps ...predicate.ChatCallback) *ChatCallbackQuery {
	ccq.predicates = append(ccq.predicates, ps...)
	return ccq
}

// Limit adds a limit step to the query.
func (ccq *ChatCallbackQuery) Limit(limit int) *ChatCallbackQuery {
	ccq.limit = &limit
	return ccq
}

// Offset adds an offset step to the query.
func (ccq *ChatCallbackQuery) Offset(offset int) *ChatCallbackQuery {
	ccq.offset = &offset
	return ccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccq *ChatCallbackQuery) Unique(unique bool) *ChatCallbackQuery {
	ccq.unique = &unique
	return ccq
}

// Order adds an order step to the query.
func (ccq *ChatCallbackQuery) Order(o ...OrderFunc) *ChatCallbackQuery {
	ccq.order = append(ccq.order, o...)
	return ccq
}

// QueryChatUser chains the current query on the "chat_user" edge.
func (ccq *ChatCallbackQuery) QueryChatUser() *ChatUserQuery {
	query := &ChatUserQuery{config: ccq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chatcallback.Table, chatcallback.FieldID, selector),
			sqlgraph.To(chatuser.Table, chatuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chatcallback.ChatUserTable, chatcallback.ChatUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRepo chains the current query on the "repo" edge.
func (ccq *ChatCallbackQuery) QueryRepo() *RepoQuery {
	query := &RepoQuery{config: ccq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chatcallback.Table, chatcallback.FieldID, selector),
			sqlgraph.To(repo.Table, repo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chatcallback.RepoTable, chatcallback.RepoColumn),
		)
		fromU = sqlgraph.SetNeighbors(ccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ChatCallback entity from the query.
// Returns a *NotFoundError when no ChatCallback was found.
func (ccq *ChatCallbackQuery) First(ctx context.Context) (*ChatCallback, error) {
	nodes, err := ccq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chatcallback.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccq *ChatCallbackQuery) FirstX(ctx context.Context) *ChatCallback {
	node, err := ccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChatCallback ID from the query.
// Returns a *NotFoundError when no ChatCallback ID was found.
func (ccq *ChatCallbackQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ccq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chatcallback.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccq *ChatCallbackQuery) FirstIDX(ctx context.Context) int {
	id, err := ccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChatCallback entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ChatCallback entity is not found.
// Returns a *NotFoundError when no ChatCallback entities are found.
func (ccq *ChatCallbackQuery) Only(ctx context.Context) (*ChatCallback, error) {
	nodes, err := ccq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chatcallback.Label}
	default:
		return nil, &NotSingularError{chatcallback.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccq *ChatCallbackQuery) OnlyX(ctx context.Context) *ChatCallback {
	node, err := ccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChatCallback ID in the query.
// Returns a *NotSingularError when exactly one ChatCallback ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ccq *ChatCallbackQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ccq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = &NotSingularError{chatcallback.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccq *ChatCallbackQuery) OnlyIDX(ctx context.Context) int {
	id, err := ccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChatCallbacks.
func (ccq *ChatCallbackQuery) All(ctx context.Context) ([]*ChatCallback, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ccq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ccq *ChatCallbackQuery) AllX(ctx context.Context) []*ChatCallback {
	nodes, err := ccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChatCallback IDs.
func (ccq *ChatCallbackQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ccq.Select(chatcallback.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccq *ChatCallbackQuery) IDsX(ctx context.Context) []int {
	ids, err := ccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccq *ChatCallbackQuery) Count(ctx context.Context) (int, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ccq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ccq *ChatCallbackQuery) CountX(ctx context.Context) int {
	count, err := ccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccq *ChatCallbackQuery) Exist(ctx context.Context) (bool, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ccq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ccq *ChatCallbackQuery) ExistX(ctx context.Context) bool {
	exist, err := ccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChatCallbackQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccq *ChatCallbackQuery) Clone() *ChatCallbackQuery {
	if ccq == nil {
		return nil
	}
	return &ChatCallbackQuery{
		config:       ccq.config,
		limit:        ccq.limit,
		offset:       ccq.offset,
		order:        append([]OrderFunc{}, ccq.order...),
		predicates:   append([]predicate.ChatCallback{}, ccq.predicates...),
		withChatUser: ccq.withChatUser.Clone(),
		withRepo:     ccq.withRepo.Clone(),
		// clone intermediate query.
		sql:  ccq.sql.Clone(),
		path: ccq.path,
	}
}

// WithChatUser tells the query-builder to eager-load the nodes that are connected to
// the "chat_user" edge. The optional arguments are used to configure the query builder of the edge.
func (ccq *ChatCallbackQuery) WithChatUser(opts ...func(*ChatUserQuery)) *ChatCallbackQuery {
	query := &ChatUserQuery{config: ccq.config}
	for _, opt := range opts {
		opt(query)
	}
	ccq.withChatUser = query
	return ccq
}

// WithRepo tells the query-builder to eager-load the nodes that are connected to
// the "repo" edge. The optional arguments are used to configure the query builder of the edge.
func (ccq *ChatCallbackQuery) WithRepo(opts ...func(*RepoQuery)) *ChatCallbackQuery {
	query := &RepoQuery{config: ccq.config}
	for _, opt := range opts {
		opt(query)
	}
	ccq.withRepo = query
	return ccq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Hash string `json:"hash"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ChatCallback.Query().
//		GroupBy(chatcallback.FieldHash).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ccq *ChatCallbackQuery) GroupBy(field string, fields ...string) *ChatCallbackGroupBy {
	group := &ChatCallbackGroupBy{config: ccq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ccq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Hash string `json:"hash"`
//	}
//
//	client.ChatCallback.Query().
//		Select(chatcallback.FieldHash).
//		Scan(ctx, &v)
//
func (ccq *ChatCallbackQuery) Select(fields ...string) *ChatCallbackSelect {
	ccq.fields = append(ccq.fields, fields...)
	return &ChatCallbackSelect{ChatCallbackQuery: ccq}
}

func (ccq *ChatCallbackQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ccq.fields {
		if !chatcallback.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccq.path != nil {
		prev, err := ccq.path(ctx)
		if err != nil {
			return err
		}
		ccq.sql = prev
	}
	return nil
}

func (ccq *ChatCallbackQuery) sqlAll(ctx context.Context) ([]*ChatCallback, error) {
	var (
		nodes       = []*ChatCallback{}
		_spec       = ccq.querySpec()
		loadedTypes = [2]bool{
			ccq.withChatUser != nil,
			ccq.withRepo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ChatCallback{config: ccq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ccq.withChatUser; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*ChatCallback)
		for i := range nodes {
			fk := nodes[i].ChatUserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(chatuser.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "chat_user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChatUser = n
			}
		}
	}

	if query := ccq.withRepo; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*ChatCallback)
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

func (ccq *ChatCallbackQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccq.querySpec()
	return sqlgraph.CountNodes(ctx, ccq.driver, _spec)
}

func (ccq *ChatCallbackQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ccq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ccq *ChatCallbackQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chatcallback.Table,
			Columns: chatcallback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chatcallback.FieldID,
			},
		},
		From:   ccq.sql,
		Unique: true,
	}
	if unique := ccq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ccq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatcallback.FieldID)
		for i := range fields {
			if fields[i] != chatcallback.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccq *ChatCallbackQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccq.driver.Dialect())
	t1 := builder.Table(chatcallback.Table)
	columns := ccq.fields
	if len(columns) == 0 {
		columns = chatcallback.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccq.sql != nil {
		selector = ccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ccq.predicates {
		p(selector)
	}
	for _, p := range ccq.order {
		p(selector)
	}
	if offset := ccq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChatCallbackGroupBy is the group-by builder for ChatCallback entities.
type ChatCallbackGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccgb *ChatCallbackGroupBy) Aggregate(fns ...AggregateFunc) *ChatCallbackGroupBy {
	ccgb.fns = append(ccgb.fns, fns...)
	return ccgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ccgb *ChatCallbackGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ccgb.path(ctx)
	if err != nil {
		return err
	}
	ccgb.sql = query
	return ccgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ccgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccgb *ChatCallbackGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ccgb.fields) > 1 {
		return nil, errors.New("ent: ChatCallbackGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ccgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) StringsX(ctx context.Context) []string {
	v, err := ccgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccgb *ChatCallbackGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ccgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = fmt.Errorf("ent: ChatCallbackGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) StringX(ctx context.Context) string {
	v, err := ccgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccgb *ChatCallbackGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ccgb.fields) > 1 {
		return nil, errors.New("ent: ChatCallbackGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ccgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) IntsX(ctx context.Context) []int {
	v, err := ccgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccgb *ChatCallbackGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ccgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = fmt.Errorf("ent: ChatCallbackGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) IntX(ctx context.Context) int {
	v, err := ccgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccgb *ChatCallbackGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ccgb.fields) > 1 {
		return nil, errors.New("ent: ChatCallbackGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ccgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ccgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccgb *ChatCallbackGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ccgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = fmt.Errorf("ent: ChatCallbackGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ccgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccgb *ChatCallbackGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ccgb.fields) > 1 {
		return nil, errors.New("ent: ChatCallbackGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ccgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ccgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccgb *ChatCallbackGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ccgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = fmt.Errorf("ent: ChatCallbackGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ccgb *ChatCallbackGroupBy) BoolX(ctx context.Context) bool {
	v, err := ccgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ccgb *ChatCallbackGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ccgb.fields {
		if !chatcallback.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ccgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ccgb *ChatCallbackGroupBy) sqlQuery() *sql.Selector {
	selector := ccgb.sql.Select()
	aggregation := make([]string, 0, len(ccgb.fns))
	for _, fn := range ccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ccgb.fields)+len(ccgb.fns))
		for _, f := range ccgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ccgb.fields...)...)
}

// ChatCallbackSelect is the builder for selecting fields of ChatCallback entities.
type ChatCallbackSelect struct {
	*ChatCallbackQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ccs *ChatCallbackSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ccs.prepareQuery(ctx); err != nil {
		return err
	}
	ccs.sql = ccs.ChatCallbackQuery.sqlQuery(ctx)
	return ccs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ccs *ChatCallbackSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ccs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ccs *ChatCallbackSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ccs.fields) > 1 {
		return nil, errors.New("ent: ChatCallbackSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ccs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ccs *ChatCallbackSelect) StringsX(ctx context.Context) []string {
	v, err := ccs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ccs *ChatCallbackSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ccs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = fmt.Errorf("ent: ChatCallbackSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ccs *ChatCallbackSelect) StringX(ctx context.Context) string {
	v, err := ccs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ccs *ChatCallbackSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ccs.fields) > 1 {
		return nil, errors.New("ent: ChatCallbackSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ccs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ccs *ChatCallbackSelect) IntsX(ctx context.Context) []int {
	v, err := ccs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ccs *ChatCallbackSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ccs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = fmt.Errorf("ent: ChatCallbackSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ccs *ChatCallbackSelect) IntX(ctx context.Context) int {
	v, err := ccs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ccs *ChatCallbackSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ccs.fields) > 1 {
		return nil, errors.New("ent: ChatCallbackSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ccs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ccs *ChatCallbackSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ccs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ccs *ChatCallbackSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ccs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = fmt.Errorf("ent: ChatCallbackSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ccs *ChatCallbackSelect) Float64X(ctx context.Context) float64 {
	v, err := ccs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ccs *ChatCallbackSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ccs.fields) > 1 {
		return nil, errors.New("ent: ChatCallbackSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ccs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ccs *ChatCallbackSelect) BoolsX(ctx context.Context) []bool {
	v, err := ccs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ccs *ChatCallbackSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ccs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatcallback.Label}
	default:
		err = fmt.Errorf("ent: ChatCallbackSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ccs *ChatCallbackSelect) BoolX(ctx context.Context) bool {
	v, err := ccs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ccs *ChatCallbackSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ccs.sql.Query()
	if err := ccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
