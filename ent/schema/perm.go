package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Perm holds the schema definition for the Perm entity.
type Perm struct {
	ent.Schema
}

// Fields of the Perm.
func (Perm) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("repo_perm").
			Values(
				"read",
				"write",
				"admin",
			).
			Default("read"),
		field.Time("synced_at").
			Optional(),
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

// Edges of the Perm.
func (Perm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("perms").
			Field("user_id").
			Unique(),
		edge.From("repo", Repo.Type).
			Ref("perms").
			Field("repo_id").
			Unique(),
	}
}

func (Perm) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("repo_id"),
	}
}
