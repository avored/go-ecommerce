package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field {
		field.String("name"),
		field.String("slug"),
		field.String("meta_title"),
		field.String("meta_description"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("categories", Category.Type),
		edge.To("cart_items", CartItem.Type),
		edge.To("order_products", OrderProduct.Type),
	}
}

func (Product) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}
