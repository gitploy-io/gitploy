package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Deployment holds the schema definition for the Deployment entity.
type Deployment struct {
	ent.Schema
}

// Fields of the Deployment.
func (Deployment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("uid").
			Optional(),
		field.Enum("type").
			Values(
				"commit",
				"branch",
				"tag",
			).
			Default("commit"),
		field.String("ref"),
		field.String("sha").
			Optional(),
		field.String("env"),
		field.Enum("status").
			Values(
				"waiting",
				"created",
				"running",
				"success",
				"failure",
			).
			Default("waiting"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("deployments").
			Unique(),
		edge.From("repo", Repo.Type).
			Ref("deployments").
			Unique(),
	}
}
