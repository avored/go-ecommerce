// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/avored/go-ecommerce/ent/adminuser"
	"github.com/avored/go-ecommerce/ent/predicate"
)

// AdminUserQuery is the builder for querying AdminUser entities.
type AdminUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AdminUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminUserQuery builder.
func (auq *AdminUserQuery) Where(ps ...predicate.AdminUser) *AdminUserQuery {
	auq.predicates = append(auq.predicates, ps...)
	return auq
}

// Limit adds a limit step to the query.
func (auq *AdminUserQuery) Limit(limit int) *AdminUserQuery {
	auq.limit = &limit
	return auq
}

// Offset adds an offset step to the query.
func (auq *AdminUserQuery) Offset(offset int) *AdminUserQuery {
	auq.offset = &offset
	return auq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (auq *AdminUserQuery) Unique(unique bool) *AdminUserQuery {
	auq.unique = &unique
	return auq
}

// Order adds an order step to the query.
func (auq *AdminUserQuery) Order(o ...OrderFunc) *AdminUserQuery {
	auq.order = append(auq.order, o...)
	return auq
}

// First returns the first AdminUser entity from the query.
// Returns a *NotFoundError when no AdminUser was found.
func (auq *AdminUserQuery) First(ctx context.Context) (*AdminUser, error) {
	nodes, err := auq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adminuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auq *AdminUserQuery) FirstX(ctx context.Context) *AdminUser {
	node, err := auq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdminUser ID from the query.
// Returns a *NotFoundError when no AdminUser ID was found.
func (auq *AdminUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = auq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adminuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (auq *AdminUserQuery) FirstIDX(ctx context.Context) int {
	id, err := auq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdminUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdminUser entity is found.
// Returns a *NotFoundError when no AdminUser entities are found.
func (auq *AdminUserQuery) Only(ctx context.Context) (*AdminUser, error) {
	nodes, err := auq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adminuser.Label}
	default:
		return nil, &NotSingularError{adminuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auq *AdminUserQuery) OnlyX(ctx context.Context) *AdminUser {
	node, err := auq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdminUser ID in the query.
// Returns a *NotSingularError when more than one AdminUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (auq *AdminUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = auq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adminuser.Label}
	default:
		err = &NotSingularError{adminuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auq *AdminUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := auq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdminUsers.
func (auq *AdminUserQuery) All(ctx context.Context) ([]*AdminUser, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return auq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (auq *AdminUserQuery) AllX(ctx context.Context) []*AdminUser {
	nodes, err := auq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdminUser IDs.
func (auq *AdminUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := auq.Select(adminuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auq *AdminUserQuery) IDsX(ctx context.Context) []int {
	ids, err := auq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auq *AdminUserQuery) Count(ctx context.Context) (int, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return auq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (auq *AdminUserQuery) CountX(ctx context.Context) int {
	count, err := auq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auq *AdminUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return auq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (auq *AdminUserQuery) ExistX(ctx context.Context) bool {
	exist, err := auq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auq *AdminUserQuery) Clone() *AdminUserQuery {
	if auq == nil {
		return nil
	}
	return &AdminUserQuery{
		config:     auq.config,
		limit:      auq.limit,
		offset:     auq.offset,
		order:      append([]OrderFunc{}, auq.order...),
		predicates: append([]predicate.AdminUser{}, auq.predicates...),
		// clone intermediate query.
		sql:    auq.sql.Clone(),
		path:   auq.path,
		unique: auq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdminUser.Query().
//		GroupBy(adminuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (auq *AdminUserQuery) GroupBy(field string, fields ...string) *AdminUserGroupBy {
	grbuild := &AdminUserGroupBy{config: auq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := auq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return auq.sqlQuery(ctx), nil
	}
	grbuild.label = adminuser.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AdminUser.Query().
//		Select(adminuser.FieldCreatedAt).
//		Scan(ctx, &v)
func (auq *AdminUserQuery) Select(fields ...string) *AdminUserSelect {
	auq.fields = append(auq.fields, fields...)
	selbuild := &AdminUserSelect{AdminUserQuery: auq}
	selbuild.label = adminuser.Label
	selbuild.flds, selbuild.scan = &auq.fields, selbuild.Scan
	return selbuild
}

func (auq *AdminUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range auq.fields {
		if !adminuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if auq.path != nil {
		prev, err := auq.path(ctx)
		if err != nil {
			return err
		}
		auq.sql = prev
	}
	return nil
}

func (auq *AdminUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdminUser, error) {
	var (
		nodes = []*AdminUser{}
		_spec = auq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AdminUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AdminUser{config: auq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, auq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (auq *AdminUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auq.querySpec()
	_spec.Node.Columns = auq.fields
	if len(auq.fields) > 0 {
		_spec.Unique = auq.unique != nil && *auq.unique
	}
	return sqlgraph.CountNodes(ctx, auq.driver, _spec)
}

func (auq *AdminUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := auq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (auq *AdminUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adminuser.Table,
			Columns: adminuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adminuser.FieldID,
			},
		},
		From:   auq.sql,
		Unique: true,
	}
	if unique := auq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := auq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminuser.FieldID)
		for i := range fields {
			if fields[i] != adminuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := auq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (auq *AdminUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(auq.driver.Dialect())
	t1 := builder.Table(adminuser.Table)
	columns := auq.fields
	if len(columns) == 0 {
		columns = adminuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if auq.sql != nil {
		selector = auq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if auq.unique != nil && *auq.unique {
		selector.Distinct()
	}
	for _, p := range auq.predicates {
		p(selector)
	}
	for _, p := range auq.order {
		p(selector)
	}
	if offset := auq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdminUserGroupBy is the group-by builder for AdminUser entities.
type AdminUserGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (augb *AdminUserGroupBy) Aggregate(fns ...AggregateFunc) *AdminUserGroupBy {
	augb.fns = append(augb.fns, fns...)
	return augb
}

// Scan applies the group-by query and scans the result into the given value.
func (augb *AdminUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := augb.path(ctx)
	if err != nil {
		return err
	}
	augb.sql = query
	return augb.sqlScan(ctx, v)
}

func (augb *AdminUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range augb.fields {
		if !adminuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := augb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := augb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (augb *AdminUserGroupBy) sqlQuery() *sql.Selector {
	selector := augb.sql.Select()
	aggregation := make([]string, 0, len(augb.fns))
	for _, fn := range augb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(augb.fields)+len(augb.fns))
		for _, f := range augb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(augb.fields...)...)
}

// AdminUserSelect is the builder for selecting fields of AdminUser entities.
type AdminUserSelect struct {
	*AdminUserQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aus *AdminUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aus.prepareQuery(ctx); err != nil {
		return err
	}
	aus.sql = aus.AdminUserQuery.sqlQuery(ctx)
	return aus.sqlScan(ctx, v)
}

func (aus *AdminUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aus.sql.Query()
	if err := aus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
