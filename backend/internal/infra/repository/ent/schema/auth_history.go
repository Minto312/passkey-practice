package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AuthHistory holds the schema definition for the AuthHistory entity.
type AuthHistory struct {
	ent.Schema
}

// Fields of the AuthHistory.
func (AuthHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("method").
			Values("password", "passkey"),
		field.Time("authenticated_at"),
		field.String("ip_address"),
		field.String("user_agent"),
		field.Int("user_id"),
	}
}

// Edges of the AuthHistory.
func (AuthHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("auth_histories").
			Field("user_id").
			Unique().
			Required(),
	}
}
