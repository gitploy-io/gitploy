// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatUsersColumns holds the columns for the "chat_users" table.
	ChatUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "refresh", Type: field.TypeString},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "bot_token", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64, Unique: true},
	}
	// ChatUsersTable holds the schema information for the "chat_users" table.
	ChatUsersTable = &schema.Table{
		Name:       "chat_users",
		Columns:    ChatUsersColumns,
		PrimaryKey: []*schema.Column{ChatUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chat_users_users_chat_user",
				Columns:    []*schema.Column{ChatUsersColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DeploymentsColumns holds the columns for the "deployments" table.
	DeploymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"commit", "branch", "tag"}, Default: "commit"},
		{Name: "env", Type: field.TypeString},
		{Name: "ref", Type: field.TypeString},
		{Name: "dynamic_payload", Type: field.TypeJSON, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"waiting", "created", "queued", "running", "success", "failure", "canceled"}, Default: "waiting"},
		{Name: "uid", Type: field.TypeInt64, Nullable: true},
		{Name: "sha", Type: field.TypeString, Nullable: true},
		{Name: "html_url", Type: field.TypeString, Nullable: true, Size: 2000},
		{Name: "production_environment", Type: field.TypeBool, Default: false},
		{Name: "is_rollback", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_approval_enabled", Type: field.TypeBool, Nullable: true},
		{Name: "required_approval_count", Type: field.TypeInt, Nullable: true},
		{Name: "repo_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// DeploymentsTable holds the schema information for the "deployments" table.
	DeploymentsTable = &schema.Table{
		Name:       "deployments",
		Columns:    DeploymentsColumns,
		PrimaryKey: []*schema.Column{DeploymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployments_repos_deployments",
				Columns:    []*schema.Column{DeploymentsColumns[16]},
				RefColumns: []*schema.Column{ReposColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "deployments_users_deployments",
				Columns:    []*schema.Column{DeploymentsColumns[17]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "deployment_repo_id_env_status_updated_at",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[16], DeploymentsColumns[3], DeploymentsColumns[6], DeploymentsColumns[13]},
			},
			{
				Name:    "deployment_repo_id_env_created_at",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[16], DeploymentsColumns[3], DeploymentsColumns[12]},
			},
			{
				Name:    "deployment_repo_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[16], DeploymentsColumns[12]},
			},
			{
				Name:    "deployment_repo_id_number",
				Unique:  true,
				Columns: []*schema.Column{DeploymentsColumns[16], DeploymentsColumns[1]},
			},
			{
				Name:    "deployment_uid",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[7]},
			},
			{
				Name:    "deployment_created_at_status",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[12], DeploymentsColumns[6]},
			},
		},
	}
	// DeploymentStatisticsColumns holds the columns for the "deployment_statistics" table.
	DeploymentStatisticsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "env", Type: field.TypeString},
		{Name: "count", Type: field.TypeInt, Default: 0},
		{Name: "rollback_count", Type: field.TypeInt, Default: 0},
		{Name: "additions", Type: field.TypeInt, Default: 0},
		{Name: "deletions", Type: field.TypeInt, Default: 0},
		{Name: "changes", Type: field.TypeInt, Default: 0},
		{Name: "lead_time_seconds", Type: field.TypeInt, Default: 0},
		{Name: "commit_count", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "repo_id", Type: field.TypeInt64},
	}
	// DeploymentStatisticsTable holds the schema information for the "deployment_statistics" table.
	DeploymentStatisticsTable = &schema.Table{
		Name:       "deployment_statistics",
		Columns:    DeploymentStatisticsColumns,
		PrimaryKey: []*schema.Column{DeploymentStatisticsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployment_statistics_repos_deployment_statistics",
				Columns:    []*schema.Column{DeploymentStatisticsColumns[11]},
				RefColumns: []*schema.Column{ReposColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "deploymentstatistics_repo_id_env",
				Unique:  true,
				Columns: []*schema.Column{DeploymentStatisticsColumns[11], DeploymentStatisticsColumns[1]},
			},
			{
				Name:    "deploymentstatistics_updated_at",
				Unique:  false,
				Columns: []*schema.Column{DeploymentStatisticsColumns[10]},
			},
		},
	}
	// DeploymentStatusColumns holds the columns for the "deployment_status" table.
	DeploymentStatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "log_url", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deployment_id", Type: field.TypeInt},
		{Name: "repo_id", Type: field.TypeInt64, Nullable: true},
	}
	// DeploymentStatusTable holds the schema information for the "deployment_status" table.
	DeploymentStatusTable = &schema.Table{
		Name:       "deployment_status",
		Columns:    DeploymentStatusColumns,
		PrimaryKey: []*schema.Column{DeploymentStatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployment_status_deployments_deployment_statuses",
				Columns:    []*schema.Column{DeploymentStatusColumns[6]},
				RefColumns: []*schema.Column{DeploymentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "deployment_status_repos_deployment_statuses",
				Columns:    []*schema.Column{DeploymentStatusColumns[7]},
				RefColumns: []*schema.Column{ReposColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"deployment", "deployment_status", "review"}},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"created", "updated", "deleted"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_id", Type: field.TypeInt, Nullable: true},
		{Name: "deployment_id", Type: field.TypeInt, Nullable: true},
		{Name: "deployment_status_id", Type: field.TypeInt, Nullable: true},
		{Name: "review_id", Type: field.TypeInt, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_deployments_event",
				Columns:    []*schema.Column{EventsColumns[5]},
				RefColumns: []*schema.Column{DeploymentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "events_deployment_status_event",
				Columns:    []*schema.Column{EventsColumns[6]},
				RefColumns: []*schema.Column{DeploymentStatusColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "events_reviews_event",
				Columns:    []*schema.Column{EventsColumns[7]},
				RefColumns: []*schema.Column{ReviewsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "event_created_at",
				Unique:  false,
				Columns: []*schema.Column{EventsColumns[3]},
			},
		},
	}
	// LocksColumns holds the columns for the "locks" table.
	LocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "env", Type: field.TypeString},
		{Name: "expired_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "repo_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// LocksTable holds the schema information for the "locks" table.
	LocksTable = &schema.Table{
		Name:       "locks",
		Columns:    LocksColumns,
		PrimaryKey: []*schema.Column{LocksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "locks_repos_locks",
				Columns:    []*schema.Column{LocksColumns[4]},
				RefColumns: []*schema.Column{ReposColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "locks_users_locks",
				Columns:    []*schema.Column{LocksColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "lock_repo_id_env",
				Unique:  true,
				Columns: []*schema.Column{LocksColumns[4], LocksColumns[1]},
			},
		},
	}
	// NotificationRecordsColumns holds the columns for the "notification_records" table.
	NotificationRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event_id", Type: field.TypeInt, Unique: true},
	}
	// NotificationRecordsTable holds the schema information for the "notification_records" table.
	NotificationRecordsTable = &schema.Table{
		Name:       "notification_records",
		Columns:    NotificationRecordsColumns,
		PrimaryKey: []*schema.Column{NotificationRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notification_records_events_notification_record",
				Columns:    []*schema.Column{NotificationRecordsColumns[1]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PermsColumns holds the columns for the "perms" table.
	PermsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "repo_perm", Type: field.TypeEnum, Enums: []string{"read", "write", "admin"}, Default: "read"},
		{Name: "synced_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "timestamp(6)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "repo_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// PermsTable holds the schema information for the "perms" table.
	PermsTable = &schema.Table{
		Name:       "perms",
		Columns:    PermsColumns,
		PrimaryKey: []*schema.Column{PermsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "perms_repos_perms",
				Columns:    []*schema.Column{PermsColumns[5]},
				RefColumns: []*schema.Column{ReposColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "perms_users_perms",
				Columns:    []*schema.Column{PermsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "perm_repo_id_user_id",
				Unique:  false,
				Columns: []*schema.Column{PermsColumns[5], PermsColumns[6]},
			},
			{
				Name:    "perm_user_id_synced_at",
				Unique:  false,
				Columns: []*schema.Column{PermsColumns[6], PermsColumns[2]},
			},
		},
	}
	// ReposColumns holds the columns for the "repos" table.
	ReposColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "namespace", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "config_path", Type: field.TypeString, Default: "deploy.yml"},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "webhook_id", Type: field.TypeInt64, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "latest_deployed_at", Type: field.TypeTime, Nullable: true},
		{Name: "owner_id", Type: field.TypeInt64, Nullable: true},
	}
	// ReposTable holds the schema information for the "repos" table.
	ReposTable = &schema.Table{
		Name:       "repos",
		Columns:    ReposColumns,
		PrimaryKey: []*schema.Column{ReposColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "repos_users_repos",
				Columns:    []*schema.Column{ReposColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "repo_namespace_name",
				Unique:  true,
				Columns: []*schema.Column{ReposColumns[1], ReposColumns[2]},
			},
			{
				Name:    "repo_name",
				Unique:  false,
				Columns: []*schema.Column{ReposColumns[2]},
			},
			{
				Name:    "repo_active",
				Unique:  false,
				Columns: []*schema.Column{ReposColumns[5]},
			},
		},
	}
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "rejected", "approved"}, Default: "pending"},
		{Name: "comment", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deployment_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reviews_deployments_reviews",
				Columns:    []*schema.Column{ReviewsColumns[5]},
				RefColumns: []*schema.Column{DeploymentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "reviews_users_reviews",
				Columns:    []*schema.Column{ReviewsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "login", Type: field.TypeString, Unique: true},
		{Name: "avatar", Type: field.TypeString},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "token", Type: field.TypeString},
		{Name: "refresh", Type: field.TypeString},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "hash", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatUsersTable,
		DeploymentsTable,
		DeploymentStatisticsTable,
		DeploymentStatusTable,
		EventsTable,
		LocksTable,
		NotificationRecordsTable,
		PermsTable,
		ReposTable,
		ReviewsTable,
		UsersTable,
	}
)

func init() {
	ChatUsersTable.ForeignKeys[0].RefTable = UsersTable
	DeploymentsTable.ForeignKeys[0].RefTable = ReposTable
	DeploymentsTable.ForeignKeys[1].RefTable = UsersTable
	DeploymentStatisticsTable.ForeignKeys[0].RefTable = ReposTable
	DeploymentStatusTable.ForeignKeys[0].RefTable = DeploymentsTable
	DeploymentStatusTable.ForeignKeys[1].RefTable = ReposTable
	EventsTable.ForeignKeys[0].RefTable = DeploymentsTable
	EventsTable.ForeignKeys[1].RefTable = DeploymentStatusTable
	EventsTable.ForeignKeys[2].RefTable = ReviewsTable
	LocksTable.ForeignKeys[0].RefTable = ReposTable
	LocksTable.ForeignKeys[1].RefTable = UsersTable
	NotificationRecordsTable.ForeignKeys[0].RefTable = EventsTable
	PermsTable.ForeignKeys[0].RefTable = ReposTable
	PermsTable.ForeignKeys[1].RefTable = UsersTable
	ReposTable.ForeignKeys[0].RefTable = UsersTable
	ReviewsTable.ForeignKeys[0].RefTable = DeploymentsTable
	ReviewsTable.ForeignKeys[1].RefTable = UsersTable
}
