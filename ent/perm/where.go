// Code generated by entc, DO NOT EDIT.

package perm

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gitploy-io/gitploy/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SyncedAt applies equality check predicate on the "synced_at" field. It's identical to SyncedAtEQ.
func SyncedAt(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyncedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// RepoID applies equality check predicate on the "repo_id" field. It's identical to RepoIDEQ.
func RepoID(v int64) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepoID), v))
	})
}

// RepoPermEQ applies the EQ predicate on the "repo_perm" field.
func RepoPermEQ(v RepoPerm) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepoPerm), v))
	})
}

// RepoPermNEQ applies the NEQ predicate on the "repo_perm" field.
func RepoPermNEQ(v RepoPerm) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRepoPerm), v))
	})
}

// RepoPermIn applies the In predicate on the "repo_perm" field.
func RepoPermIn(vs ...RepoPerm) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRepoPerm), v...))
	})
}

// RepoPermNotIn applies the NotIn predicate on the "repo_perm" field.
func RepoPermNotIn(vs ...RepoPerm) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRepoPerm), v...))
	})
}

// SyncedAtEQ applies the EQ predicate on the "synced_at" field.
func SyncedAtEQ(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtNEQ applies the NEQ predicate on the "synced_at" field.
func SyncedAtNEQ(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtIn applies the In predicate on the "synced_at" field.
func SyncedAtIn(vs ...time.Time) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSyncedAt), v...))
	})
}

// SyncedAtNotIn applies the NotIn predicate on the "synced_at" field.
func SyncedAtNotIn(vs ...time.Time) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSyncedAt), v...))
	})
}

// SyncedAtGT applies the GT predicate on the "synced_at" field.
func SyncedAtGT(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtGTE applies the GTE predicate on the "synced_at" field.
func SyncedAtGTE(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtLT applies the LT predicate on the "synced_at" field.
func SyncedAtLT(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtLTE applies the LTE predicate on the "synced_at" field.
func SyncedAtLTE(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSyncedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// RepoIDEQ applies the EQ predicate on the "repo_id" field.
func RepoIDEQ(v int64) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepoID), v))
	})
}

// RepoIDNEQ applies the NEQ predicate on the "repo_id" field.
func RepoIDNEQ(v int64) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRepoID), v))
	})
}

// RepoIDIn applies the In predicate on the "repo_id" field.
func RepoIDIn(vs ...int64) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
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
func RepoIDNotIn(vs ...int64) predicate.Perm {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Perm(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRepoID), v...))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRepo applies the HasEdge predicate on the "repo" edge.
func HasRepo() predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepoTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepoTable, RepoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepoWith applies the HasEdge predicate on the "repo" edge with a given conditions (other predicates).
func HasRepoWith(preds ...predicate.Repo) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
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
func And(predicates ...predicate.Perm) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Perm) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
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
func Not(p predicate.Perm) predicate.Perm {
	return predicate.Perm(func(s *sql.Selector) {
		p(s.Not())
	})
}
