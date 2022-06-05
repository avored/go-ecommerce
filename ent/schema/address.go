package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field {
		field.String("type"),
		field.Bool("is_default").Default(false),
		field.String("first_name"),
		field.String("last_name"),
		field.String("address1"),
		field.String("address2"),
		field.String("company_name"),
		field.String("suburb"),
		field.String("city"),
		field.String("state"),
		field.String("postcode"),
		field.String("phone"),
		field.Int("country_id"),
		field.Int("customer_id").Optional(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("customer", Customer.Type).Ref("addresses").Unique().Field("customer_id"),
		edge.To("shipping_orders", Order.Type),
		edge.To("billing_orders", Order.Type),
	}
}


func (Address) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}
