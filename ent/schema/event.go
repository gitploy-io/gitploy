package schema

import (
	"time"

	"entgo.io/ent"
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
				"deployment",
				"approval",
			),
		field.Enum("type").
			Values(
				"created",
				"updated",
				"removed",
			),
		field.Time("created_at").
			Default(time.Now),
		field.Int("deployment_id").
			Optional(),
		field.Int("approval_id").
			Optional(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment", Deployment.Type).
			Ref("event").
			Field("deployment_id").
			Unique(),
		edge.From("approval", Approval.Type).
			Ref("event").
			Field("approval_id").
			Unique(),
		edge.To("notification_record", NotificationRecord.Type).
			Unique(),
	}
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}
