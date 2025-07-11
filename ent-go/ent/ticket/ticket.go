// Code generated by ent, DO NOT EDIT.

package ticket

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ticket type in the database.
	Label = "ticket"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldQuantitySold holds the string denoting the quantity_sold field in the database.
	FieldQuantitySold = "quantity_sold"
	// FieldSaleStartDate holds the string denoting the sale_start_date field in the database.
	FieldSaleStartDate = "sale_start_date"
	// FieldSaleEndDate holds the string denoting the sale_end_date field in the database.
	FieldSaleEndDate = "sale_end_date"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldTicketType holds the string denoting the ticket_type field in the database.
	FieldTicketType = "ticket_type"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldIsRefundable holds the string denoting the is_refundable field in the database.
	FieldIsRefundable = "is_refundable"
	// Table holds the table name of the ticket in the database.
	Table = "tickets"
)

// Columns holds all SQL columns for ticket fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUUID,
	FieldName,
	FieldDescription,
	FieldPrice,
	FieldQuantity,
	FieldQuantitySold,
	FieldSaleStartDate,
	FieldSaleEndDate,
	FieldEventID,
	FieldTicketType,
	FieldIsActive,
	FieldIsRefundable,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tickets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_ticket",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// PriceValidator is a validator for the "price" field. It is called by the builders before save.
	PriceValidator func(float64) error
	// QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	QuantityValidator func(int) error
	// DefaultQuantitySold holds the default value on creation for the "quantity_sold" field.
	DefaultQuantitySold int
	// QuantitySoldValidator is a validator for the "quantity_sold" field. It is called by the builders before save.
	QuantitySoldValidator func(int) error
	// TicketTypeValidator is a validator for the "ticket_type" field. It is called by the builders before save.
	TicketTypeValidator func(string) error
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultIsRefundable holds the default value on creation for the "is_refundable" field.
	DefaultIsRefundable bool
)

// OrderOption defines the ordering options for the Ticket queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByQuantitySold orders the results by the quantity_sold field.
func ByQuantitySold(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantitySold, opts...).ToFunc()
}

// BySaleStartDate orders the results by the sale_start_date field.
func BySaleStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSaleStartDate, opts...).ToFunc()
}

// BySaleEndDate orders the results by the sale_end_date field.
func BySaleEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSaleEndDate, opts...).ToFunc()
}

// ByEventID orders the results by the event_id field.
func ByEventID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventID, opts...).ToFunc()
}

// ByTicketType orders the results by the ticket_type field.
func ByTicketType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTicketType, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByIsRefundable orders the results by the is_refundable field.
func ByIsRefundable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRefundable, opts...).ToFunc()
}
