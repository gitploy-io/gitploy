package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		// Edges
		field.String("user_id").
			Optional(),
		field.String("repo_id").
			Optional(),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("deployments").
			Field("user_id").
			Unique(),
		edge.From("repo", Repo.Type).
			Ref("deployments").
			Field("repo_id").
			Unique(),
	}
}

func (Deployment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("repo_id"),
		index.Fields("repo_id", "env", "status", "created_at"),
		index.Fields("repo_id", "env", "created_at"),
		index.Fields("repo_id", "created_at"),
	}
}
