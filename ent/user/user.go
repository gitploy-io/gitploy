// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldRefresh holds the string denoting the refresh field in the database.
	FieldRefresh = "refresh"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeChatUser holds the string denoting the chat_user edge name in mutations.
	EdgeChatUser = "chat_user"
	// EdgePerms holds the string denoting the perms edge name in mutations.
	EdgePerms = "perms"
	// EdgeDeployments holds the string denoting the deployments edge name in mutations.
	EdgeDeployments = "deployments"
	// EdgeApprovals holds the string denoting the approvals edge name in mutations.
	EdgeApprovals = "approvals"
	// EdgeReviews holds the string denoting the reviews edge name in mutations.
	EdgeReviews = "reviews"
	// EdgeLocks holds the string denoting the locks edge name in mutations.
	EdgeLocks = "locks"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ChatUserTable is the table that holds the chat_user relation/edge.
	ChatUserTable = "chat_users"
	// ChatUserInverseTable is the table name for the ChatUser entity.
	// It exists in this package in order to avoid circular dependency with the "chatuser" package.
	ChatUserInverseTable = "chat_users"
	// ChatUserColumn is the table column denoting the chat_user relation/edge.
	ChatUserColumn = "user_id"
	// PermsTable is the table that holds the perms relation/edge.
	PermsTable = "perms"
	// PermsInverseTable is the table name for the Perm entity.
	// It exists in this package in order to avoid circular dependency with the "perm" package.
	PermsInverseTable = "perms"
	// PermsColumn is the table column denoting the perms relation/edge.
	PermsColumn = "user_id"
	// DeploymentsTable is the table that holds the deployments relation/edge.
	DeploymentsTable = "deployments"
	// DeploymentsInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentsInverseTable = "deployments"
	// DeploymentsColumn is the table column denoting the deployments relation/edge.
	DeploymentsColumn = "user_id"
	// ApprovalsTable is the table that holds the approvals relation/edge.
	ApprovalsTable = "approvals"
	// ApprovalsInverseTable is the table name for the Approval entity.
	// It exists in this package in order to avoid circular dependency with the "approval" package.
	ApprovalsInverseTable = "approvals"
	// ApprovalsColumn is the table column denoting the approvals relation/edge.
	ApprovalsColumn = "user_id"
	// ReviewsTable is the table that holds the reviews relation/edge.
	ReviewsTable = "reviews"
	// ReviewsInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewsInverseTable = "reviews"
	// ReviewsColumn is the table column denoting the reviews relation/edge.
	ReviewsColumn = "user_id"
	// LocksTable is the table that holds the locks relation/edge.
	LocksTable = "locks"
	// LocksInverseTable is the table name for the Lock entity.
	// It exists in this package in order to avoid circular dependency with the "lock" package.
	LocksInverseTable = "locks"
	// LocksColumn is the table column denoting the locks relation/edge.
	LocksColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldAvatar,
	FieldAdmin,
	FieldToken,
	FieldRefresh,
	FieldExpiry,
	FieldHash,
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
	// DefaultAdmin holds the default value on creation for the "admin" field.
	DefaultAdmin bool
	// DefaultHash holds the default value on creation for the "hash" field.
	DefaultHash func() string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
