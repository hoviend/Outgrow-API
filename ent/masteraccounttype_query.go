// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"outgrow/ent/masteraccountcategory"
	"outgrow/ent/masteraccounttype"
	"outgrow/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MasterAccountTypeQuery is the builder for querying MasterAccountType entities.
type MasterAccountTypeQuery struct {
	config
	ctx            *QueryContext
	order          []OrderFunc
	inters         []Interceptor
	predicates     []predicate.MasterAccountType
	withCategories *MasterAccountCategoryQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MasterAccountTypeQuery builder.
func (matq *MasterAccountTypeQuery) Where(ps ...predicate.MasterAccountType) *MasterAccountTypeQuery {
	matq.predicates = append(matq.predicates, ps...)
	return matq
}

// Limit the number of records to be returned by this query.
func (matq *MasterAccountTypeQuery) Limit(limit int) *MasterAccountTypeQuery {
	matq.ctx.Limit = &limit
	return matq
}

// Offset to start from.
func (matq *MasterAccountTypeQuery) Offset(offset int) *MasterAccountTypeQuery {
	matq.ctx.Offset = &offset
	return matq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (matq *MasterAccountTypeQuery) Unique(unique bool) *MasterAccountTypeQuery {
	matq.ctx.Unique = &unique
	return matq
}

// Order specifies how the records should be ordered.
func (matq *MasterAccountTypeQuery) Order(o ...OrderFunc) *MasterAccountTypeQuery {
	matq.order = append(matq.order, o...)
	return matq
}

// QueryCategories chains the current query on the "categories" edge.
func (matq *MasterAccountTypeQuery) QueryCategories() *MasterAccountCategoryQuery {
	query := (&MasterAccountCategoryClient{config: matq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := matq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := matq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(masteraccounttype.Table, masteraccounttype.FieldID, selector),
			sqlgraph.To(masteraccountcategory.Table, masteraccountcategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, masteraccounttype.CategoriesTable, masteraccounttype.CategoriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(matq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MasterAccountType entity from the query.
// Returns a *NotFoundError when no MasterAccountType was found.
func (matq *MasterAccountTypeQuery) First(ctx context.Context) (*MasterAccountType, error) {
	nodes, err := matq.Limit(1).All(setContextOp(ctx, matq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{masteraccounttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (matq *MasterAccountTypeQuery) FirstX(ctx context.Context) *MasterAccountType {
	node, err := matq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MasterAccountType ID from the query.
// Returns a *NotFoundError when no MasterAccountType ID was found.
func (matq *MasterAccountTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = matq.Limit(1).IDs(setContextOp(ctx, matq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{masteraccounttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (matq *MasterAccountTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := matq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MasterAccountType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MasterAccountType entity is found.
// Returns a *NotFoundError when no MasterAccountType entities are found.
func (matq *MasterAccountTypeQuery) Only(ctx context.Context) (*MasterAccountType, error) {
	nodes, err := matq.Limit(2).All(setContextOp(ctx, matq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{masteraccounttype.Label}
	default:
		return nil, &NotSingularError{masteraccounttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (matq *MasterAccountTypeQuery) OnlyX(ctx context.Context) *MasterAccountType {
	node, err := matq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MasterAccountType ID in the query.
// Returns a *NotSingularError when more than one MasterAccountType ID is found.
// Returns a *NotFoundError when no entities are found.
func (matq *MasterAccountTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = matq.Limit(2).IDs(setContextOp(ctx, matq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{masteraccounttype.Label}
	default:
		err = &NotSingularError{masteraccounttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (matq *MasterAccountTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := matq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MasterAccountTypes.
func (matq *MasterAccountTypeQuery) All(ctx context.Context) ([]*MasterAccountType, error) {
	ctx = setContextOp(ctx, matq.ctx, "All")
	if err := matq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MasterAccountType, *MasterAccountTypeQuery]()
	return withInterceptors[[]*MasterAccountType](ctx, matq, qr, matq.inters)
}

// AllX is like All, but panics if an error occurs.
func (matq *MasterAccountTypeQuery) AllX(ctx context.Context) []*MasterAccountType {
	nodes, err := matq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MasterAccountType IDs.
func (matq *MasterAccountTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if matq.ctx.Unique == nil && matq.path != nil {
		matq.Unique(true)
	}
	ctx = setContextOp(ctx, matq.ctx, "IDs")
	if err = matq.Select(masteraccounttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (matq *MasterAccountTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := matq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (matq *MasterAccountTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, matq.ctx, "Count")
	if err := matq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, matq, querierCount[*MasterAccountTypeQuery](), matq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (matq *MasterAccountTypeQuery) CountX(ctx context.Context) int {
	count, err := matq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (matq *MasterAccountTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, matq.ctx, "Exist")
	switch _, err := matq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (matq *MasterAccountTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := matq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MasterAccountTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (matq *MasterAccountTypeQuery) Clone() *MasterAccountTypeQuery {
	if matq == nil {
		return nil
	}
	return &MasterAccountTypeQuery{
		config:         matq.config,
		ctx:            matq.ctx.Clone(),
		order:          append([]OrderFunc{}, matq.order...),
		inters:         append([]Interceptor{}, matq.inters...),
		predicates:     append([]predicate.MasterAccountType{}, matq.predicates...),
		withCategories: matq.withCategories.Clone(),
		// clone intermediate query.
		sql:  matq.sql.Clone(),
		path: matq.path,
	}
}

// WithCategories tells the query-builder to eager-load the nodes that are connected to
// the "categories" edge. The optional arguments are used to configure the query builder of the edge.
func (matq *MasterAccountTypeQuery) WithCategories(opts ...func(*MasterAccountCategoryQuery)) *MasterAccountTypeQuery {
	query := (&MasterAccountCategoryClient{config: matq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	matq.withCategories = query
	return matq
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
//	client.MasterAccountType.Query().
//		GroupBy(masteraccounttype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (matq *MasterAccountTypeQuery) GroupBy(field string, fields ...string) *MasterAccountTypeGroupBy {
	matq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MasterAccountTypeGroupBy{build: matq}
	grbuild.flds = &matq.ctx.Fields
	grbuild.label = masteraccounttype.Label
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
//	client.MasterAccountType.Query().
//		Select(masteraccounttype.FieldName).
//		Scan(ctx, &v)
func (matq *MasterAccountTypeQuery) Select(fields ...string) *MasterAccountTypeSelect {
	matq.ctx.Fields = append(matq.ctx.Fields, fields...)
	sbuild := &MasterAccountTypeSelect{MasterAccountTypeQuery: matq}
	sbuild.label = masteraccounttype.Label
	sbuild.flds, sbuild.scan = &matq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MasterAccountTypeSelect configured with the given aggregations.
func (matq *MasterAccountTypeQuery) Aggregate(fns ...AggregateFunc) *MasterAccountTypeSelect {
	return matq.Select().Aggregate(fns...)
}

func (matq *MasterAccountTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range matq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, matq); err != nil {
				return err
			}
		}
	}
	for _, f := range matq.ctx.Fields {
		if !masteraccounttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if matq.path != nil {
		prev, err := matq.path(ctx)
		if err != nil {
			return err
		}
		matq.sql = prev
	}
	return nil
}

func (matq *MasterAccountTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MasterAccountType, error) {
	var (
		nodes       = []*MasterAccountType{}
		_spec       = matq.querySpec()
		loadedTypes = [1]bool{
			matq.withCategories != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MasterAccountType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MasterAccountType{config: matq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, matq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := matq.withCategories; query != nil {
		if err := matq.loadCategories(ctx, query, nodes,
			func(n *MasterAccountType) { n.Edges.Categories = []*MasterAccountCategory{} },
			func(n *MasterAccountType, e *MasterAccountCategory) {
				n.Edges.Categories = append(n.Edges.Categories, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (matq *MasterAccountTypeQuery) loadCategories(ctx context.Context, query *MasterAccountCategoryQuery, nodes []*MasterAccountType, init func(*MasterAccountType), assign func(*MasterAccountType, *MasterAccountCategory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*MasterAccountType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.MasterAccountCategory(func(s *sql.Selector) {
		s.Where(sql.InValues(masteraccounttype.CategoriesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AccountTypeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_type_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (matq *MasterAccountTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := matq.querySpec()
	_spec.Node.Columns = matq.ctx.Fields
	if len(matq.ctx.Fields) > 0 {
		_spec.Unique = matq.ctx.Unique != nil && *matq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, matq.driver, _spec)
}

func (matq *MasterAccountTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(masteraccounttype.Table, masteraccounttype.Columns, sqlgraph.NewFieldSpec(masteraccounttype.FieldID, field.TypeInt))
	_spec.From = matq.sql
	if unique := matq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if matq.path != nil {
		_spec.Unique = true
	}
	if fields := matq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, masteraccounttype.FieldID)
		for i := range fields {
			if fields[i] != masteraccounttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := matq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := matq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := matq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := matq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (matq *MasterAccountTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(matq.driver.Dialect())
	t1 := builder.Table(masteraccounttype.Table)
	columns := matq.ctx.Fields
	if len(columns) == 0 {
		columns = masteraccounttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if matq.sql != nil {
		selector = matq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if matq.ctx.Unique != nil && *matq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range matq.predicates {
		p(selector)
	}
	for _, p := range matq.order {
		p(selector)
	}
	if offset := matq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := matq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MasterAccountTypeGroupBy is the group-by builder for MasterAccountType entities.
type MasterAccountTypeGroupBy struct {
	selector
	build *MasterAccountTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (matgb *MasterAccountTypeGroupBy) Aggregate(fns ...AggregateFunc) *MasterAccountTypeGroupBy {
	matgb.fns = append(matgb.fns, fns...)
	return matgb
}

// Scan applies the selector query and scans the result into the given value.
func (matgb *MasterAccountTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, matgb.build.ctx, "GroupBy")
	if err := matgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MasterAccountTypeQuery, *MasterAccountTypeGroupBy](ctx, matgb.build, matgb, matgb.build.inters, v)
}

func (matgb *MasterAccountTypeGroupBy) sqlScan(ctx context.Context, root *MasterAccountTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(matgb.fns))
	for _, fn := range matgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*matgb.flds)+len(matgb.fns))
		for _, f := range *matgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*matgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := matgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MasterAccountTypeSelect is the builder for selecting fields of MasterAccountType entities.
type MasterAccountTypeSelect struct {
	*MasterAccountTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mats *MasterAccountTypeSelect) Aggregate(fns ...AggregateFunc) *MasterAccountTypeSelect {
	mats.fns = append(mats.fns, fns...)
	return mats
}

// Scan applies the selector query and scans the result into the given value.
func (mats *MasterAccountTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mats.ctx, "Select")
	if err := mats.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MasterAccountTypeQuery, *MasterAccountTypeSelect](ctx, mats.MasterAccountTypeQuery, mats, mats.inters, v)
}

func (mats *MasterAccountTypeSelect) sqlScan(ctx context.Context, root *MasterAccountTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mats.fns))
	for _, fn := range mats.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mats.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
