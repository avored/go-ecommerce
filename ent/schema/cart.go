package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Fields of the Cart.
func (Cart) Fields() []ent.Field {
	return []ent.Field {
		// field.Uint("customer_id").Optional().Nillable(),
		field.String("session_id"),
		field.Float("full_amount").
            SchemaType(map[string]string{
                dialect.MySQL:    "decimal(18, 8)",
            }),
	}
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("cart_items", CartItem.Type),
	}
}



func (Cart) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}
