package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type Event struct {
	ent.Schema
}

// Mixin of the Event.
func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Unique().Default(uuid.New).Immutable().StructTag(`json:"uuid"`),
		field.String("title").NotEmpty().StructTag(`json:"title"`),
		field.String("description").NotEmpty().StructTag(`json:"description"`),
		field.Time("start_date").StructTag(`json:"start_date"`),
		field.Time("end_date").StructTag(`json:"end_date"`),
		field.String("event_type").StructTag(`json:"event_type"`).MaxLen(50).NotEmpty(),
		field.Enum("status").Values("draft", "published", "cancelled", "ended").Default("draft").StructTag(`json:"status"`),
		field.Bool("is_public").Default(false).StructTag(`json:"is_public"`),
		field.Bool("is_paid").Default(false).StructTag(`json:"is_paid"`),
		field.Int("max_participants").Default(0).StructTag(`json:"max_participants"`),
		field.Time("registration_deadline").Nillable().Optional().StructTag(`json:"registration_deadline"`),
		field.String("cover_photo_id").Nillable().Optional().StructTag(`json:"cover_photo_id"`),
		field.Int("organizer_id").StructTag(`json:"organizer_id"`),
		field.Int("location_id").Nillable().Optional().StructTag(`json:"location_id"`),
	}
}

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("event").
			Unique().
			Required().
			Field("organizer_id"),
		edge.From("location", Location.Type).Ref("event").
			Unique().Field("location_id"),
		edge.To("guest", User.Type),
		edge.To("ticket", Ticket.Type).
			StructTag(`json:"tickets"`),
		// edge.To("ratings", Rating.Type).
		// 	StructTag(`json:"ratings"`),
	}

}
