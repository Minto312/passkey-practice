// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Minto312/passkey-practice/backend/ent/authhistory"
	"github.com/Minto312/passkey-practice/backend/ent/predicate"
	"github.com/Minto312/passkey-practice/backend/ent/user"
	"github.com/google/uuid"
)

// AuthHistoryQuery is the builder for querying AuthHistory entities.
type AuthHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []authhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.AuthHistory
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthHistoryQuery builder.
func (ahq *AuthHistoryQuery) Where(ps ...predicate.AuthHistory) *AuthHistoryQuery {
	ahq.predicates = append(ahq.predicates, ps...)
	return ahq
}

// Limit the number of records to be returned by this query.
func (ahq *AuthHistoryQuery) Limit(limit int) *AuthHistoryQuery {
	ahq.ctx.Limit = &limit
	return ahq
}

// Offset to start from.
func (ahq *AuthHistoryQuery) Offset(offset int) *AuthHistoryQuery {
	ahq.ctx.Offset = &offset
	return ahq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ahq *AuthHistoryQuery) Unique(unique bool) *AuthHistoryQuery {
	ahq.ctx.Unique = &unique
	return ahq
}

// Order specifies how the records should be ordered.
func (ahq *AuthHistoryQuery) Order(o ...authhistory.OrderOption) *AuthHistoryQuery {
	ahq.order = append(ahq.order, o...)
	return ahq
}

