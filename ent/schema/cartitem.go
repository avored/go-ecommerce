package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CartItem holds the schema definition for the CartItem entity.
type CartItem struct {
	ent.Schema
}

// Fields of the CartItem.
func (CartItem) Fields() []ent.Field {
	return []ent.Field {
		field.String("name"),
		field.Float("qty"),
		field.Int("cart_id").Optional(),
		field.Int("product_id").Optional(),
		field.Float("item_price").
            SchemaType(map[string]string{
                dialect.MySQL:    "decimal(18, 8)",
            }),
	}
	
}

// Edges of the CartItem.
func (CartItem) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("cart", Cart.Type).Ref("cart_items").Unique().Field("cart_id"),
		edge.From("product", Product.Type).Ref("cart_items").Unique().Field("product_id"),
	}
}
