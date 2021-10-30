package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values(
				"approved",
				"rejected",
			),
		field.String("comment"),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),
		// Edges
		field.Int64("user_id"),
		field.Int("deployment_id"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("comments").
			Field("user_id").
			Unique().
			Required(),
		edge.From("deployment", Deployment.Type).
			Ref("comments").
			Field("deployment_id").
			Unique().
			Required(),
	}
}
