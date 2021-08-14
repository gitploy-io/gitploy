// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hanjunlee/gitploy/ent/repo"
)

// Repo is the model entity for the Repo schema.
type Repo struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Description holds the value of the "description" field.
	Description string `json:"description"`
	// ConfigPath holds the value of the "config_path" field.
	ConfigPath string `json:"config_path"`
	// Active holds the value of the "active" field.
	Active bool `json:"active"`
	// WebhookID holds the value of the "webhook_id" field.
	WebhookID int64 `json:"webhook_id"`
	// SyncedAt holds the value of the "synced_at" field.
	SyncedAt time.Time `json:"synced_at"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// LatestDeployedAt holds the value of the "latest_deployed_at" field.
	LatestDeployedAt time.Time `json:"latest_deployed_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RepoQuery when eager-loading is set.
	Edges RepoEdges `json:"edges"`
}

// RepoEdges holds the relations/edges for other nodes in the graph.
type RepoEdges struct {
	// Perms holds the value of the perms edge.
	Perms []*Perm `json:"perms,omitempty"`
	// Deployments holds the value of the deployments edge.
	Deployments []*Deployment `json:"deployments,omitempty"`
	// Callback holds the value of the callback edge.
	Callback []*Callback `json:"callback,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PermsOrErr returns the Perms value or an error if the edge
// was not loaded in eager-loading.
func (e RepoEdges) PermsOrErr() ([]*Perm, error) {
	if e.loadedTypes[0] {
		return e.Perms, nil
	}
	return nil, &NotLoadedError{edge: "perms"}
}

// DeploymentsOrErr returns the Deployments value or an error if the edge
// was not loaded in eager-loading.
func (e RepoEdges) DeploymentsOrErr() ([]*Deployment, error) {
	if e.loadedTypes[1] {
		return e.Deployments, nil
	}
	return nil, &NotLoadedError{edge: "deployments"}
}

// CallbackOrErr returns the Callback value or an error if the edge
// was not loaded in eager-loading.
func (e RepoEdges) CallbackOrErr() ([]*Callback, error) {
	if e.loadedTypes[2] {
		return e.Callback, nil
	}
	return nil, &NotLoadedError{edge: "callback"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Repo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case repo.FieldActive:
			values[i] = new(sql.NullBool)
		case repo.FieldWebhookID:
			values[i] = new(sql.NullInt64)
		case repo.FieldID, repo.FieldNamespace, repo.FieldName, repo.FieldDescription, repo.FieldConfigPath:
			values[i] = new(sql.NullString)
		case repo.FieldSyncedAt, repo.FieldCreatedAt, repo.FieldUpdatedAt, repo.FieldLatestDeployedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Repo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Repo fields.
func (r *Repo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case repo.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case repo.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				r.Namespace = value.String
			}
		case repo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case repo.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case repo.FieldConfigPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config_path", values[i])
			} else if value.Valid {
				r.ConfigPath = value.String
			}
		case repo.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				r.Active = value.Bool
			}
		case repo.FieldWebhookID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field webhook_id", values[i])
			} else if value.Valid {
				r.WebhookID = value.Int64
			}
		case repo.FieldSyncedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field synced_at", values[i])
			} else if value.Valid {
				r.SyncedAt = value.Time
			}
		case repo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case repo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case repo.FieldLatestDeployedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field latest_deployed_at", values[i])
			} else if value.Valid {
				r.LatestDeployedAt = value.Time
			}
		}
	}
	return nil
}

// QueryPerms queries the "perms" edge of the Repo entity.
func (r *Repo) QueryPerms() *PermQuery {
	return (&RepoClient{config: r.config}).QueryPerms(r)
}

// QueryDeployments queries the "deployments" edge of the Repo entity.
func (r *Repo) QueryDeployments() *DeploymentQuery {
	return (&RepoClient{config: r.config}).QueryDeployments(r)
}

// QueryCallback queries the "callback" edge of the Repo entity.
func (r *Repo) QueryCallback() *CallbackQuery {
	return (&RepoClient{config: r.config}).QueryCallback(r)
}

// Update returns a builder for updating this Repo.
// Note that you need to call Repo.Unwrap() before calling this method if this Repo
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Repo) Update() *RepoUpdateOne {
	return (&RepoClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Repo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Repo) Unwrap() *Repo {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Repo is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Repo) String() string {
	var builder strings.Builder
	builder.WriteString("Repo(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", namespace=")
	builder.WriteString(r.Namespace)
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", description=")
	builder.WriteString(r.Description)
	builder.WriteString(", config_path=")
	builder.WriteString(r.ConfigPath)
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", r.Active))
	builder.WriteString(", webhook_id=")
	builder.WriteString(fmt.Sprintf("%v", r.WebhookID))
	builder.WriteString(", synced_at=")
	builder.WriteString(r.SyncedAt.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", latest_deployed_at=")
	builder.WriteString(r.LatestDeployedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Repos is a parsable slice of Repo.
type Repos []*Repo

func (r Repos) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
