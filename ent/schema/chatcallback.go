package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ChatCallback holds the schema definition for the ChatCallback entity.
type ChatCallback struct {
	ent.Schema
}

// Fields of the ChatCallback.
func (ChatCallback) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("state"),
		field.Enum("type").
			Values(
				"deploy",
				"rollback",
			),
		field.Bool("is_opened"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the ChatCallback.
func (ChatCallback) Edges() []ent.Edge {
	return nil
}
