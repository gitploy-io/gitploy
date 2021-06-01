package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Repo holds the schema definition for the Repo entity.
type Repo struct {
	ent.Schema
}

// Fields of the Repo.
func (Repo) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("namespace"),
		field.String("name"),
		field.String("description").
			Optional(),
		field.String("config_path").
			Default("deploy.yml"),
		field.Time("synced_at").
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		// Denormalization to sort with deployment.
		field.Time("latest_deployed_at").
			Optional(),
	}
}

// Edges of the Repo.
func (Repo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("perms", Perm.Type),
		edge.To("deployments", Deployment.Type),
	}
}
