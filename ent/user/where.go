// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hanjunlee/gitploy/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDGT(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Login applies equality check predicate on the "login" field. It's identical to LoginEQ.
func Login(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogin), v))
	})
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// Admin applies equality check predicate on the "admin" field. It's identical to AdminEQ.
func Admin(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdmin), v))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// Refresh applies equality check predicate on the "refresh" field. It's identical to RefreshEQ.
func Refresh(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefresh), v))
	})
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// SyncedAt applies equality check predicate on the "synced_at" field. It's identical to SyncedAtEQ.
func SyncedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyncedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// LoginEQ applies the EQ predicate on the "login" field.
func LoginEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogin), v))
	})
}

// LoginNEQ applies the NEQ predicate on the "login" field.
func LoginNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogin), v))
	})
}

// LoginIn applies the In predicate on the "login" field.
func LoginIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogin), v...))
	})
}

// LoginNotIn applies the NotIn predicate on the "login" field.
func LoginNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogin), v...))
	})
}

// LoginGT applies the GT predicate on the "login" field.
func LoginGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogin), v))
	})
}

// LoginGTE applies the GTE predicate on the "login" field.
func LoginGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogin), v))
	})
}

// LoginLT applies the LT predicate on the "login" field.
func LoginLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogin), v))
	})
}

// LoginLTE applies the LTE predicate on the "login" field.
func LoginLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogin), v))
	})
}

// LoginContains applies the Contains predicate on the "login" field.
func LoginContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogin), v))
	})
}

// LoginHasPrefix applies the HasPrefix predicate on the "login" field.
func LoginHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogin), v))
	})
}

// LoginHasSuffix applies the HasSuffix predicate on the "login" field.
func LoginHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogin), v))
	})
}

// LoginEqualFold applies the EqualFold predicate on the "login" field.
func LoginEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogin), v))
	})
}

// LoginContainsFold applies the ContainsFold predicate on the "login" field.
func LoginContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogin), v))
	})
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvatar), v))
	})
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAvatar), v...))
	})
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAvatar), v...))
	})
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvatar), v))
	})
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvatar), v))
	})
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvatar), v))
	})
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvatar), v))
	})
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAvatar), v))
	})
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAvatar), v))
	})
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAvatar), v))
	})
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAvatar)))
	})
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAvatar)))
	})
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAvatar), v))
	})
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAvatar), v))
	})
}

// AdminEQ applies the EQ predicate on the "admin" field.
func AdminEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdmin), v))
	})
}

// AdminNEQ applies the NEQ predicate on the "admin" field.
func AdminNEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdmin), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToken), v))
	})
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToken), v))
	})
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToken), v))
	})
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToken), v))
	})
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToken), v))
	})
}

// RefreshEQ applies the EQ predicate on the "refresh" field.
func RefreshEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefresh), v))
	})
}

// RefreshNEQ applies the NEQ predicate on the "refresh" field.
func RefreshNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRefresh), v))
	})
}

// RefreshIn applies the In predicate on the "refresh" field.
func RefreshIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRefresh), v...))
	})
}

// RefreshNotIn applies the NotIn predicate on the "refresh" field.
func RefreshNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRefresh), v...))
	})
}

// RefreshGT applies the GT predicate on the "refresh" field.
func RefreshGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRefresh), v))
	})
}

// RefreshGTE applies the GTE predicate on the "refresh" field.
func RefreshGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRefresh), v))
	})
}

// RefreshLT applies the LT predicate on the "refresh" field.
func RefreshLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRefresh), v))
	})
}

// RefreshLTE applies the LTE predicate on the "refresh" field.
func RefreshLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRefresh), v))
	})
}

// RefreshContains applies the Contains predicate on the "refresh" field.
func RefreshContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRefresh), v))
	})
}

// RefreshHasPrefix applies the HasPrefix predicate on the "refresh" field.
func RefreshHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRefresh), v))
	})
}

// RefreshHasSuffix applies the HasSuffix predicate on the "refresh" field.
func RefreshHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRefresh), v))
	})
}

// RefreshEqualFold applies the EqualFold predicate on the "refresh" field.
func RefreshEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRefresh), v))
	})
}

// RefreshContainsFold applies the ContainsFold predicate on the "refresh" field.
func RefreshContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRefresh), v))
	})
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiry), v))
	})
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiry), v...))
	})
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiry), v...))
	})
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiry), v))
	})
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiry), v))
	})
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiry), v))
	})
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiry), v))
	})
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHash), v))
	})
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHash), v...))
	})
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHash), v...))
	})
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHash), v))
	})
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHash), v))
	})
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHash), v))
	})
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHash), v))
	})
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHash), v))
	})
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHash), v))
	})
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHash), v))
	})
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHash), v))
	})
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHash), v))
	})
}

// SyncedAtEQ applies the EQ predicate on the "synced_at" field.
func SyncedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtNEQ applies the NEQ predicate on the "synced_at" field.
func SyncedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtIn applies the In predicate on the "synced_at" field.
func SyncedAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
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
func SyncedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
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
func SyncedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtGTE applies the GTE predicate on the "synced_at" field.
func SyncedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtLT applies the LT predicate on the "synced_at" field.
func SyncedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtLTE applies the LTE predicate on the "synced_at" field.
func SyncedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSyncedAt), v))
	})
}

// SyncedAtIsNil applies the IsNil predicate on the "synced_at" field.
func SyncedAtIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSyncedAt)))
	})
}

// SyncedAtNotNil applies the NotNil predicate on the "synced_at" field.
func SyncedAtNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSyncedAt)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasChatUser applies the HasEdge predicate on the "chat_user" edge.
func HasChatUser() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChatUserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ChatUserTable, ChatUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChatUserWith applies the HasEdge predicate on the "chat_user" edge with a given conditions (other predicates).
func HasChatUserWith(preds ...predicate.ChatUser) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChatUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ChatUserTable, ChatUserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPerms applies the HasEdge predicate on the "perms" edge.
func HasPerms() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PermsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PermsTable, PermsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermsWith applies the HasEdge predicate on the "perms" edge with a given conditions (other predicates).
func HasPermsWith(preds ...predicate.Perm) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PermsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PermsTable, PermsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeployments applies the HasEdge predicate on the "deployments" edge.
func HasDeployments() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeploymentsTable, DeploymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentsWith applies the HasEdge predicate on the "deployments" edge with a given conditions (other predicates).
func HasDeploymentsWith(preds ...predicate.Deployment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeploymentsTable, DeploymentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasApprovals applies the HasEdge predicate on the "approvals" edge.
func HasApprovals() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ApprovalsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ApprovalsTable, ApprovalsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApprovalsWith applies the HasEdge predicate on the "approvals" edge with a given conditions (other predicates).
func HasApprovalsWith(preds ...predicate.Approval) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ApprovalsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ApprovalsTable, ApprovalsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
