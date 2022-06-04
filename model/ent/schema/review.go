package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values(
				"pending",
				"rejected",
				"approved",
			).
			Default("pending"),
		field.Text("comment").
			Optional(),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),
		// Edges
		field.Int64("user_id").
			Optional(),
		field.Int("deployment_id"),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("reviews").
			Field("user_id").
			Unique().
			Required(),
		edge.From("deployment", Deployment.Type).
			Ref("reviews").
			Field("deployment_id").
			Unique().
			Required(),
		edge.To("event", Event.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
