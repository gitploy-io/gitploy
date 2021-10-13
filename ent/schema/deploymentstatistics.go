package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DeploymentStatistics holds the schema definition for the DeploymentStatistics entity.
type DeploymentStatistics struct {
	ent.Schema
}

// Fields of the DeploymentStatistics.
func (DeploymentStatistics) Fields() []ent.Field {
	return []ent.Field{
		field.String("env"),
		field.Int("count").
			Default(1),
		field.Int("rollback_count").
			Default(0),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int64("repo_id"),
	}
}

// Edges of the DeploymentStatistics.
func (DeploymentStatistics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repo", Repo.Type).
			Ref("deployment_statistics").
			Field("repo_id").
			Unique().
			Required(),
	}
}

func (DeploymentStatistics) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "env").
			Unique(),
		// The collector searches updated records only.
		index.Fields("updated_at"),
	}
}
