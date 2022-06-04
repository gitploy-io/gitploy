package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("kind").
			Values(
				"deployment_status",
				"review",
			),
		field.Enum("type").
			Values(
				"created",
				"updated",
				"deleted",
			),
		field.Time("created_at").
			Default(nowUTC),
		field.Int("deployment_status_id").
			Optional(),
		field.Int("review_id").
			Optional(),
		// This field is filled when the type is 'deleted'.
		field.Int("deleted_id").
			Optional(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment_status", DeploymentStatus.Type).
			Ref("event").
			Field("deployment_status_id").
			Unique(),
		edge.From("review", Review.Type).
			Ref("event").
			Field("review_id").
			Unique(),
		edge.To("notification_record", NotificationRecord.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}).
			Unique(),
	}
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}
