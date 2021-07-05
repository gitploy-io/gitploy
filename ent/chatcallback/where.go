// Code generated by entc, DO NOT EDIT.

package chatcallback

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hanjunlee/gitploy/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// IsOpened applies equality check predicate on the "is_opened" field. It's identical to IsOpenedEQ.
func IsOpened(v bool) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsOpened), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ChatUserID applies equality check predicate on the "chat_user_id" field. It's identical to ChatUserIDEQ.
func ChatUserID(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChatUserID), v))
	})
}

// RepoID applies equality check predicate on the "repo_id" field. It's identical to RepoIDEQ.
func RepoID(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepoID), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// IsOpenedEQ applies the EQ predicate on the "is_opened" field.
func IsOpenedEQ(v bool) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsOpened), v))
	})
}

// IsOpenedNEQ applies the NEQ predicate on the "is_opened" field.
func IsOpenedNEQ(v bool) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsOpened), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// ChatUserIDEQ applies the EQ predicate on the "chat_user_id" field.
func ChatUserIDEQ(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDNEQ applies the NEQ predicate on the "chat_user_id" field.
func ChatUserIDNEQ(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDIn applies the In predicate on the "chat_user_id" field.
func ChatUserIDIn(vs ...string) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChatUserID), v...))
	})
}

// ChatUserIDNotIn applies the NotIn predicate on the "chat_user_id" field.
func ChatUserIDNotIn(vs ...string) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChatUserID), v...))
	})
}

// ChatUserIDGT applies the GT predicate on the "chat_user_id" field.
func ChatUserIDGT(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDGTE applies the GTE predicate on the "chat_user_id" field.
func ChatUserIDGTE(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDLT applies the LT predicate on the "chat_user_id" field.
func ChatUserIDLT(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDLTE applies the LTE predicate on the "chat_user_id" field.
func ChatUserIDLTE(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDContains applies the Contains predicate on the "chat_user_id" field.
func ChatUserIDContains(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDHasPrefix applies the HasPrefix predicate on the "chat_user_id" field.
func ChatUserIDHasPrefix(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDHasSuffix applies the HasSuffix predicate on the "chat_user_id" field.
func ChatUserIDHasSuffix(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDEqualFold applies the EqualFold predicate on the "chat_user_id" field.
func ChatUserIDEqualFold(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChatUserID), v))
	})
}

// ChatUserIDContainsFold applies the ContainsFold predicate on the "chat_user_id" field.
func ChatUserIDContainsFold(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChatUserID), v))
	})
}

// RepoIDEQ applies the EQ predicate on the "repo_id" field.
func RepoIDEQ(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepoID), v))
	})
}

// RepoIDNEQ applies the NEQ predicate on the "repo_id" field.
func RepoIDNEQ(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRepoID), v))
	})
}

// RepoIDIn applies the In predicate on the "repo_id" field.
func RepoIDIn(vs ...string) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRepoID), v...))
	})
}

// RepoIDNotIn applies the NotIn predicate on the "repo_id" field.
func RepoIDNotIn(vs ...string) predicate.ChatCallback {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatCallback(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRepoID), v...))
	})
}

// RepoIDGT applies the GT predicate on the "repo_id" field.
func RepoIDGT(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRepoID), v))
	})
}

// RepoIDGTE applies the GTE predicate on the "repo_id" field.
func RepoIDGTE(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRepoID), v))
	})
}

// RepoIDLT applies the LT predicate on the "repo_id" field.
func RepoIDLT(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRepoID), v))
	})
}

// RepoIDLTE applies the LTE predicate on the "repo_id" field.
func RepoIDLTE(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRepoID), v))
	})
}

// RepoIDContains applies the Contains predicate on the "repo_id" field.
func RepoIDContains(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRepoID), v))
	})
}

// RepoIDHasPrefix applies the HasPrefix predicate on the "repo_id" field.
func RepoIDHasPrefix(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRepoID), v))
	})
}

// RepoIDHasSuffix applies the HasSuffix predicate on the "repo_id" field.
func RepoIDHasSuffix(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRepoID), v))
	})
}

// RepoIDIsNil applies the IsNil predicate on the "repo_id" field.
func RepoIDIsNil() predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRepoID)))
	})
}

// RepoIDNotNil applies the NotNil predicate on the "repo_id" field.
func RepoIDNotNil() predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRepoID)))
	})
}

// RepoIDEqualFold applies the EqualFold predicate on the "repo_id" field.
func RepoIDEqualFold(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRepoID), v))
	})
}

// RepoIDContainsFold applies the ContainsFold predicate on the "repo_id" field.
func RepoIDContainsFold(v string) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRepoID), v))
	})
}

// HasChatUser applies the HasEdge predicate on the "chat_user" edge.
func HasChatUser() predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChatUserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChatUserTable, ChatUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChatUserWith applies the HasEdge predicate on the "chat_user" edge with a given conditions (other predicates).
func HasChatUserWith(preds ...predicate.ChatUser) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChatUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChatUserTable, ChatUserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRepo applies the HasEdge predicate on the "repo" edge.
func HasRepo() predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepoTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepoTable, RepoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepoWith applies the HasEdge predicate on the "repo" edge with a given conditions (other predicates).
func HasRepoWith(preds ...predicate.Repo) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepoInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepoTable, RepoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChatCallback) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChatCallback) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChatCallback) predicate.ChatCallback {
	return predicate.ChatCallback(func(s *sql.Selector) {
		p(s.Not())
	})
}
