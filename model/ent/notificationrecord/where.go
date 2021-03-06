// Code generated by entc, DO NOT EDIT.

package notificationrecord

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
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
func IDGT(id int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventID), v))
	})
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventID), v))
	})
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v int) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventID), v))
	})
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...int) predicate.NotificationRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotificationRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventID), v...))
	})
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...int) predicate.NotificationRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotificationRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventID), v...))
	})
}

// EventIDIsNil applies the IsNil predicate on the "event_id" field.
func EventIDIsNil() predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEventID)))
	})
}

// EventIDNotNil applies the NotNil predicate on the "event_id" field.
func EventIDNotNil() predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEventID)))
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotificationRecord) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotificationRecord) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
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
func Not(p predicate.NotificationRecord) predicate.NotificationRecord {
	return predicate.NotificationRecord(func(s *sql.Selector) {
		p(s.Not())
	})
}
