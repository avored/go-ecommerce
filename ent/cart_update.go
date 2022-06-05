// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/avored/go-ecommerce/ent/cart"
	"github.com/avored/go-ecommerce/ent/cartitem"
	"github.com/avored/go-ecommerce/ent/predicate"
)

// CartUpdate is the builder for updating Cart entities.
type CartUpdate struct {
	config
	hooks    []Hook
	mutation *CartMutation
}

// Where appends a list predicates to the CartUpdate builder.
func (cu *CartUpdate) Where(ps ...predicate.Cart) *CartUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CartUpdate) SetUpdatedAt(t time.Time) *CartUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetSessionID sets the "session_id" field.
func (cu *CartUpdate) SetSessionID(s string) *CartUpdate {
	cu.mutation.SetSessionID(s)
	return cu
}

// SetFullAmount sets the "full_amount" field.
func (cu *CartUpdate) SetFullAmount(f float64) *CartUpdate {
	cu.mutation.ResetFullAmount()
	cu.mutation.SetFullAmount(f)
	return cu
}

// AddFullAmount adds f to the "full_amount" field.
func (cu *CartUpdate) AddFullAmount(f float64) *CartUpdate {
	cu.mutation.AddFullAmount(f)
	return cu
}

// AddCartItemIDs adds the "cart_items" edge to the CartItem entity by IDs.
func (cu *CartUpdate) AddCartItemIDs(ids ...int) *CartUpdate {
	cu.mutation.AddCartItemIDs(ids...)
	return cu
}

// AddCartItems adds the "cart_items" edges to the CartItem entity.
func (cu *CartUpdate) AddCartItems(c ...*CartItem) *CartUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCartItemIDs(ids...)
}

// Mutation returns the CartMutation object of the builder.
func (cu *CartUpdate) Mutation() *CartMutation {
	return cu.mutation
}

// ClearCartItems clears all "cart_items" edges to the CartItem entity.
func (cu *CartUpdate) ClearCartItems() *CartUpdate {
	cu.mutation.ClearCartItems()
	return cu
}

// RemoveCartItemIDs removes the "cart_items" edge to CartItem entities by IDs.
func (cu *CartUpdate) RemoveCartItemIDs(ids ...int) *CartUpdate {
	cu.mutation.RemoveCartItemIDs(ids...)
	return cu
}

// RemoveCartItems removes "cart_items" edges to CartItem entities.
func (cu *CartUpdate) RemoveCartItems(c ...*CartItem) *CartUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCartItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CartUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CartUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CartUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CartUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CartUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := cart.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cart.Table,
			Columns: cart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cart.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cart.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.SessionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cart.FieldSessionID,
		})
	}
	if value, ok := cu.mutation.FullAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: cart.FieldFullAmount,
		})
	}
	if value, ok := cu.mutation.AddedFullAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: cart.FieldFullAmount,
		})
	}
	if cu.mutation.CartItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.CartItemsTable,
			Columns: []string{cart.CartItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cartitem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCartItemsIDs(); len(nodes) > 0 && !cu.mutation.CartItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.CartItemsTable,
			Columns: []string{cart.CartItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cartitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CartItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.CartItemsTable,
			Columns: []string{cart.CartItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cartitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CartUpdateOne is the builder for updating a single Cart entity.
type CartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CartUpdateOne) SetUpdatedAt(t time.Time) *CartUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetSessionID sets the "session_id" field.
func (cuo *CartUpdateOne) SetSessionID(s string) *CartUpdateOne {
	cuo.mutation.SetSessionID(s)
	return cuo
}

// SetFullAmount sets the "full_amount" field.
func (cuo *CartUpdateOne) SetFullAmount(f float64) *CartUpdateOne {
	cuo.mutation.ResetFullAmount()
	cuo.mutation.SetFullAmount(f)
	return cuo
}

// AddFullAmount adds f to the "full_amount" field.
func (cuo *CartUpdateOne) AddFullAmount(f float64) *CartUpdateOne {
	cuo.mutation.AddFullAmount(f)
	return cuo
}

// AddCartItemIDs adds the "cart_items" edge to the CartItem entity by IDs.
func (cuo *CartUpdateOne) AddCartItemIDs(ids ...int) *CartUpdateOne {
	cuo.mutation.AddCartItemIDs(ids...)
	return cuo
}

// AddCartItems adds the "cart_items" edges to the CartItem entity.
func (cuo *CartUpdateOne) AddCartItems(c ...*CartItem) *CartUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCartItemIDs(ids...)
}

// Mutation returns the CartMutation object of the builder.
func (cuo *CartUpdateOne) Mutation() *CartMutation {
	return cuo.mutation
}

// ClearCartItems clears all "cart_items" edges to the CartItem entity.
func (cuo *CartUpdateOne) ClearCartItems() *CartUpdateOne {
	cuo.mutation.ClearCartItems()
	return cuo
}

// RemoveCartItemIDs removes the "cart_items" edge to CartItem entities by IDs.
func (cuo *CartUpdateOne) RemoveCartItemIDs(ids ...int) *CartUpdateOne {
	cuo.mutation.RemoveCartItemIDs(ids...)
	return cuo
}

// RemoveCartItems removes "cart_items" edges to CartItem entities.
func (cuo *CartUpdateOne) RemoveCartItems(c ...*CartItem) *CartUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCartItemIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CartUpdateOne) Select(field string, fields ...string) *CartUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cart entity.
func (cuo *CartUpdateOne) Save(ctx context.Context) (*Cart, error) {
	var (
		err  error
		node *Cart
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CartUpdateOne) SaveX(ctx context.Context) *Cart {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CartUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CartUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CartUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := cart.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CartUpdateOne) sqlSave(ctx context.Context) (_node *Cart, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cart.Table,
			Columns: cart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cart.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cart.FieldID)
		for _, f := range fields {
			if !cart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cart.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.SessionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cart.FieldSessionID,
		})
	}
	if value, ok := cuo.mutation.FullAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: cart.FieldFullAmount,
		})
	}
	if value, ok := cuo.mutation.AddedFullAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: cart.FieldFullAmount,
		})
	}
	if cuo.mutation.CartItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.CartItemsTable,
			Columns: []string{cart.CartItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cartitem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCartItemsIDs(); len(nodes) > 0 && !cuo.mutation.CartItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.CartItemsTable,
			Columns: []string{cart.CartItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cartitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CartItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.CartItemsTable,
			Columns: []string{cart.CartItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cartitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cart{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}