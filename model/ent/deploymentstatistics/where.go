// Code generated by entc, DO NOT EDIT.

package deploymentstatistics

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func IDGT(id int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Env applies equality check predicate on the "env" field. It's identical to EnvEQ.
func Env(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnv), v))
	})
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// RollbackCount applies equality check predicate on the "rollback_count" field. It's identical to RollbackCountEQ.
func RollbackCount(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRollbackCount), v))
	})
}

// Additions applies equality check predicate on the "additions" field. It's identical to AdditionsEQ.
func Additions(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdditions), v))
	})
}

// Deletions applies equality check predicate on the "deletions" field. It's identical to DeletionsEQ.
func Deletions(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletions), v))
	})
}

// Changes applies equality check predicate on the "changes" field. It's identical to ChangesEQ.
func Changes(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChanges), v))
	})
}

// LeadTimeSeconds applies equality check predicate on the "lead_time_seconds" field. It's identical to LeadTimeSecondsEQ.
func LeadTimeSeconds(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeadTimeSeconds), v))
	})
}

// CommitCount applies equality check predicate on the "commit_count" field. It's identical to CommitCountEQ.
func CommitCount(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommitCount), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// RepoID applies equality check predicate on the "repo_id" field. It's identical to RepoIDEQ.
func RepoID(v int64) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepoID), v))
	})
}

// EnvEQ applies the EQ predicate on the "env" field.
func EnvEQ(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnv), v))
	})
}

// EnvNEQ applies the NEQ predicate on the "env" field.
func EnvNEQ(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnv), v))
	})
}

// EnvIn applies the In predicate on the "env" field.
func EnvIn(vs ...string) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnv), v...))
	})
}

// EnvNotIn applies the NotIn predicate on the "env" field.
func EnvNotIn(vs ...string) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnv), v...))
	})
}

// EnvGT applies the GT predicate on the "env" field.
func EnvGT(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnv), v))
	})
}

// EnvGTE applies the GTE predicate on the "env" field.
func EnvGTE(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnv), v))
	})
}

// EnvLT applies the LT predicate on the "env" field.
func EnvLT(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnv), v))
	})
}

// EnvLTE applies the LTE predicate on the "env" field.
func EnvLTE(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnv), v))
	})
}

// EnvContains applies the Contains predicate on the "env" field.
func EnvContains(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEnv), v))
	})
}

// EnvHasPrefix applies the HasPrefix predicate on the "env" field.
func EnvHasPrefix(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEnv), v))
	})
}

// EnvHasSuffix applies the HasSuffix predicate on the "env" field.
func EnvHasSuffix(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEnv), v))
	})
}

// EnvEqualFold applies the EqualFold predicate on the "env" field.
func EnvEqualFold(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEnv), v))
	})
}

// EnvContainsFold applies the ContainsFold predicate on the "env" field.
func EnvContainsFold(v string) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEnv), v))
	})
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCount), v))
	})
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCount), v...))
	})
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCount), v...))
	})
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCount), v))
	})
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCount), v))
	})
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCount), v))
	})
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCount), v))
	})
}

// RollbackCountEQ applies the EQ predicate on the "rollback_count" field.
func RollbackCountEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRollbackCount), v))
	})
}

// RollbackCountNEQ applies the NEQ predicate on the "rollback_count" field.
func RollbackCountNEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRollbackCount), v))
	})
}

// RollbackCountIn applies the In predicate on the "rollback_count" field.
func RollbackCountIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRollbackCount), v...))
	})
}

// RollbackCountNotIn applies the NotIn predicate on the "rollback_count" field.
func RollbackCountNotIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRollbackCount), v...))
	})
}

// RollbackCountGT applies the GT predicate on the "rollback_count" field.
func RollbackCountGT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRollbackCount), v))
	})
}

// RollbackCountGTE applies the GTE predicate on the "rollback_count" field.
func RollbackCountGTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRollbackCount), v))
	})
}

// RollbackCountLT applies the LT predicate on the "rollback_count" field.
func RollbackCountLT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRollbackCount), v))
	})
}

// RollbackCountLTE applies the LTE predicate on the "rollback_count" field.
func RollbackCountLTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRollbackCount), v))
	})
}

// AdditionsEQ applies the EQ predicate on the "additions" field.
func AdditionsEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdditions), v))
	})
}

// AdditionsNEQ applies the NEQ predicate on the "additions" field.
func AdditionsNEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdditions), v))
	})
}

// AdditionsIn applies the In predicate on the "additions" field.
func AdditionsIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdditions), v...))
	})
}

// AdditionsNotIn applies the NotIn predicate on the "additions" field.
func AdditionsNotIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdditions), v...))
	})
}

// AdditionsGT applies the GT predicate on the "additions" field.
func AdditionsGT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdditions), v))
	})
}

// AdditionsGTE applies the GTE predicate on the "additions" field.
func AdditionsGTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdditions), v))
	})
}

// AdditionsLT applies the LT predicate on the "additions" field.
func AdditionsLT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdditions), v))
	})
}

// AdditionsLTE applies the LTE predicate on the "additions" field.
func AdditionsLTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdditions), v))
	})
}

// DeletionsEQ applies the EQ predicate on the "deletions" field.
func DeletionsEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletions), v))
	})
}

