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
				"deployment_created",
				"deployment_updated",
				"approval_requested",
				"approval_responded",
			),
		// Denormalization from repository, deployment.
		field.String("repo_namespace"),
		field.String("repo_name"),
		field.Int("deployment_number"),
		field.String("deployment_type"),
		field.String("deployment_ref"),
		field.String("deployment_env"),
		field.String("deployment_status"),
		field.String("deployment_login"),
		field.String("approval_status").
			Optional(),
		field.String("approval_login").
			Optional(),
		// The notified field means it is notified by Chat or browser,
		// in meanwhile The checked field means the user has checked or not.
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
	}
}

func (Notification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("user_id", "created_at"),
	}
}
