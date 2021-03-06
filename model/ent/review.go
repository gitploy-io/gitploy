// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/review"
	"github.com/gitploy-io/gitploy/model/ent/user"
)

// Review is the model entity for the Review schema.
type Review struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status review.Status `json:"status"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitemtpy"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitemtpy"`
	// DeploymentID holds the value of the "deployment_id" field.
	DeploymentID int `json:"deployment_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewQuery when eager-loading is set.
	Edges ReviewEdges `json:"edges"`
}

// ReviewEdges holds the relations/edges for other nodes in the graph.
type ReviewEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Deployment holds the value of the deployment edge.
	Deployment *Deployment `json:"deployment,omitempty"`
	// Event holds the value of the event edge.
	Event []*Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewEdges) UserOrErr() (*User, error) {
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

// DeploymentOrErr returns the Deployment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewEdges) DeploymentOrErr() (*Deployment, error) {
	if e.loadedTypes[1] {
		if e.Deployment == nil {
			// The edge deployment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: deployment.Label}
		}
		return e.Deployment, nil
	}
	return nil, &NotLoadedError{edge: "deployment"}
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading.
func (e ReviewEdges) EventOrErr() ([]*Event, error) {
	if e.loadedTypes[2] {
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Review) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case review.FieldID, review.FieldUserID, review.FieldDeploymentID:
			values[i] = new(sql.NullInt64)
		case review.FieldStatus, review.FieldComment:
			values[i] = new(sql.NullString)
		case review.FieldCreatedAt, review.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Review", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Review fields.
func (r *Review) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case review.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case review.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = review.Status(value.String)
			}
		case review.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				r.Comment = value.String
			}
		case review.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case review.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case review.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				r.UserID = value.Int64
			}
		case review.FieldDeploymentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_id", values[i])
			} else if value.Valid {
				r.DeploymentID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Review entity.
func (r *Review) QueryUser() *UserQuery {
	return (&ReviewClient{config: r.config}).QueryUser(r)
}

// QueryDeployment queries the "deployment" edge of the Review entity.
func (r *Review) QueryDeployment() *DeploymentQuery {
	return (&ReviewClient{config: r.config}).QueryDeployment(r)
}

// QueryEvent queries the "event" edge of the Review entity.
func (r *Review) QueryEvent() *EventQuery {
	return (&ReviewClient{config: r.config}).QueryEvent(r)
}

// Update returns a builder for updating this Review.
// Note that you need to call Review.Unwrap() before calling this method if this Review
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Review) Update() *ReviewUpdateOne {
	return (&ReviewClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Review entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Review) Unwrap() *Review {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Review is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Review) String() string {
	var builder strings.Builder
	builder.WriteString("Review(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteString(", comment=")
	builder.WriteString(r.Comment)
	builder.WriteString(", created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", deployment_id=")
	builder.WriteString(fmt.Sprintf("%v", r.DeploymentID))
	builder.WriteByte(')')
	return builder.String()
}

// Reviews is a parsable slice of Review.
type Reviews []*Review

func (r Reviews) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
