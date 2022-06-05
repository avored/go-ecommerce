// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/avored/go-ecommerce/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// MetaTitle holds the value of the "meta_title" field.
	MetaTitle string `json:"meta_title,omitempty"`
	// MetaDescription holds the value of the "meta_description" field.
	MetaDescription string `json:"meta_description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges ProductEdges `json:"edges"`
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// CartItems holds the value of the cart_items edge.
	CartItems []*CartItem `json:"cart_items,omitempty"`
	// OrderProducts holds the value of the order_products edge.
	OrderProducts []*OrderProduct `json:"order_products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[0] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// CartItemsOrErr returns the CartItems value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) CartItemsOrErr() ([]*CartItem, error) {
	if e.loadedTypes[1] {
		return e.CartItems, nil
	}
	return nil, &NotLoadedError{edge: "cart_items"}
}

// OrderProductsOrErr returns the OrderProducts value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) OrderProductsOrErr() ([]*OrderProduct, error) {
	if e.loadedTypes[2] {
		return e.OrderProducts, nil
	}
	return nil, &NotLoadedError{edge: "order_products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldSlug, product.FieldMetaTitle, product.FieldMetaDescription:
			values[i] = new(sql.NullString)
		case product.FieldCreatedAt, product.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Product", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				pr.Slug = value.String
			}
		case product.FieldMetaTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_title", values[i])
			} else if value.Valid {
				pr.MetaTitle = value.String
			}
		case product.FieldMetaDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_description", values[i])
			} else if value.Valid {
				pr.MetaDescription = value.String
			}
		}
	}
	return nil
}

// QueryCategories queries the "categories" edge of the Product entity.
func (pr *Product) QueryCategories() *CategoryQuery {
	return (&ProductClient{config: pr.config}).QueryCategories(pr)
}

// QueryCartItems queries the "cart_items" edge of the Product entity.
func (pr *Product) QueryCartItems() *CartItemQuery {
	return (&ProductClient{config: pr.config}).QueryCartItems(pr)
}

// QueryOrderProducts queries the "order_products" edge of the Product entity.
func (pr *Product) QueryOrderProducts() *OrderProductQuery {
	return (&ProductClient{config: pr.config}).QueryOrderProducts(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", slug=")
	builder.WriteString(pr.Slug)
	builder.WriteString(", meta_title=")
	builder.WriteString(pr.MetaTitle)
	builder.WriteString(", meta_description=")
	builder.WriteString(pr.MetaDescription)
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}