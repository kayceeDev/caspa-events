package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Unique().Default(uuid.New).Immutable().StructTag(`json:"uuid"`),
		field.String("first_name").NotEmpty().StructTag(`json:"first_name"`),
		field.String("last_name").NotEmpty().StructTag(`json:"last_name"`),
		field.String("email").NotEmpty().Unique().StructTag(`json:"email"`),
		field.String("password").NotEmpty().StructTag(`json:"password"`).Optional(),
		field.String("phone").NotEmpty().Unique().StructTag(`json:"phone"`).Optional(),
		field.Bool("email_verified").Default(false).StructTag(`json:"email_verified"`),
		field.Bool("phone_verified").Default(false).StructTag(`json:"phone_verified"`),
		field.Time("email_verified_at").Nillable().Optional().StructTag(`json:"phone_verified"`),
		field.Time("phone_verified_at").Nillable().Optional().StructTag(`json:"phone_verified"`),
		field.Bool("disabled").Optional().Default(false).StructTag(`json:"disabled"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("event", Event.Type),
	}
}
