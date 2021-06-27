package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values(
				"deployment",
			).
			Default("deployment"),
		// notified means it is notified by Chat or browser,
		// in meanwhile checked means the status is checked directly or not.
		field.Bool("notified").
			Default(false),
		field.Bool("checked").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("user_id"),
		field.Int("deployment_id").
			Optional(),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("notification").
			Field("user_id").
			Required().
			Unique(),
		edge.From("deployment", Deployment.Type).
			Ref("notifications").
			Field("deployment_id").
			Unique(),
	}
}

func (Notification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("created_at"),
		index.Fields("user_id", "created_at"),
	}
}
