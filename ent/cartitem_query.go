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
	"github.com/avored/go-ecommerce/ent/cart"
	"github.com/avored/go-ecommerce/ent/cartitem"
	"github.com/avored/go-ecommerce/ent/predicate"
	"github.com/avored/go-ecommerce/ent/product"
)

// CartItemQuery is the builder for querying CartItem entities.
type CartItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CartItem
	// eager-loading edges.
	withCart    *CartQuery
	withProduct *ProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CartItemQuery builder.
func (ciq *CartItemQuery) Where(ps ...predicate.CartItem) *CartItemQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit adds a limit step to the query.
func (ciq *CartItemQuery) Limit(limit int) *CartItemQuery {
	ciq.limit = &limit
	return ciq
}

// Offset adds an offset step to the query.
func (ciq *CartItemQuery) Offset(offset int) *CartItemQuery {
	ciq.offset = &offset
	return ciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciq *CartItemQuery) Unique(unique bool) *CartItemQuery {
	ciq.unique = &unique
	return ciq
}

// Order adds an order step to the query.
func (ciq *CartItemQuery) Order(o ...OrderFunc) *CartItemQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// QueryCart chains the current query on the "cart" edge.
func (ciq *CartItemQuery) QueryCart() *CartQuery {
	query := &CartQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cartitem.Table, cartitem.FieldID, selector),
			sqlgraph.To(cart.Table, cart.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cartitem.CartTable, cartitem.CartColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the "product" edge.
func (ciq *CartItemQuery) QueryProduct() *ProductQuery {
	query := &ProductQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cartitem.Table, cartitem.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cartitem.ProductTable, cartitem.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CartItem entity from the query.
// Returns a *NotFoundError when no CartItem was found.
func (ciq *CartItemQuery) First(ctx context.Context) (*CartItem, error) {
	nodes, err := ciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cartitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CartItemQuery) FirstX(ctx context.Context) *CartItem {
	node, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CartItem ID from the query.
// Returns a *NotFoundError when no CartItem ID was found.
func (ciq *CartItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cartitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciq *CartItemQuery) FirstIDX(ctx context.Context) int {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CartItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CartItem entity is found.
// Returns a *NotFoundError when no CartItem entities are found.
func (ciq *CartItemQuery) Only(ctx context.Context) (*CartItem, error) {
	nodes, err := ciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cartitem.Label}
	default:
		return nil, &NotSingularError{cartitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CartItemQuery) OnlyX(ctx context.Context) *CartItem {
	node, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CartItem ID in the query.
// Returns a *NotSingularError when more than one CartItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (ciq *CartItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = &NotSingularError{cartitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CartItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CartItems.
func (ciq *CartItemQuery) All(ctx context.Context) ([]*CartItem, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CartItemQuery) AllX(ctx context.Context) []*CartItem {
	nodes, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CartItem IDs.
func (ciq *CartItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ciq.Select(cartitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CartItemQuery) IDsX(ctx context.Context) []int {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CartItemQuery) Count(ctx context.Context) (int, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CartItemQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CartItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CartItemQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CartItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CartItemQuery) Clone() *CartItemQuery {
	if ciq == nil {
		return nil
	}
	return &CartItemQuery{
		config:      ciq.config,
		limit:       ciq.limit,
		offset:      ciq.offset,
		order:       append([]OrderFunc{}, ciq.order...),
		predicates:  append([]predicate.CartItem{}, ciq.predicates...),
		withCart:    ciq.withCart.Clone(),
		withProduct: ciq.withProduct.Clone(),
		// clone intermediate query.
		sql:    ciq.sql.Clone(),
		path:   ciq.path,
		unique: ciq.unique,
	}
}

// WithCart tells the query-builder to eager-load the nodes that are connected to
// the "cart" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CartItemQuery) WithCart(opts ...func(*CartQuery)) *CartItemQuery {
	query := &CartQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withCart = query
	return ciq
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CartItemQuery) WithProduct(opts ...func(*ProductQuery)) *CartItemQuery {
	query := &ProductQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withProduct = query
	return ciq
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
//	client.CartItem.Query().
//		GroupBy(cartitem.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ciq *CartItemQuery) GroupBy(field string, fields ...string) *CartItemGroupBy {
	group := &CartItemGroupBy{config: ciq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(ctx), nil
	}
	return group
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
//	client.CartItem.Query().
//		Select(cartitem.FieldName).
//		Scan(ctx, &v)
//
func (ciq *CartItemQuery) Select(fields ...string) *CartItemSelect {
	ciq.fields = append(ciq.fields, fields...)
	return &CartItemSelect{CartItemQuery: ciq}
}

func (ciq *CartItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ciq.fields {
		if !cartitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *CartItemQuery) sqlAll(ctx context.Context) ([]*CartItem, error) {
	var (
		nodes       = []*CartItem{}
		_spec       = ciq.querySpec()
		loadedTypes = [2]bool{
			ciq.withCart != nil,
			ciq.withProduct != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CartItem{config: ciq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ciq.withCart; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CartItem)
		for i := range nodes {
			fk := nodes[i].CartID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(cart.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "cart_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Cart = n
			}
		}
	}

	if query := ciq.withProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CartItem)
		for i := range nodes {
			fk := nodes[i].ProductID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(product.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Product = n
			}
		}
	}

	return nodes, nil
}

func (ciq *CartItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	_spec.Node.Columns = ciq.fields
	if len(ciq.fields) > 0 {
		_spec.Unique = ciq.unique != nil && *ciq.unique
	}
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CartItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ciq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ciq *CartItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cartitem.Table,
			Columns: cartitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cartitem.FieldID,
			},
		},
		From:   ciq.sql,
		Unique: true,
	}
	if unique := ciq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ciq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cartitem.FieldID)
		for i := range fields {
			if fields[i] != cartitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CartItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(cartitem.Table)
	columns := ciq.fields
	if len(columns) == 0 {
		columns = cartitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ciq.unique != nil && *ciq.unique {
		selector.Distinct()
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CartItemGroupBy is the group-by builder for CartItem entities.
type CartItemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CartItemGroupBy) Aggregate(fns ...AggregateFunc) *CartItemGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the group-by query and scans the result into the given value.
func (cigb *CartItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cigb.path(ctx)
	if err != nil {
		return err
	}
	cigb.sql = query
	return cigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cigb *CartItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cigb *CartItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CartItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cigb *CartItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := cigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cigb *CartItemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = fmt.Errorf("ent: CartItemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cigb *CartItemGroupBy) StringX(ctx context.Context) string {
	v, err := cigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cigb *CartItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CartItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cigb *CartItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := cigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cigb *CartItemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = fmt.Errorf("ent: CartItemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cigb *CartItemGroupBy) IntX(ctx context.Context) int {
	v, err := cigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cigb *CartItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CartItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cigb *CartItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cigb *CartItemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = fmt.Errorf("ent: CartItemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cigb *CartItemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cigb *CartItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CartItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cigb *CartItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cigb *CartItemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = fmt.Errorf("ent: CartItemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cigb *CartItemGroupBy) BoolX(ctx context.Context) bool {
	v, err := cigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cigb *CartItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cigb.fields {
		if !cartitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cigb *CartItemGroupBy) sqlQuery() *sql.Selector {
	selector := cigb.sql.Select()
	aggregation := make([]string, 0, len(cigb.fns))
	for _, fn := range cigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cigb.fields)+len(cigb.fns))
		for _, f := range cigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cigb.fields...)...)
}

// CartItemSelect is the builder for selecting fields of CartItem entities.
type CartItemSelect struct {
	*CartItemQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cis *CartItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cis.prepareQuery(ctx); err != nil {
		return err
	}
	cis.sql = cis.CartItemQuery.sqlQuery(ctx)
	return cis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cis *CartItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cis *CartItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CartItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cis *CartItemSelect) StringsX(ctx context.Context) []string {
	v, err := cis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cis *CartItemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = fmt.Errorf("ent: CartItemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cis *CartItemSelect) StringX(ctx context.Context) string {
	v, err := cis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cis *CartItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CartItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cis *CartItemSelect) IntsX(ctx context.Context) []int {
	v, err := cis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cis *CartItemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = fmt.Errorf("ent: CartItemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cis *CartItemSelect) IntX(ctx context.Context) int {
	v, err := cis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cis *CartItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CartItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cis *CartItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cis *CartItemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = fmt.Errorf("ent: CartItemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cis *CartItemSelect) Float64X(ctx context.Context) float64 {
	v, err := cis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cis *CartItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CartItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cis *CartItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := cis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cis *CartItemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = fmt.Errorf("ent: CartItemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cis *CartItemSelect) BoolX(ctx context.Context) bool {
	v, err := cis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cis *CartItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cis.sql.Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
