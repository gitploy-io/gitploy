package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DeploymentCount holds the schema definition for the DeploymentCount entity.
// It is the kind of statistics table to count deployments.
type DeploymentCount struct {
	ent.Schema
}

// Fields of the DeploymentCount.
func (DeploymentCount) Fields() []ent.Field {
	return []ent.Field{
		field.String("namespace"),
		field.String("name"),
		field.String("env"),
		field.Int("count").
			Default(1),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the DeploymentCount.
func (DeploymentCount) Edges() []ent.Edge {
	return nil
}

func (DeploymentCount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "name", "env").
			Unique(),
		// The collector searches updated records only.
		index.Fields("updated_at"),
	}
}
