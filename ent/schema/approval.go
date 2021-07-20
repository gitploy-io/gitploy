package schema

import (
	"time"

	"entgo.io/ent"
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
		field.Bool("is_approved").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		// Edges
		field.String("user_id"),
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
	}
}

func (Approval) Index() []ent.Index {
	return []ent.Index{
		index.Fields("deployment_id", "user_id").
			Unique(),
	}
}
