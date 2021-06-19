package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ChatUser holds the schema definition for the ChatUser entity.
type ChatUser struct {
	ent.Schema
}

// Fields of the ChatUser.
func (ChatUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("token"),
		field.String("refresh"),
		field.Time("expiry"),
		field.String("bot_token"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("user_id"),
	}
}

// Edges of the ChatUser.
func (ChatUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat_callback", ChatCallback.Type),
		edge.From("user", User.Type).
			Ref("chat_user").
			Field("user_id").
			Unique().
			Required(),
	}
}

func (ChatUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
