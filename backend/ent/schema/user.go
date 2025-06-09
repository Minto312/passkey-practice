package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password_hash").NotEmpty(),
		field.String("display_name"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("passkeys", Passkey{}),
		edge.To("auth_histories", AuthHistory{}),
		edge.To("sessions", Session{}),
	}
}
