// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApprovalsColumns holds the columns for the "approvals" table.
	ApprovalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "declined", "approved"}, Default: "pending"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deployment_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// ApprovalsTable holds the schema information for the "approvals" table.
	ApprovalsTable = &schema.Table{
		Name:       "approvals",
		Columns:    ApprovalsColumns,
		PrimaryKey: []*schema.Column{ApprovalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "approvals_deployments_approvals",
				Columns:    []*schema.Column{ApprovalsColumns[4]},
				RefColumns: []*schema.Column{DeploymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "approvals_users_approvals",
				Columns:    []*schema.Column{ApprovalsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChatCallbacksColumns holds the columns for the "chat_callbacks" table.
	ChatCallbacksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "state", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"deploy", "rollback"}},
		{Name: "is_opened", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "chat_user_id", Type: field.TypeString, Nullable: true},
		{Name: "repo_id", Type: field.TypeString, Nullable: true},
	}
	// ChatCallbacksTable holds the schema information for the "chat_callbacks" table.
	ChatCallbacksTable = &schema.Table{
		Name:       "chat_callbacks",
		Columns:    ChatCallbacksColumns,
		PrimaryKey: []*schema.Column{ChatCallbacksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chat_callbacks_chat_users_chat_callback",
				Columns:    []*schema.Column{ChatCallbacksColumns[6]},
				RefColumns: []*schema.Column{ChatUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "chat_callbacks_repos_chat_callback",
				Columns:    []*schema.Column{ChatCallbacksColumns[7]},
				RefColumns: []*schema.Column{ReposColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChatUsersColumns holds the columns for the "chat_users" table.
	ChatUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "refresh", Type: field.TypeString},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "bot_token", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeString, Unique: true, Nullable: true},
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
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "chatuser_user_id",
				Unique:  false,
				Columns: []*schema.Column{ChatUsersColumns[7]},
			},
		},
	}
	// DeploymentsColumns holds the columns for the "deployments" table.
	DeploymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"commit", "branch", "tag"}, Default: "commit"},
		{Name: "ref", Type: field.TypeString},
		{Name: "sha", Type: field.TypeString},
		{Name: "env", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"waiting", "created", "running", "success", "failure"}, Default: "waiting"},
		{Name: "uid", Type: field.TypeInt64, Nullable: true},
		{Name: "is_rollback", Type: field.TypeBool, Default: false},
		{Name: "is_approval_enabled", Type: field.TypeBool, Default: false},
		{Name: "required_approval_count", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "repo_id", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// DeploymentsTable holds the schema information for the "deployments" table.
	DeploymentsTable = &schema.Table{
		Name:       "deployments",
		Columns:    DeploymentsColumns,
		PrimaryKey: []*schema.Column{DeploymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployments_repos_deployments",
				Columns:    []*schema.Column{DeploymentsColumns[13]},
				RefColumns: []*schema.Column{ReposColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "deployments_users_deployments",
				Columns:    []*schema.Column{DeploymentsColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "deployment_user_id",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[14]},
			},
			{
				Name:    "deployment_repo_id",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[13]},
			},
			{
				Name:    "deployment_repo_id_env_created_at",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[13], DeploymentsColumns[5], DeploymentsColumns[11]},
			},
			{
				Name:    "deployment_repo_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[13], DeploymentsColumns[11]},
			},
			{
				Name:    "deployment_number_repo_id",
				Unique:  true,
				Columns: []*schema.Column{DeploymentsColumns[1], DeploymentsColumns[13]},
			},
			{
				Name:    "deployment_uid",
				Unique:  false,
				Columns: []*schema.Column{DeploymentsColumns[7]},
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
		{Name: "deployment_id", Type: field.TypeInt, Nullable: true},
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
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "deploymentstatus_deployment_id",
				Unique:  false,
				Columns: []*schema.Column{DeploymentStatusColumns[6]},
			},
		},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"deployment", "approval_requested", "approval_responded"}, Default: "deployment"},
		{Name: "repo_namespace", Type: field.TypeString},
		{Name: "repo_name", Type: field.TypeString},
		{Name: "deployment_number", Type: field.TypeInt},
		{Name: "deployment_type", Type: field.TypeString},
		{Name: "deployment_ref", Type: field.TypeString},
		{Name: "deployment_env", Type: field.TypeString},
		{Name: "deployment_status", Type: field.TypeString},
		{Name: "deployment_login", Type: field.TypeString},
		{Name: "approval_status", Type: field.TypeString, Nullable: true},
		{Name: "approval_login", Type: field.TypeString, Nullable: true},
		{Name: "notified", Type: field.TypeBool, Default: false},
		{Name: "checked", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notifications_users_notification",
				Columns:    []*schema.Column{NotificationsColumns[16]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "notification_user_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationsColumns[16]},
			},
			{
				Name:    "notification_user_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{NotificationsColumns[16], NotificationsColumns[14]},
			},
		},
	}
	// PermsColumns holds the columns for the "perms" table.
	PermsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "repo_perm", Type: field.TypeEnum, Enums: []string{"read", "write", "admin"}, Default: "read"},
		{Name: "synced_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "repo_id", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
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
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "perms_users_perms",
				Columns:    []*schema.Column{PermsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "perm_user_id",
				Unique:  false,
				Columns: []*schema.Column{PermsColumns[6]},
			},
			{
				Name:    "perm_repo_id",
				Unique:  false,
				Columns: []*schema.Column{PermsColumns[5]},
			},
			{
				Name:    "perm_repo_id_user_id",
				Unique:  false,
				Columns: []*schema.Column{PermsColumns[5], PermsColumns[6]},
			},
		},
	}
	// ReposColumns holds the columns for the "repos" table.
	ReposColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "namespace", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "config_path", Type: field.TypeString, Default: "deploy.yml"},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "webhook_id", Type: field.TypeInt64, Nullable: true},
		{Name: "synced_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "latest_deployed_at", Type: field.TypeTime, Nullable: true},
	}
	// ReposTable holds the schema information for the "repos" table.
	ReposTable = &schema.Table{
		Name:        "repos",
		Columns:     ReposColumns,
		PrimaryKey:  []*schema.Column{ReposColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "login", Type: field.TypeString, Unique: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "token", Type: field.TypeString},
		{Name: "refresh", Type: field.TypeString},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "hash", Type: field.TypeString, Unique: true},
		{Name: "synced_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApprovalsTable,
		ChatCallbacksTable,
		ChatUsersTable,
		DeploymentsTable,
		DeploymentStatusTable,
		NotificationsTable,
		PermsTable,
		ReposTable,
		UsersTable,
	}
)

func init() {
	ApprovalsTable.ForeignKeys[0].RefTable = DeploymentsTable
	ApprovalsTable.ForeignKeys[1].RefTable = UsersTable
	ChatCallbacksTable.ForeignKeys[0].RefTable = ChatUsersTable
	ChatCallbacksTable.ForeignKeys[1].RefTable = ReposTable
	ChatUsersTable.ForeignKeys[0].RefTable = UsersTable
	DeploymentsTable.ForeignKeys[0].RefTable = ReposTable
	DeploymentsTable.ForeignKeys[1].RefTable = UsersTable
	DeploymentStatusTable.ForeignKeys[0].RefTable = DeploymentsTable
	NotificationsTable.ForeignKeys[0].RefTable = UsersTable
	PermsTable.ForeignKeys[0].RefTable = ReposTable
	PermsTable.ForeignKeys[1].RefTable = UsersTable
}
