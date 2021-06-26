// Code generated by entc, DO NOT EDIT.

package notification

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the notification type in the database.
	Label = "notification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldResourceID holds the string denoting the resource_id field in the database.
	FieldResourceID = "resource_id"
	// FieldNotified holds the string denoting the notified field in the database.
	FieldNotified = "notified"
	// FieldChecked holds the string denoting the checked field in the database.
	FieldChecked = "checked"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the notification in the database.
	Table = "notifications"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "notifications"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for notification fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldResourceID,
	FieldNotified,
	FieldChecked,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
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
	// DefaultNotified holds the default value on creation for the "notified" field.
	DefaultNotified bool
	// DefaultChecked holds the default value on creation for the "checked" field.
	DefaultChecked bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// TypeDeployment is the default value of the Type enum.
const DefaultType = TypeDeployment

// Type values.
const (
	TypeDeployment Type = "deployment"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeDeployment:
		return nil
	default:
		return fmt.Errorf("notification: invalid enum value for type field: %q", _type)
	}
}