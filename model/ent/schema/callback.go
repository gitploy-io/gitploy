package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Callback holds the schema definition for the Callback entity.
type Callback struct {
	ent.Schema
}

// Fields of the Callback.
func (Callback) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").
			Immutable().
			Unique().
			DefaultFunc(generateHash).
			Sensitive(),
		field.Enum("type").
			Values(
				"deploy",
				"rollback",
				"lock",
				"unlock",
			),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),
		field.Int64("repo_id"),
	}
}

// Edges of the Callback.
func (Callback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repo", Repo.Type).
			Ref("callback").
			Field("repo_id").
			Unique().
			Required(),
	}
}
