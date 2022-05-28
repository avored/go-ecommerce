// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/avored/go-ecommerce/ent/adminuser"
)

// AdminUser is the model entity for the AdminUser schema.
type AdminUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AdminUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case adminuser.FieldID:
			values[i] = new(sql.NullInt64)
		case adminuser.FieldFirstName, adminuser.FieldLastName, adminuser.FieldEmail, adminuser.FieldPassword:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AdminUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdminUser fields.
func (au *AdminUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case adminuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			au.ID = int(value.Int64)
		case adminuser.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				au.FirstName = value.String
			}
		case adminuser.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				au.LastName = value.String
			}
		case adminuser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				au.Email = value.String
			}
		case adminuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				au.Password = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AdminUser.
// Note that you need to call AdminUser.Unwrap() before calling this method if this AdminUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (au *AdminUser) Update() *AdminUserUpdateOne {
	return (&AdminUserClient{config: au.config}).UpdateOne(au)
}

// Unwrap unwraps the AdminUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (au *AdminUser) Unwrap() *AdminUser {
	tx, ok := au.config.driver.(*txDriver)
	if !ok {
		panic("ent: AdminUser is not a transactional entity")
	}
	au.config.driver = tx.drv
	return au
}

// String implements the fmt.Stringer.
func (au *AdminUser) String() string {
	var builder strings.Builder
	builder.WriteString("AdminUser(")
	builder.WriteString(fmt.Sprintf("id=%v", au.ID))
	builder.WriteString(", first_name=")
	builder.WriteString(au.FirstName)
	builder.WriteString(", last_name=")
	builder.WriteString(au.LastName)
	builder.WriteString(", email=")
	builder.WriteString(au.Email)
	builder.WriteString(", password=")
	builder.WriteString(au.Password)
	builder.WriteByte(')')
	return builder.String()
}

// AdminUsers is a parsable slice of AdminUser.
type AdminUsers []*AdminUser

func (au AdminUsers) config(cfg config) {
	for _i := range au {
		au[_i].config = cfg
	}
}