// DeletionsNEQ applies the NEQ predicate on the "deletions" field.
func DeletionsNEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletions), v))
	})
}

// DeletionsIn applies the In predicate on the "deletions" field.
func DeletionsIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletions), v...))
	})
}

// DeletionsNotIn applies the NotIn predicate on the "deletions" field.
func DeletionsNotIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletions), v...))
	})
}

// DeletionsGT applies the GT predicate on the "deletions" field.
func DeletionsGT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletions), v))
	})
}

// DeletionsGTE applies the GTE predicate on the "deletions" field.
func DeletionsGTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletions), v))
	})
}

// DeletionsLT applies the LT predicate on the "deletions" field.
func DeletionsLT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletions), v))
	})
}

// DeletionsLTE applies the LTE predicate on the "deletions" field.
func DeletionsLTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletions), v))
	})
}

// ChangesEQ applies the EQ predicate on the "changes" field.
func ChangesEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChanges), v))
	})
}

// ChangesNEQ applies the NEQ predicate on the "changes" field.
func ChangesNEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChanges), v))
	})
}

// ChangesIn applies the In predicate on the "changes" field.
func ChangesIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChanges), v...))
	})
}

// ChangesNotIn applies the NotIn predicate on the "changes" field.
func ChangesNotIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChanges), v...))
	})
}

// ChangesGT applies the GT predicate on the "changes" field.
func ChangesGT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChanges), v))
	})
}

// ChangesGTE applies the GTE predicate on the "changes" field.
func ChangesGTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChanges), v))
	})
}

// ChangesLT applies the LT predicate on the "changes" field.
func ChangesLT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChanges), v))
	})
}

// ChangesLTE applies the LTE predicate on the "changes" field.
func ChangesLTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChanges), v))
	})
}

// LeadTimeSecondsEQ applies the EQ predicate on the "lead_time_seconds" field.
func LeadTimeSecondsEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeadTimeSeconds), v))
	})
}

// LeadTimeSecondsNEQ applies the NEQ predicate on the "lead_time_seconds" field.
func LeadTimeSecondsNEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLeadTimeSeconds), v))
	})
}

// LeadTimeSecondsIn applies the In predicate on the "lead_time_seconds" field.
func LeadTimeSecondsIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLeadTimeSeconds), v...))
	})
}

// LeadTimeSecondsNotIn applies the NotIn predicate on the "lead_time_seconds" field.
func LeadTimeSecondsNotIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLeadTimeSeconds), v...))
	})
}

// LeadTimeSecondsGT applies the GT predicate on the "lead_time_seconds" field.
func LeadTimeSecondsGT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLeadTimeSeconds), v))
	})
}

// LeadTimeSecondsGTE applies the GTE predicate on the "lead_time_seconds" field.
func LeadTimeSecondsGTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLeadTimeSeconds), v))
	})
}

// LeadTimeSecondsLT applies the LT predicate on the "lead_time_seconds" field.
func LeadTimeSecondsLT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLeadTimeSeconds), v))
	})
}

// LeadTimeSecondsLTE applies the LTE predicate on the "lead_time_seconds" field.
func LeadTimeSecondsLTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLeadTimeSeconds), v))
	})
}

// CommitCountEQ applies the EQ predicate on the "commit_count" field.
func CommitCountEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommitCount), v))
	})
}

// CommitCountNEQ applies the NEQ predicate on the "commit_count" field.
func CommitCountNEQ(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCommitCount), v))
	})
}

// CommitCountIn applies the In predicate on the "commit_count" field.
func CommitCountIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCommitCount), v...))
	})
}

// CommitCountNotIn applies the NotIn predicate on the "commit_count" field.
func CommitCountNotIn(vs ...int) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCommitCount), v...))
	})
}

// CommitCountGT applies the GT predicate on the "commit_count" field.
func CommitCountGT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCommitCount), v))
	})
}

// CommitCountGTE applies the GTE predicate on the "commit_count" field.
func CommitCountGTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCommitCount), v))
	})
}

// CommitCountLT applies the LT predicate on the "commit_count" field.
func CommitCountLT(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCommitCount), v))
	})
}

// CommitCountLTE applies the LTE predicate on the "commit_count" field.
func CommitCountLTE(v int) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCommitCount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// RepoIDEQ applies the EQ predicate on the "repo_id" field.
func RepoIDEQ(v int64) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepoID), v))
	})
}

// RepoIDNEQ applies the NEQ predicate on the "repo_id" field.
func RepoIDNEQ(v int64) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRepoID), v))
	})
}

// RepoIDIn applies the In predicate on the "repo_id" field.
func RepoIDIn(vs ...int64) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func RepoIDNotIn(vs ...int64) predicate.DeploymentStatistics {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRepoID), v...))
	})
}

// HasRepo applies the HasEdge predicate on the "repo" edge.
func HasRepo() predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepoTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepoTable, RepoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepoWith applies the HasEdge predicate on the "repo" edge with a given conditions (other predicates).
func HasRepoWith(preds ...predicate.Repo) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func And(predicates ...predicate.DeploymentStatistics) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeploymentStatistics) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
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
func Not(p predicate.DeploymentStatistics) predicate.DeploymentStatistics {
	return predicate.DeploymentStatistics(func(s *sql.Selector) {
		p(s.Not())
	})
}
