package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field {
        field.String("first_name"),
        field.String("last_name"),
        field.String("email").Unique(),
        field.String("password"),
    }
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("addresses", Address.Type),
		edge.To("orders", Order.Type),
	}
}


func (Customer) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}
