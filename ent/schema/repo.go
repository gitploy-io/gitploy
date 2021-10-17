package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Repo holds the schema definition for the Repo entity.
type Repo struct {
	ent.Schema
}

// Fields of the Repo.
func (Repo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("namespace"),
		field.String("name"),
		field.String("description"),
		field.String("config_path").
			Default("deploy.yml"),
		// Activated repo has the webhook to update the deployment status.
		field.Bool("active").
			Default(false),
		field.Int64("webhook_id").
			Optional(),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),
		// Denormalization to sort with deployment.
		field.Time("latest_deployed_at").
			Optional(),
	}
}

// Edges of the Repo.
func (Repo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("perms", Perm.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("deployments", Deployment.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("callback", Callback.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("locks", Lock.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("deployment_statistics", DeploymentStatistics.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Repo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "name").
			Unique(),
		index.Fields("name"),
	}
}
