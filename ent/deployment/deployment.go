// Code generated by entc, DO NOT EDIT.

package deployment

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the deployment type in the database.
	Label = "deployment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
	// FieldSha holds the string denoting the sha field in the database.
	FieldSha = "sha"
	// FieldEnv holds the string denoting the env field in the database.
	FieldEnv = "env"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRepoID holds the string denoting the repo_id field in the database.
	FieldRepoID = "repo_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRepo holds the string denoting the repo edge name in mutations.
	EdgeRepo = "repo"
	// EdgeApprovals holds the string denoting the approvals edge name in mutations.
	EdgeApprovals = "approvals"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// Table holds the table name of the deployment in the database.
	Table = "deployments"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "deployments"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// RepoTable is the table the holds the repo relation/edge.
	RepoTable = "deployments"
	// RepoInverseTable is the table name for the Repo entity.
	// It exists in this package in order to avoid circular dependency with the "repo" package.
	RepoInverseTable = "repos"
	// RepoColumn is the table column denoting the repo relation/edge.
	RepoColumn = "repo_id"
	// ApprovalsTable is the table the holds the approvals relation/edge.
	ApprovalsTable = "approvals"
	// ApprovalsInverseTable is the table name for the Approval entity.
	// It exists in this package in order to avoid circular dependency with the "approval" package.
	ApprovalsInverseTable = "approvals"
	// ApprovalsColumn is the table column denoting the approvals relation/edge.
	ApprovalsColumn = "deployment_id"
	// NotificationsTable is the table the holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "deployment_id"
)

// Columns holds all SQL columns for deployment fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldUID,
	FieldType,
	FieldRef,
	FieldSha,
	FieldEnv,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
	FieldRepoID,
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

// TypeCommit is the default value of the Type enum.
const DefaultType = TypeCommit

// Type values.
const (
	TypeCommit Type = "commit"
	TypeBranch Type = "branch"
	TypeTag    Type = "tag"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeCommit, TypeBranch, TypeTag:
		return nil
	default:
		return fmt.Errorf("deployment: invalid enum value for type field: %q", _type)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusWaiting is the default value of the Status enum.
const DefaultStatus = StatusWaiting

// Status values.
const (
	StatusWaiting Status = "waiting"
	StatusCreated Status = "created"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusWaiting, StatusCreated, StatusRunning, StatusSuccess, StatusFailure:
		return nil
	default:
		return fmt.Errorf("deployment: invalid enum value for status field: %q", s)
	}
}
