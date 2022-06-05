// Code generated by entc, DO NOT EDIT.

package cart

import (
	"time"
)

const (
	// Label holds the string label denoting the cart type in the database.
	Label = "cart"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSessionID holds the string denoting the session_id field in the database.
	FieldSessionID = "session_id"
	// FieldFullAmount holds the string denoting the full_amount field in the database.
	FieldFullAmount = "full_amount"
	// EdgeCartItems holds the string denoting the cart_items edge name in mutations.
	EdgeCartItems = "cart_items"
	// Table holds the table name of the cart in the database.
	Table = "carts"
	// CartItemsTable is the table that holds the cart_items relation/edge.
	CartItemsTable = "cart_items"
	// CartItemsInverseTable is the table name for the CartItem entity.
	// It exists in this package in order to avoid circular dependency with the "cartitem" package.
	CartItemsInverseTable = "cart_items"
	// CartItemsColumn is the table column denoting the cart_items relation/edge.
	CartItemsColumn = "cart_id"
)

// Columns holds all SQL columns for cart fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSessionID,
	FieldFullAmount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
