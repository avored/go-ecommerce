// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/avored/go-ecommerce/ent/address"
	"github.com/avored/go-ecommerce/ent/customer"
)

// Address is the model entity for the Address schema.
type Address struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// IsDefault holds the value of the "is_default" field.
	IsDefault bool `json:"is_default,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Address1 holds the value of the "address1" field.
	Address1 string `json:"address1,omitempty"`
	// Address2 holds the value of the "address2" field.
	Address2 string `json:"address2,omitempty"`
	// CompanyName holds the value of the "company_name" field.
	CompanyName string `json:"company_name,omitempty"`
	// Suburb holds the value of the "suburb" field.
	Suburb string `json:"suburb,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Postcode holds the value of the "postcode" field.
	Postcode string `json:"postcode,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// CountryID holds the value of the "country_id" field.
	CountryID int `json:"country_id,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID int `json:"customer_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AddressQuery when eager-loading is set.
	Edges AddressEdges `json:"edges"`
}

// AddressEdges holds the relations/edges for other nodes in the graph.
type AddressEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// ShippingOrders holds the value of the shipping_orders edge.
	ShippingOrders []*Order `json:"shipping_orders,omitempty"`
	// BillingOrders holds the value of the billing_orders edge.
	BillingOrders []*Order `json:"billing_orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AddressEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.Customer == nil {
			// The edge customer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// ShippingOrdersOrErr returns the ShippingOrders value or an error if the edge
// was not loaded in eager-loading.
func (e AddressEdges) ShippingOrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[1] {
		return e.ShippingOrders, nil
	}
	return nil, &NotLoadedError{edge: "shipping_orders"}
}

// BillingOrdersOrErr returns the BillingOrders value or an error if the edge
// was not loaded in eager-loading.
func (e AddressEdges) BillingOrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[2] {
		return e.BillingOrders, nil
	}
	return nil, &NotLoadedError{edge: "billing_orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Address) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case address.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case address.FieldID, address.FieldCountryID, address.FieldCustomerID:
			values[i] = new(sql.NullInt64)
		case address.FieldType, address.FieldFirstName, address.FieldLastName, address.FieldAddress1, address.FieldAddress2, address.FieldCompanyName, address.FieldSuburb, address.FieldCity, address.FieldState, address.FieldPostcode, address.FieldPhone:
			values[i] = new(sql.NullString)
		case address.FieldCreatedAt, address.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Address", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Address fields.
func (a *Address) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case address.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case address.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case address.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case address.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = value.String
			}
		case address.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				a.IsDefault = value.Bool
			}
		case address.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				a.FirstName = value.String
			}
		case address.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				a.LastName = value.String
			}
		case address.FieldAddress1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address1", values[i])
			} else if value.Valid {
				a.Address1 = value.String
			}
		case address.FieldAddress2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address2", values[i])
			} else if value.Valid {
				a.Address2 = value.String
			}
		case address.FieldCompanyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company_name", values[i])
			} else if value.Valid {
				a.CompanyName = value.String
			}
		case address.FieldSuburb:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field suburb", values[i])
			} else if value.Valid {
				a.Suburb = value.String
			}
		case address.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				a.City = value.String
			}
		case address.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				a.State = value.String
			}
		case address.FieldPostcode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postcode", values[i])
			} else if value.Valid {
				a.Postcode = value.String
			}
		case address.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				a.Phone = value.String
			}
		case address.FieldCountryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field country_id", values[i])
			} else if value.Valid {
				a.CountryID = int(value.Int64)
			}
		case address.FieldCustomerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				a.CustomerID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCustomer queries the "customer" edge of the Address entity.
func (a *Address) QueryCustomer() *CustomerQuery {
	return (&AddressClient{config: a.config}).QueryCustomer(a)
}

// QueryShippingOrders queries the "shipping_orders" edge of the Address entity.
func (a *Address) QueryShippingOrders() *OrderQuery {
	return (&AddressClient{config: a.config}).QueryShippingOrders(a)
}

// QueryBillingOrders queries the "billing_orders" edge of the Address entity.
func (a *Address) QueryBillingOrders() *OrderQuery {
	return (&AddressClient{config: a.config}).QueryBillingOrders(a)
}

// Update returns a builder for updating this Address.
// Note that you need to call Address.Unwrap() before calling this method if this Address
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Address) Update() *AddressUpdateOne {
	return (&AddressClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Address entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Address) Unwrap() *Address {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Address is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Address) String() string {
	var builder strings.Builder
	builder.WriteString("Address(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", type=")
	builder.WriteString(a.Type)
	builder.WriteString(", is_default=")
	builder.WriteString(fmt.Sprintf("%v", a.IsDefault))
	builder.WriteString(", first_name=")
	builder.WriteString(a.FirstName)
	builder.WriteString(", last_name=")
	builder.WriteString(a.LastName)
	builder.WriteString(", address1=")
	builder.WriteString(a.Address1)
	builder.WriteString(", address2=")
	builder.WriteString(a.Address2)
	builder.WriteString(", company_name=")
	builder.WriteString(a.CompanyName)
	builder.WriteString(", suburb=")
	builder.WriteString(a.Suburb)
	builder.WriteString(", city=")
	builder.WriteString(a.City)
	builder.WriteString(", state=")
	builder.WriteString(a.State)
	builder.WriteString(", postcode=")
	builder.WriteString(a.Postcode)
	builder.WriteString(", phone=")
	builder.WriteString(a.Phone)
	builder.WriteString(", country_id=")
	builder.WriteString(fmt.Sprintf("%v", a.CountryID))
	builder.WriteString(", customer_id=")
	builder.WriteString(fmt.Sprintf("%v", a.CustomerID))
	builder.WriteByte(')')
	return builder.String()
}

// Addresses is a parsable slice of Address.
type Addresses []*Address

func (a Addresses) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}