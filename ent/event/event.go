// Code generated by entc, DO NOT EDIT.

package event

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeploymentID holds the string denoting the deployment_id field in the database.
	FieldDeploymentID = "deployment_id"
	// FieldApprovalID holds the string denoting the approval_id field in the database.
	FieldApprovalID = "approval_id"
	// EdgeDeployment holds the string denoting the deployment edge name in mutations.
	EdgeDeployment = "deployment"
	// EdgeApproval holds the string denoting the approval edge name in mutations.
	EdgeApproval = "approval"
	// EdgeNotificationRecord holds the string denoting the notification_record edge name in mutations.
	EdgeNotificationRecord = "notification_record"
	// Table holds the table name of the event in the database.
	Table = "events"
	// DeploymentTable is the table that holds the deployment relation/edge.
	DeploymentTable = "events"
	// DeploymentInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentInverseTable = "deployments"
	// DeploymentColumn is the table column denoting the deployment relation/edge.
	DeploymentColumn = "deployment_id"
	// ApprovalTable is the table that holds the approval relation/edge.
	ApprovalTable = "events"
	// ApprovalInverseTable is the table name for the Approval entity.
	// It exists in this package in order to avoid circular dependency with the "approval" package.
	ApprovalInverseTable = "approvals"
	// ApprovalColumn is the table column denoting the approval relation/edge.
	ApprovalColumn = "approval_id"
	// NotificationRecordTable is the table that holds the notification_record relation/edge.
	NotificationRecordTable = "notification_records"
	// NotificationRecordInverseTable is the table name for the NotificationRecord entity.
	// It exists in this package in order to avoid circular dependency with the "notificationrecord" package.
	NotificationRecordInverseTable = "notification_records"
	// NotificationRecordColumn is the table column denoting the notification_record relation/edge.
	NotificationRecordColumn = "event_id"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldKind,
	FieldType,
	FieldCreatedAt,
	FieldDeploymentID,
	FieldApprovalID,
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
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindDeployment Kind = "deployment"
	KindApproval   Kind = "approval"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindDeployment, KindApproval:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for kind field: %q", k)
	}
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeCreated Type = "created"
	TypeUpdated Type = "updated"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeCreated, TypeUpdated:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for type field: %q", _type)
	}
}
