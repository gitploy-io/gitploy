package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NotificationRecord holds the schema definition for the NotificationRecord entity.
type NotificationRecord struct {
	ent.Schema
}

// Fields of the NotificationRecord.
func (NotificationRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("event_id"),
	}
}

// Edges of the NotificationRecord.
func (NotificationRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("notification_record").
			Field("event_id").
			Unique().
			Required(),
	}
}
