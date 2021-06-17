package schema

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"entgo.io/ent"
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
		field.String("id"),
		field.String("login").
			Unique(),
		field.String("avatar").
			Optional(),
		field.Bool("admin").
			Default(false),
		field.String("token"),
		field.String("refresh"),
		field.Time("expiry"),
		field.String("hash").
			Immutable().
			Unique().
			DefaultFunc(func() string {
				b := make([]byte, 16)
				rand.Read(b)
				return base64.URLEncoding.EncodeToString(b)
			}),
		field.Time("synced_at").
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat_user", ChatUser.Type).
			Unique(),
		edge.To("perms", Perm.Type),
		edge.To("deployments", Deployment.Type),
	}
}
