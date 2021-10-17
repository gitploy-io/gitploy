// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitploy-io/gitploy/ent/lock"
	"github.com/gitploy-io/gitploy/ent/repo"
	"github.com/gitploy-io/gitploy/ent/user"
)

// Lock is the model entity for the Lock schema.
type Lock struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Env holds the value of the "env" field.
	Env string `json:"env"`
	// ExpiredAt holds the value of the "expired_at" field.
	ExpiredAt time.Time `json:"expired_at,omitemtpy"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id"`
	// RepoID holds the value of the "repo_id" field.
	RepoID int64 `json:"repo_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LockQuery when eager-loading is set.
	Edges LockEdges `json:"edges"`
}

// LockEdges holds the relations/edges for other nodes in the graph.
type LockEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Repo holds the value of the repo edge.
	Repo *Repo `json:"repo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LockEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RepoOrErr returns the Repo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LockEdges) RepoOrErr() (*Repo, error) {
	if e.loadedTypes[1] {
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
func (*Lock) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case lock.FieldID, lock.FieldUserID, lock.FieldRepoID:
			values[i] = new(sql.NullInt64)
		case lock.FieldEnv:
			values[i] = new(sql.NullString)
		case lock.FieldExpiredAt, lock.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Lock", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Lock fields.
func (l *Lock) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lock.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case lock.FieldEnv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field env", values[i])
			} else if value.Valid {
				l.Env = value.String
			}
		case lock.FieldExpiredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expired_at", values[i])
			} else if value.Valid {
				l.ExpiredAt = value.Time
			}
		case lock.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case lock.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				l.UserID = value.Int64
			}
		case lock.FieldRepoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field repo_id", values[i])
			} else if value.Valid {
				l.RepoID = value.Int64
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Lock entity.
func (l *Lock) QueryUser() *UserQuery {
	return (&LockClient{config: l.config}).QueryUser(l)
}

// QueryRepo queries the "repo" edge of the Lock entity.
func (l *Lock) QueryRepo() *RepoQuery {
	return (&LockClient{config: l.config}).QueryRepo(l)
}

// Update returns a builder for updating this Lock.
// Note that you need to call Lock.Unwrap() before calling this method if this Lock
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Lock) Update() *LockUpdateOne {
	return (&LockClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Lock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Lock) Unwrap() *Lock {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Lock is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Lock) String() string {
	var builder strings.Builder
	builder.WriteString("Lock(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", env=")
	builder.WriteString(l.Env)
	builder.WriteString(", expired_at=")
	builder.WriteString(l.ExpiredAt.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", l.UserID))
	builder.WriteString(", repo_id=")
	builder.WriteString(fmt.Sprintf("%v", l.RepoID))
	builder.WriteByte(')')
	return builder.String()
}

// Locks is a parsable slice of Lock.
type Locks []*Lock

func (l Locks) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
