package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field {
		field.String("payment_method"),
		field.String("shipping_method"),
		field.Int("shipping_address_id").Optional(),
		field.Int("billing_address_id").Optional(),
		field.Int("customer_id").Optional(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("shipping_address", Address.Type).Ref("shipping_orders").Unique().Field("shipping_address_id"),
		edge.From("billing_address", Address.Type).Ref("billing_orders").Unique().Field("billing_address_id"),
		edge.From("customer_order", Customer.Type).Ref("orders").Unique().Field("customer_id"),
		edge.To("order_products", OrderProduct.Type),
	}
}

func (Order) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}
