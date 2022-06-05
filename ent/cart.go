// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/avored/go-ecommerce/ent/cart"
)

// Cart is the model entity for the Cart schema.
type Cart struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// SessionID holds the value of the "session_id" field.
	SessionID string `json:"session_id,omitempty"`
	// FullAmount holds the value of the "full_amount" field.
	FullAmount float64 `json:"full_amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CartQuery when eager-loading is set.
	Edges CartEdges `json:"edges"`
}

// CartEdges holds the relations/edges for other nodes in the graph.
type CartEdges struct {
	// CartItems holds the value of the cart_items edge.
	CartItems []*CartItem `json:"cart_items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CartItemsOrErr returns the CartItems value or an error if the edge
// was not loaded in eager-loading.
func (e CartEdges) CartItemsOrErr() ([]*CartItem, error) {
	if e.loadedTypes[0] {
		return e.CartItems, nil
	}
	return nil, &NotLoadedError{edge: "cart_items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cart) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cart.FieldFullAmount:
			values[i] = new(sql.NullFloat64)
		case cart.FieldID:
			values[i] = new(sql.NullInt64)
		case cart.FieldSessionID:
			values[i] = new(sql.NullString)
		case cart.FieldCreatedAt, cart.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cart", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cart fields.
func (c *Cart) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cart.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cart.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cart.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case cart.FieldSessionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field session_id", values[i])
			} else if value.Valid {
				c.SessionID = value.String
			}
		case cart.FieldFullAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field full_amount", values[i])
			} else if value.Valid {
				c.FullAmount = value.Float64
			}
		}
	}
	return nil
}

// QueryCartItems queries the "cart_items" edge of the Cart entity.
func (c *Cart) QueryCartItems() *CartItemQuery {
	return (&CartClient{config: c.config}).QueryCartItems(c)
}

// Update returns a builder for updating this Cart.
// Note that you need to call Cart.Unwrap() before calling this method if this Cart
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cart) Update() *CartUpdateOne {
	return (&CartClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cart) Unwrap() *Cart {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cart is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cart) String() string {
	var builder strings.Builder
	builder.WriteString("Cart(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", session_id=")
	builder.WriteString(c.SessionID)
	builder.WriteString(", full_amount=")
	builder.WriteString(fmt.Sprintf("%v", c.FullAmount))
	builder.WriteByte(')')
	return builder.String()
}

// Carts is a parsable slice of Cart.
type Carts []*Cart

func (c Carts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
