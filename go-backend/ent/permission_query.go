// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/avored/go-ecommerce/ent/permission"
	"github.com/avored/go-ecommerce/ent/predicate"
	"github.com/avored/go-ecommerce/ent/role"
)

// PermissionQuery is the builder for querying Permission entities.
type PermissionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Permission
	// eager-loading edges.
	withRoles *RoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PermissionQuery builder.
func (pq *PermissionQuery) Where(ps ...predicate.Permission) *PermissionQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PermissionQuery) Limit(limit int) *PermissionQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PermissionQuery) Offset(offset int) *PermissionQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PermissionQuery) Unique(unique bool) *PermissionQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PermissionQuery) Order(o ...OrderFunc) *PermissionQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryRoles chains the current query on the "roles" edge.
func (pq *PermissionQuery) QueryRoles() *RoleQuery {
	query := &RoleQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, permission.RolesTable, permission.RolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Permission entity from the query.
// Returns a *NotFoundError when no Permission was found.
func (pq *PermissionQuery) First(ctx context.Context) (*Permission, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{permission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PermissionQuery) FirstX(ctx context.Context) *Permission {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Permission ID from the query.
// Returns a *NotFoundError when no Permission ID was found.
func (pq *PermissionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{permission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PermissionQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Permission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Permission entity is found.
// Returns a *NotFoundError when no Permission entities are found.
func (pq *PermissionQuery) Only(ctx context.Context) (*Permission, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{permission.Label}
	default:
		return nil, &NotSingularError{permission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PermissionQuery) OnlyX(ctx context.Context) *Permission {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Permission ID in the query.
// Returns a *NotSingularError when more than one Permission ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PermissionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{permission.Label}
	default:
		err = &NotSingularError{permission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PermissionQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Permissions.
func (pq *PermissionQuery) All(ctx context.Context) ([]*Permission, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PermissionQuery) AllX(ctx context.Context) []*Permission {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Permission IDs.
func (pq *PermissionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(permission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PermissionQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PermissionQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PermissionQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PermissionQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PermissionQuery) Clone() *PermissionQuery {
	if pq == nil {
		return nil
	}
	return &PermissionQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		predicates: append([]predicate.Permission{}, pq.predicates...),
		withRoles:  pq.withRoles.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PermissionQuery) WithRoles(opts ...func(*RoleQuery)) *PermissionQuery {
	query := &RoleQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withRoles = query
	return pq
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
//	client.Permission.Query().
//		GroupBy(permission.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PermissionQuery) GroupBy(field string, fields ...string) *PermissionGroupBy {
	grbuild := &PermissionGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = permission.Label
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
//	client.Permission.Query().
//		Select(permission.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PermissionQuery) Select(fields ...string) *PermissionSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &PermissionSelect{PermissionQuery: pq}
	selbuild.label = permission.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

func (pq *PermissionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !permission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PermissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Permission, error) {
	var (
		nodes       = []*Permission{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withRoles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Permission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Permission{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withRoles; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[int]*Permission)
		nids := make(map[int]map[*Permission]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Roles = []*Role{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(permission.RolesTable)
			s.Join(joinT).On(s.C(role.FieldID), joinT.C(permission.RolesPrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(permission.RolesPrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(permission.RolesPrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Permission]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "roles" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Roles = append(kn.Edges.Roles, n)
			}
		}
	}

	return nodes, nil
}

func (pq *PermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PermissionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *PermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permission.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for i := range fields {
			if fields[i] != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(permission.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = permission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PermissionGroupBy is the group-by builder for Permission entities.
type PermissionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PermissionGroupBy) Aggregate(fns ...AggregateFunc) *PermissionGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PermissionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *PermissionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !permission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PermissionGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PermissionSelect is the builder for selecting fields of Permission entities.
type PermissionSelect struct {
	*PermissionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PermissionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PermissionQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *PermissionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}