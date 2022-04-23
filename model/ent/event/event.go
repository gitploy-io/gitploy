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
	// FieldDeploymentStatusID holds the string denoting the deployment_status_id field in the database.
	FieldDeploymentStatusID = "deployment_status_id"
	// FieldReviewID holds the string denoting the review_id field in the database.
	FieldReviewID = "review_id"
	// FieldDeletedID holds the string denoting the deleted_id field in the database.
	FieldDeletedID = "deleted_id"
	// EdgeDeploymentStatus holds the string denoting the deployment_status edge name in mutations.
	EdgeDeploymentStatus = "deployment_status"
	// EdgeReview holds the string denoting the review edge name in mutations.
	EdgeReview = "review"
	// EdgeNotificationRecord holds the string denoting the notification_record edge name in mutations.
	EdgeNotificationRecord = "notification_record"
	// Table holds the table name of the event in the database.
	Table = "events"
	// DeploymentStatusTable is the table that holds the deployment_status relation/edge.
	DeploymentStatusTable = "events"
	// DeploymentStatusInverseTable is the table name for the DeploymentStatus entity.
	// It exists in this package in order to avoid circular dependency with the "deploymentstatus" package.
	DeploymentStatusInverseTable = "deployment_status"
	// DeploymentStatusColumn is the table column denoting the deployment_status relation/edge.
	DeploymentStatusColumn = "deployment_status_id"
	// ReviewTable is the table that holds the review relation/edge.
	ReviewTable = "events"
	// ReviewInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewInverseTable = "reviews"
	// ReviewColumn is the table column denoting the review relation/edge.
	ReviewColumn = "review_id"
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
	FieldDeploymentStatusID,
	FieldReviewID,
	FieldDeletedID,
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
	KindDeploymentStatus Kind = "deployment_status"
	KindReview           Kind = "review"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindDeploymentStatus, KindReview:
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
	TypeDeleted Type = "deleted"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeCreated, TypeUpdated, TypeDeleted:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for type field: %q", _type)
	}
}
