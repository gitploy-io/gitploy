package schema

import (
	"time"

	"entgo.io/ent"
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

// Edges of the DeploymentStatistics.
func (DeploymentStatistics) Edges() []ent.Edge {
	return nil
}

func (DeploymentStatistics) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "name", "env").
			Unique(),
		// The collector searches updated records only.
		index.Fields("updated_at"),
	}
}
