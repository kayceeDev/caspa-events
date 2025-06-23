package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Add any mixins here if needed, e.g., for timestamps or soft deletes.
		mixin.Time{},
	}
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Unique().Default(uuid.New).Immutable().StructTag(`json:"uuid"`),
		field.String("name").MaxLen(255).NotEmpty(),
		field.String("address").Optional(),
		field.String("city").MaxLen(100).Optional(),
		field.String("state").MaxLen(100).Optional(),
		field.String("country").MaxLen(100).Optional(),
		field.String("postal_code").MaxLen(20).Optional(),
		field.Float("latitude").Optional(),
		field.Float("longitude").Optional(),
		field.Int("capacity").Optional(),
	}
}

func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		// Define edges if this entity has relationships with other entities.
		// For example, if Location is related to Event:
		edge.To("event", Event.Type),
	}
}
