package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatUser holds the schema definition for the ChatUser entity.
type ChatUser struct {
	ent.Schema
}

// Fields of the ChatUser.
func (ChatUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("token").
			Sensitive(),
		field.String("refresh").
			Sensitive(),
		field.Time("expiry"),
		field.String("bot_token").
			Sensitive(),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),
		field.Int64("user_id"),
	}
}

// Edges of the ChatUser.
func (ChatUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("chat_user").
			Field("user_id").
			Unique().
			Required(),
	}
}
