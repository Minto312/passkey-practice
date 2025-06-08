package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"time"
)

type Passkey struct {
	ent.Schema
}

func (Passkey) Fields() []ent.Field {
	return []ent.Field{
		field.String("credential_id"),
		field.String("public_key"),
		field.String("device_name"),
		field.Time("created_at").Default(time.Now),
		field.Time("last_used_at").Default(time.Now),
	}
}

func (Passkey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("passkeys"),
	}
} 