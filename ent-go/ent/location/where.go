// Code generated by ent, DO NOT EDIT.

package location

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kayceeDev/caspa-events/ent-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldUpdateTime, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldUUID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldName, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldAddress, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCity, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldState, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCountry, v))
}

// PostalCode applies equality check predicate on the "postal_code" field. It's identical to PostalCodeEQ.
func PostalCode(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldPostalCode, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldLongitude, v))
}

// Capacity applies equality check predicate on the "capacity" field. It's identical to CapacityEQ.
func Capacity(v int) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCapacity, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldUpdateTime, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldUUID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldName, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldAddress, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldCity, v))
}

// CityIsNil applies the IsNil predicate on the "city" field.
func CityIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldCity))
}

// CityNotNil applies the NotNil predicate on the "city" field.
func CityNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldCity))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldCity, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldState, v))
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldState))
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldState))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldState, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryIsNil applies the IsNil predicate on the "country" field.
func CountryIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldCountry))
}

// CountryNotNil applies the NotNil predicate on the "country" field.
func CountryNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldCountry))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldCountry, v))
}

// PostalCodeEQ applies the EQ predicate on the "postal_code" field.
func PostalCodeEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldPostalCode, v))
}

// PostalCodeNEQ applies the NEQ predicate on the "postal_code" field.
func PostalCodeNEQ(v string) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldPostalCode, v))
}

// PostalCodeIn applies the In predicate on the "postal_code" field.
func PostalCodeIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldPostalCode, vs...))
}

// PostalCodeNotIn applies the NotIn predicate on the "postal_code" field.
func PostalCodeNotIn(vs ...string) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldPostalCode, vs...))
}

// PostalCodeGT applies the GT predicate on the "postal_code" field.
func PostalCodeGT(v string) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldPostalCode, v))
}

// PostalCodeGTE applies the GTE predicate on the "postal_code" field.
func PostalCodeGTE(v string) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldPostalCode, v))
}

// PostalCodeLT applies the LT predicate on the "postal_code" field.
func PostalCodeLT(v string) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldPostalCode, v))
}

// PostalCodeLTE applies the LTE predicate on the "postal_code" field.
func PostalCodeLTE(v string) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldPostalCode, v))
}

// PostalCodeContains applies the Contains predicate on the "postal_code" field.
func PostalCodeContains(v string) predicate.Location {
	return predicate.Location(sql.FieldContains(FieldPostalCode, v))
}

// PostalCodeHasPrefix applies the HasPrefix predicate on the "postal_code" field.
func PostalCodeHasPrefix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasPrefix(FieldPostalCode, v))
}

// PostalCodeHasSuffix applies the HasSuffix predicate on the "postal_code" field.
func PostalCodeHasSuffix(v string) predicate.Location {
	return predicate.Location(sql.FieldHasSuffix(FieldPostalCode, v))
}

// PostalCodeIsNil applies the IsNil predicate on the "postal_code" field.
func PostalCodeIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldPostalCode))
}

// PostalCodeNotNil applies the NotNil predicate on the "postal_code" field.
func PostalCodeNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldPostalCode))
}

// PostalCodeEqualFold applies the EqualFold predicate on the "postal_code" field.
func PostalCodeEqualFold(v string) predicate.Location {
	return predicate.Location(sql.FieldEqualFold(FieldPostalCode, v))
}

// PostalCodeContainsFold applies the ContainsFold predicate on the "postal_code" field.
func PostalCodeContainsFold(v string) predicate.Location {
	return predicate.Location(sql.FieldContainsFold(FieldPostalCode, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldLatitude, v))
}

// LatitudeIsNil applies the IsNil predicate on the "latitude" field.
func LatitudeIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldLatitude))
}

// LatitudeNotNil applies the NotNil predicate on the "latitude" field.
func LatitudeNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldLatitude))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldLongitude, v))
}

// LongitudeIsNil applies the IsNil predicate on the "longitude" field.
func LongitudeIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldLongitude))
}

// LongitudeNotNil applies the NotNil predicate on the "longitude" field.
func LongitudeNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldLongitude))
}

// CapacityEQ applies the EQ predicate on the "capacity" field.
func CapacityEQ(v int) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldCapacity, v))
}

// CapacityNEQ applies the NEQ predicate on the "capacity" field.
func CapacityNEQ(v int) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldCapacity, v))
}

// CapacityIn applies the In predicate on the "capacity" field.
func CapacityIn(vs ...int) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldCapacity, vs...))
}

// CapacityNotIn applies the NotIn predicate on the "capacity" field.
func CapacityNotIn(vs ...int) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldCapacity, vs...))
}

// CapacityGT applies the GT predicate on the "capacity" field.
func CapacityGT(v int) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldCapacity, v))
}

// CapacityGTE applies the GTE predicate on the "capacity" field.
func CapacityGTE(v int) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldCapacity, v))
}

// CapacityLT applies the LT predicate on the "capacity" field.
func CapacityLT(v int) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldCapacity, v))
}

// CapacityLTE applies the LTE predicate on the "capacity" field.
func CapacityLTE(v int) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldCapacity, v))
}

// CapacityIsNil applies the IsNil predicate on the "capacity" field.
func CapacityIsNil() predicate.Location {
	return predicate.Location(sql.FieldIsNull(FieldCapacity))
}

// CapacityNotNil applies the NotNil predicate on the "capacity" field.
func CapacityNotNil() predicate.Location {
	return predicate.Location(sql.FieldNotNull(FieldCapacity))
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Location) predicate.Location {
	return predicate.Location(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Location) predicate.Location {
	return predicate.Location(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Location) predicate.Location {
	return predicate.Location(sql.NotPredicates(p))
}
