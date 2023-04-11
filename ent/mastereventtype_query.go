// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"outgrow/ent/mastereventtype"
	"outgrow/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MasterEventTypeQuery is the builder for querying MasterEventType entities.
type MasterEventTypeQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.MasterEventType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MasterEventTypeQuery builder.
func (metq *MasterEventTypeQuery) Where(ps ...predicate.MasterEventType) *MasterEventTypeQuery {
	metq.predicates = append(metq.predicates, ps...)
	return metq
}

// Limit the number of records to be returned by this query.
func (metq *MasterEventTypeQuery) Limit(limit int) *MasterEventTypeQuery {
	metq.ctx.Limit = &limit
	return metq
}

// Offset to start from.
func (metq *MasterEventTypeQuery) Offset(offset int) *MasterEventTypeQuery {
	metq.ctx.Offset = &offset
	return metq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (metq *MasterEventTypeQuery) Unique(unique bool) *MasterEventTypeQuery {
	metq.ctx.Unique = &unique
	return metq
}

// Order specifies how the records should be ordered.
func (metq *MasterEventTypeQuery) Order(o ...OrderFunc) *MasterEventTypeQuery {
	metq.order = append(metq.order, o...)
	return metq
}

// First returns the first MasterEventType entity from the query.
// Returns a *NotFoundError when no MasterEventType was found.
func (metq *MasterEventTypeQuery) First(ctx context.Context) (*MasterEventType, error) {
	nodes, err := metq.Limit(1).All(setContextOp(ctx, metq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mastereventtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (metq *MasterEventTypeQuery) FirstX(ctx context.Context) *MasterEventType {
	node, err := metq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MasterEventType ID from the query.
// Returns a *NotFoundError when no MasterEventType ID was found.
func (metq *MasterEventTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = metq.Limit(1).IDs(setContextOp(ctx, metq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mastereventtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (metq *MasterEventTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := metq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MasterEventType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MasterEventType entity is found.
// Returns a *NotFoundError when no MasterEventType entities are found.
func (metq *MasterEventTypeQuery) Only(ctx context.Context) (*MasterEventType, error) {
	nodes, err := metq.Limit(2).All(setContextOp(ctx, metq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mastereventtype.Label}
	default:
		return nil, &NotSingularError{mastereventtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (metq *MasterEventTypeQuery) OnlyX(ctx context.Context) *MasterEventType {
	node, err := metq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MasterEventType ID in the query.
// Returns a *NotSingularError when more than one MasterEventType ID is found.
// Returns a *NotFoundError when no entities are found.
func (metq *MasterEventTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = metq.Limit(2).IDs(setContextOp(ctx, metq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mastereventtype.Label}
	default:
		err = &NotSingularError{mastereventtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (metq *MasterEventTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := metq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MasterEventTypes.
func (metq *MasterEventTypeQuery) All(ctx context.Context) ([]*MasterEventType, error) {
	ctx = setContextOp(ctx, metq.ctx, "All")
	if err := metq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MasterEventType, *MasterEventTypeQuery]()
	return withInterceptors[[]*MasterEventType](ctx, metq, qr, metq.inters)
}

// AllX is like All, but panics if an error occurs.
func (metq *MasterEventTypeQuery) AllX(ctx context.Context) []*MasterEventType {
	nodes, err := metq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MasterEventType IDs.
func (metq *MasterEventTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if metq.ctx.Unique == nil && metq.path != nil {
		metq.Unique(true)
	}
	ctx = setContextOp(ctx, metq.ctx, "IDs")
	if err = metq.Select(mastereventtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (metq *MasterEventTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := metq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (metq *MasterEventTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, metq.ctx, "Count")
	if err := metq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, metq, querierCount[*MasterEventTypeQuery](), metq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (metq *MasterEventTypeQuery) CountX(ctx context.Context) int {
	count, err := metq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (metq *MasterEventTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, metq.ctx, "Exist")
	switch _, err := metq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (metq *MasterEventTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := metq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MasterEventTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (metq *MasterEventTypeQuery) Clone() *MasterEventTypeQuery {
	if metq == nil {
		return nil
	}
	return &MasterEventTypeQuery{
		config:     metq.config,
		ctx:        metq.ctx.Clone(),
		order:      append([]OrderFunc{}, metq.order...),
		inters:     append([]Interceptor{}, metq.inters...),
		predicates: append([]predicate.MasterEventType{}, metq.predicates...),
		// clone intermediate query.
		sql:  metq.sql.Clone(),
		path: metq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MasterEventType.Query().
//		GroupBy(mastereventtype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (metq *MasterEventTypeQuery) GroupBy(field string, fields ...string) *MasterEventTypeGroupBy {
	metq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MasterEventTypeGroupBy{build: metq}
	grbuild.flds = &metq.ctx.Fields
	grbuild.label = mastereventtype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.MasterEventType.Query().
//		Select(mastereventtype.FieldName).
//		Scan(ctx, &v)
func (metq *MasterEventTypeQuery) Select(fields ...string) *MasterEventTypeSelect {
	metq.ctx.Fields = append(metq.ctx.Fields, fields...)
	sbuild := &MasterEventTypeSelect{MasterEventTypeQuery: metq}
	sbuild.label = mastereventtype.Label
	sbuild.flds, sbuild.scan = &metq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MasterEventTypeSelect configured with the given aggregations.
func (metq *MasterEventTypeQuery) Aggregate(fns ...AggregateFunc) *MasterEventTypeSelect {
	return metq.Select().Aggregate(fns...)
}

func (metq *MasterEventTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range metq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, metq); err != nil {
				return err
			}
		}
	}
	for _, f := range metq.ctx.Fields {
		if !mastereventtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if metq.path != nil {
		prev, err := metq.path(ctx)
		if err != nil {
			return err
		}
		metq.sql = prev
	}
	return nil
}

func (metq *MasterEventTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MasterEventType, error) {
	var (
		nodes = []*MasterEventType{}
		_spec = metq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MasterEventType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MasterEventType{config: metq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, metq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (metq *MasterEventTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := metq.querySpec()
	_spec.Node.Columns = metq.ctx.Fields
	if len(metq.ctx.Fields) > 0 {
		_spec.Unique = metq.ctx.Unique != nil && *metq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, metq.driver, _spec)
}

func (metq *MasterEventTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(mastereventtype.Table, mastereventtype.Columns, sqlgraph.NewFieldSpec(mastereventtype.FieldID, field.TypeInt))
	_spec.From = metq.sql
	if unique := metq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if metq.path != nil {
		_spec.Unique = true
	}
	if fields := metq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mastereventtype.FieldID)
		for i := range fields {
			if fields[i] != mastereventtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := metq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := metq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := metq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := metq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (metq *MasterEventTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(metq.driver.Dialect())
	t1 := builder.Table(mastereventtype.Table)
	columns := metq.ctx.Fields
	if len(columns) == 0 {
		columns = mastereventtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if metq.sql != nil {
		selector = metq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if metq.ctx.Unique != nil && *metq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range metq.predicates {
		p(selector)
	}
	for _, p := range metq.order {
		p(selector)
	}
	if offset := metq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := metq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MasterEventTypeGroupBy is the group-by builder for MasterEventType entities.
type MasterEventTypeGroupBy struct {
	selector
	build *MasterEventTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (metgb *MasterEventTypeGroupBy) Aggregate(fns ...AggregateFunc) *MasterEventTypeGroupBy {
	metgb.fns = append(metgb.fns, fns...)
	return metgb
}

// Scan applies the selector query and scans the result into the given value.
func (metgb *MasterEventTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, metgb.build.ctx, "GroupBy")
	if err := metgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MasterEventTypeQuery, *MasterEventTypeGroupBy](ctx, metgb.build, metgb, metgb.build.inters, v)
}

func (metgb *MasterEventTypeGroupBy) sqlScan(ctx context.Context, root *MasterEventTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(metgb.fns))
	for _, fn := range metgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*metgb.flds)+len(metgb.fns))
		for _, f := range *metgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*metgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := metgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MasterEventTypeSelect is the builder for selecting fields of MasterEventType entities.
type MasterEventTypeSelect struct {
	*MasterEventTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mets *MasterEventTypeSelect) Aggregate(fns ...AggregateFunc) *MasterEventTypeSelect {
	mets.fns = append(mets.fns, fns...)
	return mets
}

// Scan applies the selector query and scans the result into the given value.
func (mets *MasterEventTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mets.ctx, "Select")
	if err := mets.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MasterEventTypeQuery, *MasterEventTypeSelect](ctx, mets.MasterEventTypeQuery, mets, mets.inters, v)
}

func (mets *MasterEventTypeSelect) sqlScan(ctx context.Context, root *MasterEventTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mets.fns))
	for _, fn := range mets.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mets.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
