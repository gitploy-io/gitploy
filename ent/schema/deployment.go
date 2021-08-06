package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Int("number"),
		field.Enum("type").
			Values(
				"commit",
				"branch",
				"tag",
			).
			Default("commit"),
		field.String("ref"),
		field.String("sha"),
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
		// UID is determined from SCM.
		// The waiting status can not have UID.
		field.Int64("uid").
			Optional(),
		field.Bool("is_rollback").
			Default(false),
		field.Bool("is_approval_enabled").
			Default(false),
		field.Int("required_approval_count").
			Default(0),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		// Edges
		field.String("user_id"),
		field.String("repo_id"),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("deployments").
			Field("user_id").
			Unique().
			Required(),
		edge.From("repo", Repo.Type).
			Ref("deployments").
			Field("repo_id").
			Unique().
			Required(),
		edge.To("approvals", Approval.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("deployment_statuses", DeploymentStatus.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Deployment) Indexes() []ent.Index {
	return []ent.Index{
		// Basically deployments are ordered by created_at field.
		index.Fields("repo_id", "env", "created_at"),
		index.Fields("repo_id", "created_at"),
		// The deployment number is unique for the repo.
		index.Fields("number", "repo_id").
			Unique(),
		// Find by UID when the hook is coming.
		index.Fields("uid"),
	}
}
