package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		// user_id is null if the user is deleted.
		field.String("user_id").
			Optional(),
		field.Int("deployment_id"),
	}
}

// Edges of the Approval.
func (Approval) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("approvals").
			Field("user_id").
			Unique(),
		edge.From("deployment", Deployment.Type).
			Ref("approvals").
			Field("deployment_id").
			Unique().
			Required(),
	}
}

func (Approval) Index() []ent.Index {
	return []ent.Index{}
}
