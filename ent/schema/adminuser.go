package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AdminUser holds the schema definition for the AdminUser entity.
type AdminUser struct {
	ent.Schema
}

// Fields of the AdminUser.
func (AdminUser) Fields() []ent.Field {
	return []ent.Field {
        field.String("first_name"),
        field.String("last_name"),
        field.String("email").Unique(),
        field.String("password"),
    }
}

// Edges of the AdminUser.
func (AdminUser) Edges() []ent.Edge {
	return nil
}
