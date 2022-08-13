package schema


import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field {
        field.String("name"),
        field.String("identifier"),
        field.String("meta_title"),
        field.Text("meta_description"),
    }
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}


func (Category) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}
