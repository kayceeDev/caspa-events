package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type Ticket struct {
	ent.Schema
}

func (Ticket) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Unique().Default(uuid.New).Immutable().StructTag(`json:"uuid"`),
		field.String("name").NotEmpty().StructTag(`json:"name"`),
		field.String("description").NotEmpty().StructTag(`json:"description"`),
		field.Float("price").Positive().StructTag(`json:"price"`),
		field.Int("quantity").Positive().StructTag(`json:"quantity"`),
		field.Int("quantity_sold").Positive().Default(0).StructTag(`json:"quantity_sold"`),
		field.Time("sale_start_date").StructTag(`json:"sale_start_date"`),
		field.Time("sale_end_date").Nillable().Optional().StructTag(`json:"sale_end_date"`),
		field.String("event_id").StructTag(`json:"event_id"`),
		field.String("ticket_type").StructTag(`json:"ticket_type"`).MaxLen(50).NotEmpty(),
		field.Bool("is_active").Default(true).StructTag(`json:"is_active"`),
		field.Bool("is_refundable").Default(false).StructTag(`json:"is_refundable"`),
	}
}
