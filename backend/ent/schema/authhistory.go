package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type AuthHistory struct {
	ent.Schema
}

func (AuthHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("method"),
		field.Time("authenticated_at"),
		field.String("ip_address"),
		field.String("user_agent"),
	}
}

func (AuthHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("auth_histories").Unique(),
	}
}
