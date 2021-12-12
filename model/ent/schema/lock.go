package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Lock holds the schema definition for the Lock entity.
type Lock struct {
	ent.Schema
}

// Fields of the Lock.
func (Lock) Fields() []ent.Field {
	return []ent.Field{
		field.String("env"),
		field.Time("expired_at").
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(nowUTC),
		// Edges
		field.Int64("user_id"),
		field.Int64("repo_id"),
	}
}

// Edges of the Lock.
func (Lock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("locks").
			Field("user_id").
			Unique().
			Required(),
		edge.From("repo", Repo.Type).
			Ref("locks").
			Field("repo_id").
			Unique().
			Required(),
	}
}

func (Lock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "env").
			Unique(),
	}
}
