// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitploy-io/gitploy/model/ent/chatuser"
	"github.com/gitploy-io/gitploy/model/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Login holds the value of the "login" field.
	Login string `json:"login"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar"`
	// Admin holds the value of the "admin" field.
	Admin bool `json:"admin"`
	// Token holds the value of the "token" field.
	Token string `json:"-"`
	// Refresh holds the value of the "refresh" field.
	Refresh string `json:"-"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// ChatUser holds the value of the chat_user edge.
	ChatUser *ChatUser `json:"chat_user,omitempty"`
	// Perms holds the value of the perms edge.
	Perms []*Perm `json:"perms,omitempty"`
	// Deployments holds the value of the deployments edge.
	Deployments []*Deployment `json:"deployments,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// Locks holds the value of the locks edge.
	Locks []*Lock `json:"locks,omitempty"`
	// Repos holds the value of the repos edge.
	Repos []*Repo `json:"repos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ChatUserOrErr returns the ChatUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ChatUserOrErr() (*ChatUser, error) {
	if e.loadedTypes[0] {
		if e.ChatUser == nil {
			// The edge chat_user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: chatuser.Label}
		}
		return e.ChatUser, nil
	}
	return nil, &NotLoadedError{edge: "chat_user"}
}

// PermsOrErr returns the Perms value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PermsOrErr() ([]*Perm, error) {
	if e.loadedTypes[1] {
		return e.Perms, nil
	}
	return nil, &NotLoadedError{edge: "perms"}
}

// DeploymentsOrErr returns the Deployments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DeploymentsOrErr() ([]*Deployment, error) {
	if e.loadedTypes[2] {
		return e.Deployments, nil
	}
	return nil, &NotLoadedError{edge: "deployments"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[3] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// LocksOrErr returns the Locks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LocksOrErr() ([]*Lock, error) {
	if e.loadedTypes[4] {
		return e.Locks, nil
	}
	return nil, &NotLoadedError{edge: "locks"}
}

// ReposOrErr returns the Repos value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReposOrErr() ([]*Repo, error) {
	if e.loadedTypes[5] {
		return e.Repos, nil
	}
	return nil, &NotLoadedError{edge: "repos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldAdmin:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldLogin, user.FieldAvatar, user.FieldToken, user.FieldRefresh, user.FieldHash:
			values[i] = new(sql.NullString)
		case user.FieldExpiry, user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldLogin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login", values[i])
			} else if value.Valid {
				u.Login = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field admin", values[i])
			} else if value.Valid {
				u.Admin = value.Bool
			}
		case user.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				u.Token = value.String
			}
		case user.FieldRefresh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh", values[i])
			} else if value.Valid {
				u.Refresh = value.String
			}
		case user.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				u.Expiry = value.Time
			}
		case user.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				u.Hash = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryChatUser queries the "chat_user" edge of the User entity.
func (u *User) QueryChatUser() *ChatUserQuery {
	return (&UserClient{config: u.config}).QueryChatUser(u)
}

// QueryPerms queries the "perms" edge of the User entity.
func (u *User) QueryPerms() *PermQuery {
	return (&UserClient{config: u.config}).QueryPerms(u)
}

// QueryDeployments queries the "deployments" edge of the User entity.
func (u *User) QueryDeployments() *DeploymentQuery {
	return (&UserClient{config: u.config}).QueryDeployments(u)
}

// QueryReviews queries the "reviews" edge of the User entity.
func (u *User) QueryReviews() *ReviewQuery {
	return (&UserClient{config: u.config}).QueryReviews(u)
}

// QueryLocks queries the "locks" edge of the User entity.
func (u *User) QueryLocks() *LockQuery {
	return (&UserClient{config: u.config}).QueryLocks(u)
}

// QueryRepos queries the "repos" edge of the User entity.
func (u *User) QueryRepos() *RepoQuery {
	return (&UserClient{config: u.config}).QueryRepos(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", login=")
	builder.WriteString(u.Login)
	builder.WriteString(", avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", admin=")
	builder.WriteString(fmt.Sprintf("%v", u.Admin))
	builder.WriteString(", token=<sensitive>")
	builder.WriteString(", refresh=<sensitive>")
	builder.WriteString(", expiry=")
	builder.WriteString(u.Expiry.Format(time.ANSIC))
	builder.WriteString(", hash=<sensitive>")
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
