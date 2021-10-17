package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Approval holds the schema definition for the Approval entity.
type Approval struct {
	ent.Schema
}

// Fields of the Approval.
func (Approval) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values(
				"pending",
				"declined",
				"approved",
			).
			Default("pending"),
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

// Edges of the Approval.
func (Approval) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("approvals").
			Field("user_id").
			Unique().
			Required(),
		edge.From("deployment", Deployment.Type).
			Ref("approvals").
			Field("deployment_id").
			Unique().
			Required(),
		edge.To("event", Event.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Approval) Index() []ent.Index {
	return []ent.Index{
		index.Fields("deployment_id", "user_id").
			Unique(),
	}
}
