// Code generated by entc, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// DeploymentID applies equality check predicate on the "deployment_id" field. It's identical to DeploymentIDEQ.
func DeploymentID(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeploymentID), v))
	})
}

// DeploymentStatusID applies equality check predicate on the "deployment_status_id" field. It's identical to DeploymentStatusIDEQ.
func DeploymentStatusID(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeploymentStatusID), v))
	})
}

// ReviewID applies equality check predicate on the "review_id" field. It's identical to ReviewIDEQ.
func ReviewID(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReviewID), v))
	})
}

// DeletedID applies equality check predicate on the "deleted_id" field. It's identical to DeletedIDEQ.
func DeletedID(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedID), v))
	})
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v Kind) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKind), v))
	})
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v Kind) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKind), v))
	})
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...Kind) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKind), v...))
	})
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...Kind) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKind), v...))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
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
func TypeNotIn(vs ...Type) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// DeploymentIDEQ applies the EQ predicate on the "deployment_id" field.
func DeploymentIDEQ(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeploymentID), v))
	})
}

// DeploymentIDNEQ applies the NEQ predicate on the "deployment_id" field.
func DeploymentIDNEQ(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeploymentID), v))
	})
}

// DeploymentIDIn applies the In predicate on the "deployment_id" field.
func DeploymentIDIn(vs ...int) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeploymentID), v...))
	})
}

// DeploymentIDNotIn applies the NotIn predicate on the "deployment_id" field.
func DeploymentIDNotIn(vs ...int) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeploymentID), v...))
	})
}

// DeploymentIDIsNil applies the IsNil predicate on the "deployment_id" field.
func DeploymentIDIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeploymentID)))
	})
}

// DeploymentIDNotNil applies the NotNil predicate on the "deployment_id" field.
func DeploymentIDNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeploymentID)))
	})
}

// DeploymentStatusIDEQ applies the EQ predicate on the "deployment_status_id" field.
func DeploymentStatusIDEQ(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeploymentStatusID), v))
	})
}

// DeploymentStatusIDNEQ applies the NEQ predicate on the "deployment_status_id" field.
func DeploymentStatusIDNEQ(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeploymentStatusID), v))
	})
}

// DeploymentStatusIDIn applies the In predicate on the "deployment_status_id" field.
func DeploymentStatusIDIn(vs ...int) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeploymentStatusID), v...))
	})
}

// DeploymentStatusIDNotIn applies the NotIn predicate on the "deployment_status_id" field.
func DeploymentStatusIDNotIn(vs ...int) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeploymentStatusID), v...))
	})
}

// DeploymentStatusIDIsNil applies the IsNil predicate on the "deployment_status_id" field.
func DeploymentStatusIDIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeploymentStatusID)))
	})
}

// DeploymentStatusIDNotNil applies the NotNil predicate on the "deployment_status_id" field.
func DeploymentStatusIDNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeploymentStatusID)))
	})
}

// ReviewIDEQ applies the EQ predicate on the "review_id" field.
func ReviewIDEQ(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReviewID), v))
	})
}

// ReviewIDNEQ applies the NEQ predicate on the "review_id" field.
func ReviewIDNEQ(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReviewID), v))
	})
}

// ReviewIDIn applies the In predicate on the "review_id" field.
func ReviewIDIn(vs ...int) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReviewID), v...))
	})
}

// ReviewIDNotIn applies the NotIn predicate on the "review_id" field.
func ReviewIDNotIn(vs ...int) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReviewID), v...))
	})
}

// ReviewIDIsNil applies the IsNil predicate on the "review_id" field.
func ReviewIDIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReviewID)))
	})
}

// ReviewIDNotNil applies the NotNil predicate on the "review_id" field.
func ReviewIDNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReviewID)))
	})
}

// DeletedIDEQ applies the EQ predicate on the "deleted_id" field.
func DeletedIDEQ(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedID), v))
	})
}

// DeletedIDNEQ applies the NEQ predicate on the "deleted_id" field.
func DeletedIDNEQ(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedID), v))
	})
}

// DeletedIDIn applies the In predicate on the "deleted_id" field.
func DeletedIDIn(vs ...int) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedID), v...))
	})
}

// DeletedIDNotIn applies the NotIn predicate on the "deleted_id" field.
func DeletedIDNotIn(vs ...int) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedID), v...))
	})
}

// DeletedIDGT applies the GT predicate on the "deleted_id" field.
func DeletedIDGT(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedID), v))
	})
}

// DeletedIDGTE applies the GTE predicate on the "deleted_id" field.
func DeletedIDGTE(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedID), v))
	})
}

// DeletedIDLT applies the LT predicate on the "deleted_id" field.
func DeletedIDLT(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedID), v))
	})
}

// DeletedIDLTE applies the LTE predicate on the "deleted_id" field.
func DeletedIDLTE(v int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedID), v))
	})
}

// DeletedIDIsNil applies the IsNil predicate on the "deleted_id" field.
func DeletedIDIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedID)))
	})
}

// DeletedIDNotNil applies the NotNil predicate on the "deleted_id" field.
func DeletedIDNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedID)))
	})
}

// HasDeployment applies the HasEdge predicate on the "deployment" edge.
func HasDeployment() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentTable, DeploymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentWith applies the HasEdge predicate on the "deployment" edge with a given conditions (other predicates).
func HasDeploymentWith(preds ...predicate.Deployment) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentTable, DeploymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeploymentStatus applies the HasEdge predicate on the "deployment_status" edge.
func HasDeploymentStatus() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentStatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentStatusTable, DeploymentStatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentStatusWith applies the HasEdge predicate on the "deployment_status" edge with a given conditions (other predicates).
func HasDeploymentStatusWith(preds ...predicate.DeploymentStatus) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentStatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentStatusTable, DeploymentStatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReview applies the HasEdge predicate on the "review" edge.
func HasReview() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReviewTable, ReviewColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewWith applies the HasEdge predicate on the "review" edge with a given conditions (other predicates).
func HasReviewWith(preds ...predicate.Review) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReviewTable, ReviewColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotificationRecord applies the HasEdge predicate on the "notification_record" edge.
func HasNotificationRecord() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotificationRecordTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NotificationRecordTable, NotificationRecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationRecordWith applies the HasEdge predicate on the "notification_record" edge with a given conditions (other predicates).
func HasNotificationRecordWith(preds ...predicate.NotificationRecord) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotificationRecordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NotificationRecordTable, NotificationRecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		p(s.Not())
	})
}
