// Code generated by entc, DO NOT EDIT.

package lock

import (
	"time"
)

const (
	// Label holds the string label denoting the lock type in the database.
	Label = "lock"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEnv holds the string denoting the env field in the database.
	FieldEnv = "env"
	// FieldExpiredAt holds the string denoting the expired_at field in the database.
	FieldExpiredAt = "expired_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRepoID holds the string denoting the repo_id field in the database.
	FieldRepoID = "repo_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRepo holds the string denoting the repo edge name in mutations.
	EdgeRepo = "repo"
	// Table holds the table name of the lock in the database.
	Table = "locks"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "locks"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// RepoTable is the table that holds the repo relation/edge.
	RepoTable = "locks"
	// RepoInverseTable is the table name for the Repo entity.
	// It exists in this package in order to avoid circular dependency with the "repo" package.
	RepoInverseTable = "repos"
	// RepoColumn is the table column denoting the repo relation/edge.
	RepoColumn = "repo_id"
)

// Columns holds all SQL columns for lock fields.
var Columns = []string{
	FieldID,
	FieldEnv,
	FieldExpiredAt,
	FieldCreatedAt,
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
)
