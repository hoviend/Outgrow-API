// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"outgrow/ent/organizationaccount"
	"outgrow/ent/organizationaccountcategory"
	"outgrow/ent/organizationaccounttype"
	"outgrow/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationAccountCategoryQuery is the builder for querying OrganizationAccountCategory entities.
type OrganizationAccountCategoryQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.OrganizationAccountCategory
	withAccounts *OrganizationAccountQuery
	withType     *OrganizationAccountTypeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationAccountCategoryQuery builder.
func (oacq *OrganizationAccountCategoryQuery) Where(ps ...predicate.OrganizationAccountCategory) *OrganizationAccountCategoryQuery {
	oacq.predicates = append(oacq.predicates, ps...)
	return oacq
}

// Limit the number of records to be returned by this query.
func (oacq *OrganizationAccountCategoryQuery) Limit(limit int) *OrganizationAccountCategoryQuery {
	oacq.ctx.Limit = &limit
	return oacq
}

// Offset to start from.
func (oacq *OrganizationAccountCategoryQuery) Offset(offset int) *OrganizationAccountCategoryQuery {
	oacq.ctx.Offset = &offset
	return oacq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oacq *OrganizationAccountCategoryQuery) Unique(unique bool) *OrganizationAccountCategoryQuery {
	oacq.ctx.Unique = &unique
	return oacq
}

// Order specifies how the records should be ordered.
func (oacq *OrganizationAccountCategoryQuery) Order(o ...OrderFunc) *OrganizationAccountCategoryQuery {
	oacq.order = append(oacq.order, o...)
	return oacq
}

