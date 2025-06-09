package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("expires_at"),
		field.String("refresh_token"),
		field.String("ip_address"),
		field.String("user_agent"),
	}
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("sessions"),
	}
}
