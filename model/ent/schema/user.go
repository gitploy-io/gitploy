package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("login").
			Unique(),
		field.String("avatar"),
		field.Bool("admin").
			Default(false),
		field.String("token").
			Sensitive(),
		field.String("refresh").
			Sensitive(),
		field.Time("expiry"),
		field.String("hash").
			Immutable().
			Unique().
			DefaultFunc(generateHash).
			Sensitive(),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat_user", ChatUser.Type).
			Unique().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("perms", Perm.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		// In order to record the history, when a user is deleted, it is filled with a null value.
		edge.To("deployments", Deployment.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
		edge.To("reviews", Review.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
		edge.To("locks", Lock.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
		edge.To("repos", Repo.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
	}
}
