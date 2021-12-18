// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitploy-io/gitploy/model/ent/deploymentstatistics"
	"github.com/gitploy-io/gitploy/model/ent/repo"
)

// DeploymentStatistics is the model entity for the DeploymentStatistics schema.
type DeploymentStatistics struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Env holds the value of the "env" field.
	Env string `json:"env"`
	// Count holds the value of the "count" field.
	Count int `json:"count"`
	// RollbackCount holds the value of the "rollback_count" field.
	RollbackCount int `json:"rollback_count"`
	// Additions holds the value of the "additions" field.
	Additions int `json:"additions"`
	// Deletions holds the value of the "deletions" field.
	Deletions int `json:"deletions"`
	// Changes holds the value of the "changes" field.
	Changes int `json:"changes"`
	// LeadTimeSeconds holds the value of the "lead_time_seconds" field.
	LeadTimeSeconds int `json:"lead_time_seconds"`
	// CommitCount holds the value of the "commit_count" field.
	CommitCount int `json:"commit_count"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// RepoID holds the value of the "repo_id" field.
	RepoID int64 `json:"repo_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeploymentStatisticsQuery when eager-loading is set.
	Edges DeploymentStatisticsEdges `json:"edges"`
}

// DeploymentStatisticsEdges holds the relations/edges for other nodes in the graph.
type DeploymentStatisticsEdges struct {
	// Repo holds the value of the repo edge.
	Repo *Repo `json:"repo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RepoOrErr returns the Repo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentStatisticsEdges) RepoOrErr() (*Repo, error) {
	if e.loadedTypes[0] {
		if e.Repo == nil {
			// The edge repo was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: repo.Label}
		}
		return e.Repo, nil
	}
	return nil, &NotLoadedError{edge: "repo"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeploymentStatistics) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deploymentstatistics.FieldID, deploymentstatistics.FieldCount, deploymentstatistics.FieldRollbackCount, deploymentstatistics.FieldAdditions, deploymentstatistics.FieldDeletions, deploymentstatistics.FieldChanges, deploymentstatistics.FieldLeadTimeSeconds, deploymentstatistics.FieldCommitCount, deploymentstatistics.FieldRepoID:
			values[i] = new(sql.NullInt64)
		case deploymentstatistics.FieldEnv:
			values[i] = new(sql.NullString)
		case deploymentstatistics.FieldCreatedAt, deploymentstatistics.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeploymentStatistics", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeploymentStatistics fields.
func (ds *DeploymentStatistics) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deploymentstatistics.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ds.ID = int(value.Int64)
		case deploymentstatistics.FieldEnv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field env", values[i])
			} else if value.Valid {
				ds.Env = value.String
			}
		case deploymentstatistics.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				ds.Count = int(value.Int64)
			}
		case deploymentstatistics.FieldRollbackCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rollback_count", values[i])
			} else if value.Valid {
				ds.RollbackCount = int(value.Int64)
			}
		case deploymentstatistics.FieldAdditions:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field additions", values[i])
			} else if value.Valid {
				ds.Additions = int(value.Int64)
			}
		case deploymentstatistics.FieldDeletions:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deletions", values[i])
			} else if value.Valid {
				ds.Deletions = int(value.Int64)
			}
		case deploymentstatistics.FieldChanges:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field changes", values[i])
			} else if value.Valid {
				ds.Changes = int(value.Int64)
			}
		case deploymentstatistics.FieldLeadTimeSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lead_time_seconds", values[i])
			} else if value.Valid {
				ds.LeadTimeSeconds = int(value.Int64)
			}
		case deploymentstatistics.FieldCommitCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field commit_count", values[i])
			} else if value.Valid {
				ds.CommitCount = int(value.Int64)
			}
		case deploymentstatistics.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ds.CreatedAt = value.Time
			}
		case deploymentstatistics.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ds.UpdatedAt = value.Time
			}
		case deploymentstatistics.FieldRepoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field repo_id", values[i])
			} else if value.Valid {
				ds.RepoID = value.Int64
			}
		}
	}
	return nil
}

// QueryRepo queries the "repo" edge of the DeploymentStatistics entity.
func (ds *DeploymentStatistics) QueryRepo() *RepoQuery {
	return (&DeploymentStatisticsClient{config: ds.config}).QueryRepo(ds)
}

// Update returns a builder for updating this DeploymentStatistics.
// Note that you need to call DeploymentStatistics.Unwrap() before calling this method if this DeploymentStatistics
// was returned from a transaction, and the transaction was committed or rolled back.
func (ds *DeploymentStatistics) Update() *DeploymentStatisticsUpdateOne {
	return (&DeploymentStatisticsClient{config: ds.config}).UpdateOne(ds)
}

// Unwrap unwraps the DeploymentStatistics entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ds *DeploymentStatistics) Unwrap() *DeploymentStatistics {
	tx, ok := ds.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeploymentStatistics is not a transactional entity")
	}
	ds.config.driver = tx.drv
	return ds
}

// String implements the fmt.Stringer.
func (ds *DeploymentStatistics) String() string {
	var builder strings.Builder
	builder.WriteString("DeploymentStatistics(")
	builder.WriteString(fmt.Sprintf("id=%v", ds.ID))
	builder.WriteString(", env=")
	builder.WriteString(ds.Env)
	builder.WriteString(", count=")
	builder.WriteString(fmt.Sprintf("%v", ds.Count))
	builder.WriteString(", rollback_count=")
	builder.WriteString(fmt.Sprintf("%v", ds.RollbackCount))
	builder.WriteString(", additions=")
	builder.WriteString(fmt.Sprintf("%v", ds.Additions))
	builder.WriteString(", deletions=")
	builder.WriteString(fmt.Sprintf("%v", ds.Deletions))
	builder.WriteString(", changes=")
	builder.WriteString(fmt.Sprintf("%v", ds.Changes))
	builder.WriteString(", lead_time_seconds=")
	builder.WriteString(fmt.Sprintf("%v", ds.LeadTimeSeconds))
	builder.WriteString(", commit_count=")
	builder.WriteString(fmt.Sprintf("%v", ds.CommitCount))
	builder.WriteString(", created_at=")
	builder.WriteString(ds.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ds.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", repo_id=")
	builder.WriteString(fmt.Sprintf("%v", ds.RepoID))
	builder.WriteByte(')')
	return builder.String()
}

// DeploymentStatisticsSlice is a parsable slice of DeploymentStatistics.
type DeploymentStatisticsSlice []*DeploymentStatistics

func (ds DeploymentStatisticsSlice) config(cfg config) {
	for _i := range ds {
		ds[_i].config = cfg
	}
}