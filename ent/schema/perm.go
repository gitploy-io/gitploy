package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
	}
}

// Edges of the Perm.
func (Perm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("perms").
			Unique(),
		edge.From("repo", Repo.Type).
			Ref("perms").
			Unique(),
	}
}
