// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/chatuser"
	"github.com/hanjunlee/gitploy/ent/repo"
)

// ChatCallback is the model entity for the ChatCallback schema.
type ChatCallback struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state"`
	// Type holds the value of the "type" field.
	Type chatcallback.Type `json:"type"`
	// IsOpened holds the value of the "is_opened" field.
	IsOpened bool `json:"is_opened"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// ChatUserID holds the value of the "chat_user_id" field.
	ChatUserID string `json:"chat_user_id"`
	// RepoID holds the value of the "repo_id" field.
	RepoID string `json:"repo_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChatCallbackQuery when eager-loading is set.
	Edges ChatCallbackEdges `json:"edges"`
}

// ChatCallbackEdges holds the relations/edges for other nodes in the graph.
type ChatCallbackEdges struct {
	// ChatUser holds the value of the chat_user edge.
	ChatUser *ChatUser `json:"chat_user,omitempty"`
	// Repo holds the value of the repo edge.
	Repo *Repo `json:"repo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChatUserOrErr returns the ChatUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatCallbackEdges) ChatUserOrErr() (*ChatUser, error) {
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

// RepoOrErr returns the Repo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatCallbackEdges) RepoOrErr() (*Repo, error) {
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
func (*ChatCallback) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case chatcallback.FieldIsOpened:
			values[i] = new(sql.NullBool)
		case chatcallback.FieldID:
			values[i] = new(sql.NullInt64)
		case chatcallback.FieldState, chatcallback.FieldType, chatcallback.FieldChatUserID, chatcallback.FieldRepoID:
			values[i] = new(sql.NullString)
		case chatcallback.FieldCreatedAt, chatcallback.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ChatCallback", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChatCallback fields.
func (cc *ChatCallback) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chatcallback.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cc.ID = int(value.Int64)
		case chatcallback.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				cc.State = value.String
			}
		case chatcallback.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				cc.Type = chatcallback.Type(value.String)
			}
		case chatcallback.FieldIsOpened:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_opened", values[i])
			} else if value.Valid {
				cc.IsOpened = value.Bool
			}
		case chatcallback.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cc.CreatedAt = value.Time
			}
		case chatcallback.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cc.UpdatedAt = value.Time
			}
		case chatcallback.FieldChatUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chat_user_id", values[i])
			} else if value.Valid {
				cc.ChatUserID = value.String
			}
		case chatcallback.FieldRepoID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repo_id", values[i])
			} else if value.Valid {
				cc.RepoID = value.String
			}
		}
	}
	return nil
}

// QueryChatUser queries the "chat_user" edge of the ChatCallback entity.
func (cc *ChatCallback) QueryChatUser() *ChatUserQuery {
	return (&ChatCallbackClient{config: cc.config}).QueryChatUser(cc)
}

// QueryRepo queries the "repo" edge of the ChatCallback entity.
func (cc *ChatCallback) QueryRepo() *RepoQuery {
	return (&ChatCallbackClient{config: cc.config}).QueryRepo(cc)
}

// Update returns a builder for updating this ChatCallback.
// Note that you need to call ChatCallback.Unwrap() before calling this method if this ChatCallback
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *ChatCallback) Update() *ChatCallbackUpdateOne {
	return (&ChatCallbackClient{config: cc.config}).UpdateOne(cc)
}

// Unwrap unwraps the ChatCallback entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *ChatCallback) Unwrap() *ChatCallback {
	tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChatCallback is not a transactional entity")
	}
	cc.config.driver = tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *ChatCallback) String() string {
	var builder strings.Builder
	builder.WriteString("ChatCallback(")
	builder.WriteString(fmt.Sprintf("id=%v", cc.ID))
	builder.WriteString(", state=")
	builder.WriteString(cc.State)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", cc.Type))
	builder.WriteString(", is_opened=")
	builder.WriteString(fmt.Sprintf("%v", cc.IsOpened))
	builder.WriteString(", created_at=")
	builder.WriteString(cc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(cc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", chat_user_id=")
	builder.WriteString(cc.ChatUserID)
	builder.WriteString(", repo_id=")
	builder.WriteString(cc.RepoID)
	builder.WriteByte(')')
	return builder.String()
}

// ChatCallbacks is a parsable slice of ChatCallback.
type ChatCallbacks []*ChatCallback

func (cc ChatCallbacks) config(cfg config) {
	for _i := range cc {
		cc[_i].config = cfg
	}
}
