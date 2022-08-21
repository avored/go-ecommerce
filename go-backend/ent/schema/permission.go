package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}


// Fields of the Role.
func (Permission) Fields() []ent.Field {
	return []ent.Field {
        field.String("name"),
        field.String("identifier"),
        field.Text("description"),
    }
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}

func (Permission) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}
