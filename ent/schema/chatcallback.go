package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatCallback holds the schema definition for the ChatCallback entity.
type ChatCallback struct {
	ent.Schema
}

// Fields of the ChatCallback.
func (ChatCallback) Fields() []ent.Field {
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
			),
		field.Bool("is_opened").
			Default(true),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("chat_user_id"),
		field.String("repo_id"),
	}
}

// Edges of the ChatCallback.
func (ChatCallback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chat_user", ChatUser.Type).
			Ref("chat_callback").
			Field("chat_user_id").
			Unique().
			Required(),
		edge.From("repo", Repo.Type).
			Ref("chat_callback").
			Field("repo_id").
			Unique().
			Required(),
	}
}
