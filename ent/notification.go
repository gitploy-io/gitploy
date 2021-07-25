// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hanjunlee/gitploy/ent/notification"
	"github.com/hanjunlee/gitploy/ent/user"
)

// Notification is the model entity for the Notification schema.
type Notification struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type notification.Type `json:"type"`
	// RepoNamespace holds the value of the "repo_namespace" field.
	RepoNamespace string `json:"repo_namespace"`
	// RepoName holds the value of the "repo_name" field.
	RepoName string `json:"repo_name"`
	// DeploymentNumber holds the value of the "deployment_number" field.
	DeploymentNumber int `json:"deployment_number"`
	// DeploymentType holds the value of the "deployment_type" field.
	DeploymentType string `json:"deployment_type"`
	// DeploymentRef holds the value of the "deployment_ref" field.
	DeploymentRef string `json:"deployment_ref"`
	// DeploymentEnv holds the value of the "deployment_env" field.
	DeploymentEnv string `json:"deployment_env"`
	// DeploymentStatus holds the value of the "deployment_status" field.
	DeploymentStatus string `json:"deployment_status"`
	// DeploymentLogin holds the value of the "deployment_login" field.
	DeploymentLogin string `json:"deployment_login"`
	// ApprovalStatus holds the value of the "approval_status" field.
	ApprovalStatus string `json:"approval_status"`
	// Notified holds the value of the "notified" field.
	Notified bool `json:"notified"`
	// Checked holds the value of the "checked" field.
	Checked bool `json:"checked"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotificationQuery when eager-loading is set.
	Edges NotificationEdges `json:"edges"`
}

// NotificationEdges holds the relations/edges for other nodes in the graph.
type NotificationEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotificationEdges) UserOrErr() (*User, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*Notification) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case notification.FieldNotified, notification.FieldChecked:
			values[i] = new(sql.NullBool)
		case notification.FieldID, notification.FieldDeploymentNumber:
			values[i] = new(sql.NullInt64)
		case notification.FieldType, notification.FieldRepoNamespace, notification.FieldRepoName, notification.FieldDeploymentType, notification.FieldDeploymentRef, notification.FieldDeploymentEnv, notification.FieldDeploymentStatus, notification.FieldDeploymentLogin, notification.FieldApprovalStatus, notification.FieldUserID:
			values[i] = new(sql.NullString)
		case notification.FieldCreatedAt, notification.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Notification", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notification fields.
func (n *Notification) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notification.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case notification.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				n.Type = notification.Type(value.String)
			}
		case notification.FieldRepoNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repo_namespace", values[i])
			} else if value.Valid {
				n.RepoNamespace = value.String
			}
		case notification.FieldRepoName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repo_name", values[i])
			} else if value.Valid {
				n.RepoName = value.String
			}
		case notification.FieldDeploymentNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_number", values[i])
			} else if value.Valid {
				n.DeploymentNumber = int(value.Int64)
			}
		case notification.FieldDeploymentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_type", values[i])
			} else if value.Valid {
				n.DeploymentType = value.String
			}
		case notification.FieldDeploymentRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_ref", values[i])
			} else if value.Valid {
				n.DeploymentRef = value.String
			}
		case notification.FieldDeploymentEnv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_env", values[i])
			} else if value.Valid {
				n.DeploymentEnv = value.String
			}
		case notification.FieldDeploymentStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_status", values[i])
			} else if value.Valid {
				n.DeploymentStatus = value.String
			}
		case notification.FieldDeploymentLogin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_login", values[i])
			} else if value.Valid {
				n.DeploymentLogin = value.String
			}
		case notification.FieldApprovalStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field approval_status", values[i])
			} else if value.Valid {
				n.ApprovalStatus = value.String
			}
		case notification.FieldNotified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field notified", values[i])
			} else if value.Valid {
				n.Notified = value.Bool
			}
		case notification.FieldChecked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field checked", values[i])
			} else if value.Valid {
				n.Checked = value.Bool
			}
		case notification.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case notification.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case notification.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				n.UserID = value.String
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Notification entity.
func (n *Notification) QueryUser() *UserQuery {
	return (&NotificationClient{config: n.config}).QueryUser(n)
}

// Update returns a builder for updating this Notification.
// Note that you need to call Notification.Unwrap() before calling this method if this Notification
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notification) Update() *NotificationUpdateOne {
	return (&NotificationClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Notification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Notification) Unwrap() *Notification {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Notification is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notification) String() string {
	var builder strings.Builder
	builder.WriteString("Notification(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", n.Type))
	builder.WriteString(", repo_namespace=")
	builder.WriteString(n.RepoNamespace)
	builder.WriteString(", repo_name=")
	builder.WriteString(n.RepoName)
	builder.WriteString(", deployment_number=")
	builder.WriteString(fmt.Sprintf("%v", n.DeploymentNumber))
	builder.WriteString(", deployment_type=")
	builder.WriteString(n.DeploymentType)
	builder.WriteString(", deployment_ref=")
	builder.WriteString(n.DeploymentRef)
	builder.WriteString(", deployment_env=")
	builder.WriteString(n.DeploymentEnv)
	builder.WriteString(", deployment_status=")
	builder.WriteString(n.DeploymentStatus)
	builder.WriteString(", deployment_login=")
	builder.WriteString(n.DeploymentLogin)
	builder.WriteString(", approval_status=")
	builder.WriteString(n.ApprovalStatus)
	builder.WriteString(", notified=")
	builder.WriteString(fmt.Sprintf("%v", n.Notified))
	builder.WriteString(", checked=")
	builder.WriteString(fmt.Sprintf("%v", n.Checked))
	builder.WriteString(", created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(n.UserID)
	builder.WriteByte(')')
	return builder.String()
}

// Notifications is a parsable slice of Notification.
type Notifications []*Notification

func (n Notifications) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
