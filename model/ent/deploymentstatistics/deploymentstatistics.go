// Code generated by entc, DO NOT EDIT.

package deploymentstatistics

import (
	"time"
)

const (
	// Label holds the string label denoting the deploymentstatistics type in the database.
	Label = "deployment_statistics"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEnv holds the string denoting the env field in the database.
	FieldEnv = "env"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldRollbackCount holds the string denoting the rollback_count field in the database.
	FieldRollbackCount = "rollback_count"
	// FieldAdditions holds the string denoting the additions field in the database.
	FieldAdditions = "additions"
	// FieldDeletions holds the string denoting the deletions field in the database.
	FieldDeletions = "deletions"
	// FieldChanges holds the string denoting the changes field in the database.
	FieldChanges = "changes"
	// FieldLeadTimeSeconds holds the string denoting the lead_time_seconds field in the database.
	FieldLeadTimeSeconds = "lead_time_seconds"
	// FieldCommitCount holds the string denoting the commit_count field in the database.
	FieldCommitCount = "commit_count"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldRepoID holds the string denoting the repo_id field in the database.
	FieldRepoID = "repo_id"
	// EdgeRepo holds the string denoting the repo edge name in mutations.
	EdgeRepo = "repo"
	// Table holds the table name of the deploymentstatistics in the database.
	Table = "deployment_statistics"
	// RepoTable is the table that holds the repo relation/edge.
	RepoTable = "deployment_statistics"
	// RepoInverseTable is the table name for the Repo entity.
	// It exists in this package in order to avoid circular dependency with the "repo" package.
	RepoInverseTable = "repos"
	// RepoColumn is the table column denoting the repo relation/edge.
	RepoColumn = "repo_id"
)

// Columns holds all SQL columns for deploymentstatistics fields.
var Columns = []string{
	FieldID,
	FieldEnv,
	FieldCount,
	FieldRollbackCount,
	FieldAdditions,
	FieldDeletions,
	FieldChanges,
	FieldLeadTimeSeconds,
	FieldCommitCount,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultCount holds the default value on creation for the "count" field.
	DefaultCount int
	// DefaultRollbackCount holds the default value on creation for the "rollback_count" field.
	DefaultRollbackCount int
	// DefaultAdditions holds the default value on creation for the "additions" field.
	DefaultAdditions int
	// DefaultDeletions holds the default value on creation for the "deletions" field.
	DefaultDeletions int
	// DefaultChanges holds the default value on creation for the "changes" field.
	DefaultChanges int
	// DefaultLeadTimeSeconds holds the default value on creation for the "lead_time_seconds" field.
	DefaultLeadTimeSeconds int
	// DefaultCommitCount holds the default value on creation for the "commit_count" field.
	DefaultCommitCount int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
