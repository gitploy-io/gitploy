package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Lock holds the schema definition for the Lock entity.
type Lock struct {
	ent.Schema
}

// Fields of the Lock.
func (Lock) Fields() []ent.Field {
	return []ent.Field{
		field.String("env"),
		field.Time("created_at").
			Default(time.Now),
		// Edges
		field.String("user_id"),
		field.String("repo_id"),
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
	return []ent.Index{}
}