// QueryUser chains the current query on the "user" edge.
func (ahq *AuthHistoryQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ahq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ahq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ahq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authhistory.Table, authhistory.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, authhistory.UserTable, authhistory.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ahq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthHistory entity from the query.
// Returns a *NotFoundError when no AuthHistory was found.
func (ahq *AuthHistoryQuery) First(ctx context.Context) (*AuthHistory, error) {
	nodes, err := ahq.Limit(1).All(setContextOp(ctx, ahq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ahq *AuthHistoryQuery) FirstX(ctx context.Context) *AuthHistory {
	node, err := ahq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthHistory ID from the query.
// Returns a *NotFoundError when no AuthHistory ID was found.
func (ahq *AuthHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ahq.Limit(1).IDs(setContextOp(ctx, ahq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ahq *AuthHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ahq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AuthHistory entity is found.
// Returns a *NotFoundError when no AuthHistory entities are found.
func (ahq *AuthHistoryQuery) Only(ctx context.Context) (*AuthHistory, error) {
	nodes, err := ahq.Limit(2).All(setContextOp(ctx, ahq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authhistory.Label}
	default:
		return nil, &NotSingularError{authhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ahq *AuthHistoryQuery) OnlyX(ctx context.Context) *AuthHistory {
	node, err := ahq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthHistory ID in the query.
// Returns a *NotSingularError when more than one AuthHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ahq *AuthHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ahq.Limit(2).IDs(setContextOp(ctx, ahq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = &NotSingularError{authhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ahq *AuthHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ahq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthHistories.
func (ahq *AuthHistoryQuery) All(ctx context.Context) ([]*AuthHistory, error) {
	ctx = setContextOp(ctx, ahq.ctx, ent.OpQueryAll)
	if err := ahq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AuthHistory, *AuthHistoryQuery]()
	return withInterceptors[[]*AuthHistory](ctx, ahq, qr, ahq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ahq *AuthHistoryQuery) AllX(ctx context.Context) []*AuthHistory {
	nodes, err := ahq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthHistory IDs.
func (ahq *AuthHistoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ahq.ctx.Unique == nil && ahq.path != nil {
		ahq.Unique(true)
	}
	ctx = setContextOp(ctx, ahq.ctx, ent.OpQueryIDs)
	if err = ahq.Select(authhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ahq *AuthHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ahq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ahq *AuthHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ahq.ctx, ent.OpQueryCount)
	if err := ahq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ahq, querierCount[*AuthHistoryQuery](), ahq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ahq *AuthHistoryQuery) CountX(ctx context.Context) int {
	count, err := ahq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ahq *AuthHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ahq.ctx, ent.OpQueryExist)
	switch _, err := ahq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ahq *AuthHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ahq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ahq *AuthHistoryQuery) Clone() *AuthHistoryQuery {
	if ahq == nil {
		return nil
	}
	return &AuthHistoryQuery{
		config:     ahq.config,
		ctx:        ahq.ctx.Clone(),
		order:      append([]authhistory.OrderOption{}, ahq.order...),
		inters:     append([]Interceptor{}, ahq.inters...),
		predicates: append([]predicate.AuthHistory{}, ahq.predicates...),
		withUser:   ahq.withUser.Clone(),
		// clone intermediate query.
		sql:  ahq.sql.Clone(),
		path: ahq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ahq *AuthHistoryQuery) WithUser(opts ...func(*UserQuery)) *AuthHistoryQuery {
	query := (&UserClient{config: ahq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ahq.withUser = query
	return ahq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Method string `json:"method,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthHistory.Query().
//		GroupBy(authhistory.FieldMethod).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ahq *AuthHistoryQuery) GroupBy(field string, fields ...string) *AuthHistoryGroupBy {
	ahq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AuthHistoryGroupBy{build: ahq}
	grbuild.flds = &ahq.ctx.Fields
	grbuild.label = authhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Method string `json:"method,omitempty"`
//	}
//
//	client.AuthHistory.Query().
//		Select(authhistory.FieldMethod).
//		Scan(ctx, &v)
func (ahq *AuthHistoryQuery) Select(fields ...string) *AuthHistorySelect {
	ahq.ctx.Fields = append(ahq.ctx.Fields, fields...)
	sbuild := &AuthHistorySelect{AuthHistoryQuery: ahq}
	sbuild.label = authhistory.Label
	sbuild.flds, sbuild.scan = &ahq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AuthHistorySelect configured with the given aggregations.
func (ahq *AuthHistoryQuery) Aggregate(fns ...AggregateFunc) *AuthHistorySelect {
	return ahq.Select().Aggregate(fns...)
}

func (ahq *AuthHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ahq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ahq); err != nil {
				return err
			}
		}
	}
	for _, f := range ahq.ctx.Fields {
		if !authhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ahq.path != nil {
		prev, err := ahq.path(ctx)
		if err != nil {
			return err
		}
		ahq.sql = prev
	}
	return nil
}

func (ahq *AuthHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AuthHistory, error) {
	var (
		nodes       = []*AuthHistory{}
		withFKs     = ahq.withFKs
		_spec       = ahq.querySpec()
		loadedTypes = [1]bool{
			ahq.withUser != nil,
		}
	)
	if ahq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, authhistory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AuthHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AuthHistory{config: ahq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ahq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ahq.withUser; query != nil {
		if err := ahq.loadUser(ctx, query, nodes, nil,
			func(n *AuthHistory, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ahq *AuthHistoryQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*AuthHistory, init func(*AuthHistory), assign func(*AuthHistory, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AuthHistory)
	for i := range nodes {
		if nodes[i].user_auth_histories == nil {
			continue
		}
		fk := *nodes[i].user_auth_histories
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_auth_histories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ahq *AuthHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ahq.querySpec()
	_spec.Node.Columns = ahq.ctx.Fields
	if len(ahq.ctx.Fields) > 0 {
		_spec.Unique = ahq.ctx.Unique != nil && *ahq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ahq.driver, _spec)
}

func (ahq *AuthHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(authhistory.Table, authhistory.Columns, sqlgraph.NewFieldSpec(authhistory.FieldID, field.TypeUUID))
	_spec.From = ahq.sql
	if unique := ahq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ahq.path != nil {
		_spec.Unique = true
	}
	if fields := ahq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authhistory.FieldID)
		for i := range fields {
			if fields[i] != authhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ahq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ahq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ahq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ahq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ahq *AuthHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ahq.driver.Dialect())
	t1 := builder.Table(authhistory.Table)
	columns := ahq.ctx.Fields
	if len(columns) == 0 {
		columns = authhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ahq.sql != nil {
		selector = ahq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ahq.ctx.Unique != nil && *ahq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ahq.predicates {
		p(selector)
	}
	for _, p := range ahq.order {
		p(selector)
	}
	if offset := ahq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ahq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthHistoryGroupBy is the group-by builder for AuthHistory entities.
type AuthHistoryGroupBy struct {
	selector
	build *AuthHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ahgb *AuthHistoryGroupBy) Aggregate(fns ...AggregateFunc) *AuthHistoryGroupBy {
	ahgb.fns = append(ahgb.fns, fns...)
	return ahgb
}

// Scan applies the selector query and scans the result into the given value.
func (ahgb *AuthHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ahgb.build.ctx, ent.OpQueryGroupBy)
	if err := ahgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthHistoryQuery, *AuthHistoryGroupBy](ctx, ahgb.build, ahgb, ahgb.build.inters, v)
}

func (ahgb *AuthHistoryGroupBy) sqlScan(ctx context.Context, root *AuthHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ahgb.fns))
	for _, fn := range ahgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ahgb.flds)+len(ahgb.fns))
		for _, f := range *ahgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ahgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ahgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AuthHistorySelect is the builder for selecting fields of AuthHistory entities.
type AuthHistorySelect struct {
	*AuthHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ahs *AuthHistorySelect) Aggregate(fns ...AggregateFunc) *AuthHistorySelect {
	ahs.fns = append(ahs.fns, fns...)
	return ahs
}

// Scan applies the selector query and scans the result into the given value.
func (ahs *AuthHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ahs.ctx, ent.OpQuerySelect)
	if err := ahs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthHistoryQuery, *AuthHistorySelect](ctx, ahs.AuthHistoryQuery, ahs, ahs.inters, v)
}

func (ahs *AuthHistorySelect) sqlScan(ctx context.Context, root *AuthHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ahs.fns))
	for _, fn := range ahs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ahs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ahs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
