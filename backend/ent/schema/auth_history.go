package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuthHistory holds the schema definition for the AuthHistory entity.
type AuthHistory struct {
	ent.Schema
}

// Fields of the AuthHistory.
func (AuthHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("method"),
		field.Time("authenticated_at").
			Default(time.Now).
			Immutable(),
		field.String("ip_address"),
		field.String("user_agent"),
	}
}

// Edges of the AuthHistory.
func (AuthHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("auth_histories").
			Unique().
			Required(),
	}
}
