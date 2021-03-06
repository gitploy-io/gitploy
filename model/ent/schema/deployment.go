package schema

import (
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
		field.String("env"),
		field.String("ref"),
		field.JSON("dynamic_payload", map[string]interface{}{}).
			Optional(),
		field.Enum("status").
			Values(
				"waiting",
				"created",
				"queued",
				"running",
				"success",
				"failure",
				"canceled",
			).
			Default("waiting"),
		// UID, SHA, and HTLM URL are returned after
		// the remote deployment is created.
		field.Int64("uid").
			Optional(),
		field.String("sha").
			Optional(),
		field.String("html_url").
			MaxLen(2000).
			Optional(),
		field.Bool("production_environment").
			Default(false),
		field.Bool("is_rollback").
			Default(false),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),
		// Edges
		field.Int64("user_id").
			Optional(),
		field.Int64("repo_id"),

		// Deprecated fields.
		field.Bool("is_approval_enabled").
			Optional().
			Nillable(),
		field.Int("required_approval_count").
			Optional().
			Nillable(),
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
			Unique().
			Required(),
		edge.To("reviews", Review.Type).
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
		// Find the latest succeed deployment.
		index.Fields("repo_id", "env", "status", "updated_at"),
		// List deployments by the env.
		index.Fields("repo_id", "env", "created_at"),
		index.Fields("repo_id", "created_at"),
		// The deployment number is unique for the repo.
		index.Fields("repo_id", "number").
			Unique(),
		// Find by UID when the hook is coming.
		index.Fields("uid"),
		// List inactive deployments for 30 minutes.
		// Or search deployments that were created between the start time and the end time.
		index.Fields("created_at", "status"),
	}
}
