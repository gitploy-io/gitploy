// Code generated by entc, DO NOT EDIT.

package chatcallback

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the chatcallback type in the database.
	Label = "chat_callback"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsOpened holds the string denoting the is_opened field in the database.
	FieldIsOpened = "is_opened"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the chatcallback in the database.
	Table = "chat_callbacks"
)

// Columns holds all SQL columns for chatcallback fields.
var Columns = []string{
	FieldID,
	FieldState,
	FieldType,
	FieldIsOpened,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeDeploy   Type = "deploy"
	TypeRollback Type = "rollback"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeDeploy, TypeRollback:
		return nil
	default:
		return fmt.Errorf("chatcallback: invalid enum value for type field: %q", _type)
	}
}
