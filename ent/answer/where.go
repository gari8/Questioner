// Code generated by entc, DO NOT EDIT.

package answer

import (
	"faves4/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AnswerType applies equality check predicate on the "answer_type" field. It's identical to AnswerTypeEQ.
func AnswerType(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnswerType), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// CratedAt applies equality check predicate on the "crated_at" field. It's identical to CratedAtEQ.
func CratedAt(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCratedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// AnswerTypeEQ applies the EQ predicate on the "answer_type" field.
func AnswerTypeEQ(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeNEQ applies the NEQ predicate on the "answer_type" field.
func AnswerTypeNEQ(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeIn applies the In predicate on the "answer_type" field.
func AnswerTypeIn(vs ...string) predicate.Answer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Answer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAnswerType), v...))
	})
}

// AnswerTypeNotIn applies the NotIn predicate on the "answer_type" field.
func AnswerTypeNotIn(vs ...string) predicate.Answer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Answer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAnswerType), v...))
	})
}

// AnswerTypeGT applies the GT predicate on the "answer_type" field.
func AnswerTypeGT(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeGTE applies the GTE predicate on the "answer_type" field.
func AnswerTypeGTE(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeLT applies the LT predicate on the "answer_type" field.
func AnswerTypeLT(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeLTE applies the LTE predicate on the "answer_type" field.
func AnswerTypeLTE(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeContains applies the Contains predicate on the "answer_type" field.
func AnswerTypeContains(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeHasPrefix applies the HasPrefix predicate on the "answer_type" field.
func AnswerTypeHasPrefix(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeHasSuffix applies the HasSuffix predicate on the "answer_type" field.
func AnswerTypeHasSuffix(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeIsNil applies the IsNil predicate on the "answer_type" field.
func AnswerTypeIsNil() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAnswerType)))
	})
}

// AnswerTypeNotNil applies the NotNil predicate on the "answer_type" field.
func AnswerTypeNotNil() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAnswerType)))
	})
}

// AnswerTypeEqualFold applies the EqualFold predicate on the "answer_type" field.
func AnswerTypeEqualFold(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAnswerType), v))
	})
}

// AnswerTypeContainsFold applies the ContainsFold predicate on the "answer_type" field.
func AnswerTypeContainsFold(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAnswerType), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Answer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Answer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Answer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Answer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// CratedAtEQ applies the EQ predicate on the "crated_at" field.
func CratedAtEQ(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCratedAt), v))
	})
}

// CratedAtNEQ applies the NEQ predicate on the "crated_at" field.
func CratedAtNEQ(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCratedAt), v))
	})
}

// CratedAtIn applies the In predicate on the "crated_at" field.
func CratedAtIn(vs ...time.Time) predicate.Answer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Answer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCratedAt), v...))
	})
}

// CratedAtNotIn applies the NotIn predicate on the "crated_at" field.
func CratedAtNotIn(vs ...time.Time) predicate.Answer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Answer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCratedAt), v...))
	})
}

// CratedAtGT applies the GT predicate on the "crated_at" field.
func CratedAtGT(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCratedAt), v))
	})
}

// CratedAtGTE applies the GTE predicate on the "crated_at" field.
func CratedAtGTE(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCratedAt), v))
	})
}

// CratedAtLT applies the LT predicate on the "crated_at" field.
func CratedAtLT(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCratedAt), v))
	})
}

// CratedAtLTE applies the LTE predicate on the "crated_at" field.
func CratedAtLTE(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCratedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Answer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Answer(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Answer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Answer(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Question) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Answer) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Answer) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
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
func Not(p predicate.Answer) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		p(s.Not())
	})
}
