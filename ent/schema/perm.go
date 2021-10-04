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
		field.Time("synced_at"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		// Edges
		field.Int64("user_id"),
		field.Int64("repo_id"),
	}
}

// Edges of the Perm.
func (Perm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("perms").
			Field("user_id").
			Unique().
			Required(),
		edge.From("repo", Repo.Type).
			Ref("perms").
			Field("repo_id").
			Unique().
			Required(),
	}
}

func (Perm) Indexes() []ent.Index {
	return []ent.Index{
		// Find the perm for the repository.
		index.Fields("repo_id", "user_id"),
		// Delete staled perms after synchronization
		index.Fields("user_id", "synced_at"),
	}
}