// QueryAccounts chains the current query on the "accounts" edge.
func (oacq *OrganizationAccountCategoryQuery) QueryAccounts() *OrganizationAccountQuery {
	query := (&OrganizationAccountClient{config: oacq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oacq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oacq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationaccountcategory.Table, organizationaccountcategory.FieldID, selector),
			sqlgraph.To(organizationaccount.Table, organizationaccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organizationaccountcategory.AccountsTable, organizationaccountcategory.AccountsColumn),
		)
		fromU = sqlgraph.SetNeighbors(oacq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryType chains the current query on the "type" edge.
func (oacq *OrganizationAccountCategoryQuery) QueryType() *OrganizationAccountTypeQuery {
	query := (&OrganizationAccountTypeClient{config: oacq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oacq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oacq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationaccountcategory.Table, organizationaccountcategory.FieldID, selector),
			sqlgraph.To(organizationaccounttype.Table, organizationaccounttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, organizationaccountcategory.TypeTable, organizationaccountcategory.TypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(oacq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrganizationAccountCategory entity from the query.
// Returns a *NotFoundError when no OrganizationAccountCategory was found.
func (oacq *OrganizationAccountCategoryQuery) First(ctx context.Context) (*OrganizationAccountCategory, error) {
	nodes, err := oacq.Limit(1).All(setContextOp(ctx, oacq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationaccountcategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oacq *OrganizationAccountCategoryQuery) FirstX(ctx context.Context) *OrganizationAccountCategory {
	node, err := oacq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationAccountCategory ID from the query.
// Returns a *NotFoundError when no OrganizationAccountCategory ID was found.
func (oacq *OrganizationAccountCategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oacq.Limit(1).IDs(setContextOp(ctx, oacq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationaccountcategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oacq *OrganizationAccountCategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := oacq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationAccountCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationAccountCategory entity is found.
// Returns a *NotFoundError when no OrganizationAccountCategory entities are found.
func (oacq *OrganizationAccountCategoryQuery) Only(ctx context.Context) (*OrganizationAccountCategory, error) {
	nodes, err := oacq.Limit(2).All(setContextOp(ctx, oacq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationaccountcategory.Label}
	default:
		return nil, &NotSingularError{organizationaccountcategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oacq *OrganizationAccountCategoryQuery) OnlyX(ctx context.Context) *OrganizationAccountCategory {
	node, err := oacq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationAccountCategory ID in the query.
// Returns a *NotSingularError when more than one OrganizationAccountCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (oacq *OrganizationAccountCategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oacq.Limit(2).IDs(setContextOp(ctx, oacq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationaccountcategory.Label}
	default:
		err = &NotSingularError{organizationaccountcategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oacq *OrganizationAccountCategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := oacq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationAccountCategories.
func (oacq *OrganizationAccountCategoryQuery) All(ctx context.Context) ([]*OrganizationAccountCategory, error) {
	ctx = setContextOp(ctx, oacq.ctx, "All")
	if err := oacq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrganizationAccountCategory, *OrganizationAccountCategoryQuery]()
	return withInterceptors[[]*OrganizationAccountCategory](ctx, oacq, qr, oacq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oacq *OrganizationAccountCategoryQuery) AllX(ctx context.Context) []*OrganizationAccountCategory {
	nodes, err := oacq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationAccountCategory IDs.
func (oacq *OrganizationAccountCategoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if oacq.ctx.Unique == nil && oacq.path != nil {
		oacq.Unique(true)
	}
	ctx = setContextOp(ctx, oacq.ctx, "IDs")
	if err = oacq.Select(organizationaccountcategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oacq *OrganizationAccountCategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := oacq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oacq *OrganizationAccountCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oacq.ctx, "Count")
	if err := oacq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oacq, querierCount[*OrganizationAccountCategoryQuery](), oacq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oacq *OrganizationAccountCategoryQuery) CountX(ctx context.Context) int {
	count, err := oacq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oacq *OrganizationAccountCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oacq.ctx, "Exist")
	switch _, err := oacq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oacq *OrganizationAccountCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := oacq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationAccountCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oacq *OrganizationAccountCategoryQuery) Clone() *OrganizationAccountCategoryQuery {
	if oacq == nil {
		return nil
	}
	return &OrganizationAccountCategoryQuery{
		config:       oacq.config,
		ctx:          oacq.ctx.Clone(),
		order:        append([]OrderFunc{}, oacq.order...),
		inters:       append([]Interceptor{}, oacq.inters...),
		predicates:   append([]predicate.OrganizationAccountCategory{}, oacq.predicates...),
		withAccounts: oacq.withAccounts.Clone(),
		withType:     oacq.withType.Clone(),
		// clone intermediate query.
		sql:  oacq.sql.Clone(),
		path: oacq.path,
	}
}

// WithAccounts tells the query-builder to eager-load the nodes that are connected to
// the "accounts" edge. The optional arguments are used to configure the query builder of the edge.
func (oacq *OrganizationAccountCategoryQuery) WithAccounts(opts ...func(*OrganizationAccountQuery)) *OrganizationAccountCategoryQuery {
	query := (&OrganizationAccountClient{config: oacq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oacq.withAccounts = query
	return oacq
}

// WithType tells the query-builder to eager-load the nodes that are connected to
// the "type" edge. The optional arguments are used to configure the query builder of the edge.
func (oacq *OrganizationAccountCategoryQuery) WithType(opts ...func(*OrganizationAccountTypeQuery)) *OrganizationAccountCategoryQuery {
	query := (&OrganizationAccountTypeClient{config: oacq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oacq.withType = query
	return oacq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AccountTypeID int `json:"account_type_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrganizationAccountCategory.Query().
//		GroupBy(organizationaccountcategory.FieldAccountTypeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oacq *OrganizationAccountCategoryQuery) GroupBy(field string, fields ...string) *OrganizationAccountCategoryGroupBy {
	oacq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationAccountCategoryGroupBy{build: oacq}
	grbuild.flds = &oacq.ctx.Fields
	grbuild.label = organizationaccountcategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AccountTypeID int `json:"account_type_id,omitempty"`
//	}
//
//	client.OrganizationAccountCategory.Query().
//		Select(organizationaccountcategory.FieldAccountTypeID).
//		Scan(ctx, &v)
func (oacq *OrganizationAccountCategoryQuery) Select(fields ...string) *OrganizationAccountCategorySelect {
	oacq.ctx.Fields = append(oacq.ctx.Fields, fields...)
	sbuild := &OrganizationAccountCategorySelect{OrganizationAccountCategoryQuery: oacq}
	sbuild.label = organizationaccountcategory.Label
	sbuild.flds, sbuild.scan = &oacq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationAccountCategorySelect configured with the given aggregations.
func (oacq *OrganizationAccountCategoryQuery) Aggregate(fns ...AggregateFunc) *OrganizationAccountCategorySelect {
	return oacq.Select().Aggregate(fns...)
}

func (oacq *OrganizationAccountCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oacq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oacq); err != nil {
				return err
			}
		}
	}
	for _, f := range oacq.ctx.Fields {
		if !organizationaccountcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oacq.path != nil {
		prev, err := oacq.path(ctx)
		if err != nil {
			return err
		}
		oacq.sql = prev
	}
	return nil
}

func (oacq *OrganizationAccountCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationAccountCategory, error) {
	var (
		nodes       = []*OrganizationAccountCategory{}
		_spec       = oacq.querySpec()
		loadedTypes = [2]bool{
			oacq.withAccounts != nil,
			oacq.withType != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrganizationAccountCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrganizationAccountCategory{config: oacq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oacq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oacq.withAccounts; query != nil {
		if err := oacq.loadAccounts(ctx, query, nodes,
			func(n *OrganizationAccountCategory) { n.Edges.Accounts = []*OrganizationAccount{} },
			func(n *OrganizationAccountCategory, e *OrganizationAccount) {
				n.Edges.Accounts = append(n.Edges.Accounts, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := oacq.withType; query != nil {
		if err := oacq.loadType(ctx, query, nodes, nil,
			func(n *OrganizationAccountCategory, e *OrganizationAccountType) { n.Edges.Type = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oacq *OrganizationAccountCategoryQuery) loadAccounts(ctx context.Context, query *OrganizationAccountQuery, nodes []*OrganizationAccountCategory, init func(*OrganizationAccountCategory), assign func(*OrganizationAccountCategory, *OrganizationAccount)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*OrganizationAccountCategory)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.OrganizationAccount(func(s *sql.Selector) {
		s.Where(sql.InValues(organizationaccountcategory.AccountsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CategoryID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "category_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oacq *OrganizationAccountCategoryQuery) loadType(ctx context.Context, query *OrganizationAccountTypeQuery, nodes []*OrganizationAccountCategory, init func(*OrganizationAccountCategory), assign func(*OrganizationAccountCategory, *OrganizationAccountType)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OrganizationAccountCategory)
	for i := range nodes {
		fk := nodes[i].AccountTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organizationaccounttype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (oacq *OrganizationAccountCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oacq.querySpec()
	_spec.Node.Columns = oacq.ctx.Fields
	if len(oacq.ctx.Fields) > 0 {
		_spec.Unique = oacq.ctx.Unique != nil && *oacq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oacq.driver, _spec)
}

func (oacq *OrganizationAccountCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organizationaccountcategory.Table, organizationaccountcategory.Columns, sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt))
	_spec.From = oacq.sql
	if unique := oacq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oacq.path != nil {
		_spec.Unique = true
	}
	if fields := oacq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationaccountcategory.FieldID)
		for i := range fields {
			if fields[i] != organizationaccountcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oacq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oacq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oacq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oacq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oacq *OrganizationAccountCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oacq.driver.Dialect())
	t1 := builder.Table(organizationaccountcategory.Table)
	columns := oacq.ctx.Fields
	if len(columns) == 0 {
		columns = organizationaccountcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oacq.sql != nil {
		selector = oacq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oacq.ctx.Unique != nil && *oacq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oacq.predicates {
		p(selector)
	}
	for _, p := range oacq.order {
		p(selector)
	}
	if offset := oacq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oacq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrganizationAccountCategoryGroupBy is the group-by builder for OrganizationAccountCategory entities.
type OrganizationAccountCategoryGroupBy struct {
	selector
	build *OrganizationAccountCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oacgb *OrganizationAccountCategoryGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationAccountCategoryGroupBy {
	oacgb.fns = append(oacgb.fns, fns...)
	return oacgb
}

// Scan applies the selector query and scans the result into the given value.
func (oacgb *OrganizationAccountCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oacgb.build.ctx, "GroupBy")
	if err := oacgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationAccountCategoryQuery, *OrganizationAccountCategoryGroupBy](ctx, oacgb.build, oacgb, oacgb.build.inters, v)
}

func (oacgb *OrganizationAccountCategoryGroupBy) sqlScan(ctx context.Context, root *OrganizationAccountCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oacgb.fns))
	for _, fn := range oacgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oacgb.flds)+len(oacgb.fns))
		for _, f := range *oacgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oacgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oacgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationAccountCategorySelect is the builder for selecting fields of OrganizationAccountCategory entities.
type OrganizationAccountCategorySelect struct {
	*OrganizationAccountCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oacs *OrganizationAccountCategorySelect) Aggregate(fns ...AggregateFunc) *OrganizationAccountCategorySelect {
	oacs.fns = append(oacs.fns, fns...)
	return oacs
}

// Scan applies the selector query and scans the result into the given value.
func (oacs *OrganizationAccountCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oacs.ctx, "Select")
	if err := oacs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationAccountCategoryQuery, *OrganizationAccountCategorySelect](ctx, oacs.OrganizationAccountCategoryQuery, oacs, oacs.inters, v)
}

func (oacs *OrganizationAccountCategorySelect) sqlScan(ctx context.Context, root *OrganizationAccountCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oacs.fns))
	for _, fn := range oacs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oacs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oacs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
