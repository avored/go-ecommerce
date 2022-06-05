package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderProduct holds the schema definition for the OrderProduct entity.
type OrderProduct struct {
	ent.Schema
}

// Fields of the OrderProduct.
func (OrderProduct) Fields() []ent.Field {
	return []ent.Field {
		field.Int("order_id").Optional(),
		field.Int("product_id").Optional(),
		field.Float("qty"),
		field.Float("amount").
            SchemaType(map[string]string{
                dialect.MySQL:    "decimal(18, 8)",
            }),
	}
}

// Edges of the OrderProduct.
func (OrderProduct) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("order_product_order", Order.Type).Ref("order_products").Unique().Field("order_id"),
		edge.From("order_product_product", Product.Type).Ref("order_products").Unique().Field("product_id"),

	}
}


func (OrderProduct) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}
