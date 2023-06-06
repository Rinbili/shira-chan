// Code generated by ent, DO NOT EDIT.

package user

import (
	"shira-chan-dev/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Uname applies equality check predicate on the "uname" field. It's identical to UnameEQ.
func Uname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUname, v))
}

// Passwd applies equality check predicate on the "passwd" field. It's identical to PasswdEQ.
func Passwd(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswd, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// IsAdmin applies equality check predicate on the "is_admin" field. It's identical to IsAdminEQ.
func IsAdmin(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsAdmin, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsActive, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UnameEQ applies the EQ predicate on the "uname" field.
func UnameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUname, v))
}

// UnameNEQ applies the NEQ predicate on the "uname" field.
func UnameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUname, v))
}

// UnameIn applies the In predicate on the "uname" field.
func UnameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUname, vs...))
}

// UnameNotIn applies the NotIn predicate on the "uname" field.
func UnameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUname, vs...))
}

// UnameGT applies the GT predicate on the "uname" field.
func UnameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUname, v))
}

// UnameGTE applies the GTE predicate on the "uname" field.
func UnameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUname, v))
}

// UnameLT applies the LT predicate on the "uname" field.
func UnameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUname, v))
}

// UnameLTE applies the LTE predicate on the "uname" field.
func UnameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUname, v))
}

// UnameContains applies the Contains predicate on the "uname" field.
func UnameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUname, v))
}

// UnameHasPrefix applies the HasPrefix predicate on the "uname" field.
func UnameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUname, v))
}

// UnameHasSuffix applies the HasSuffix predicate on the "uname" field.
func UnameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUname, v))
}

// UnameEqualFold applies the EqualFold predicate on the "uname" field.
func UnameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUname, v))
}

// UnameContainsFold applies the ContainsFold predicate on the "uname" field.
func UnameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUname, v))
}

// PasswdEQ applies the EQ predicate on the "passwd" field.
func PasswdEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswd, v))
}

// PasswdNEQ applies the NEQ predicate on the "passwd" field.
func PasswdNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPasswd, v))
}

// PasswdIn applies the In predicate on the "passwd" field.
func PasswdIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPasswd, vs...))
}

// PasswdNotIn applies the NotIn predicate on the "passwd" field.
func PasswdNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPasswd, vs...))
}

// PasswdGT applies the GT predicate on the "passwd" field.
func PasswdGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPasswd, v))
}

// PasswdGTE applies the GTE predicate on the "passwd" field.
func PasswdGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPasswd, v))
}

// PasswdLT applies the LT predicate on the "passwd" field.
func PasswdLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPasswd, v))
}

// PasswdLTE applies the LTE predicate on the "passwd" field.
func PasswdLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPasswd, v))
}

// PasswdContains applies the Contains predicate on the "passwd" field.
func PasswdContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPasswd, v))
}

// PasswdHasPrefix applies the HasPrefix predicate on the "passwd" field.
func PasswdHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPasswd, v))
}

// PasswdHasSuffix applies the HasSuffix predicate on the "passwd" field.
func PasswdHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPasswd, v))
}

// PasswdEqualFold applies the EqualFold predicate on the "passwd" field.
func PasswdEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPasswd, v))
}

// PasswdContainsFold applies the ContainsFold predicate on the "passwd" field.
func PasswdContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPasswd, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhone, v))
}

// IsAdminEQ applies the EQ predicate on the "is_admin" field.
func IsAdminEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsAdmin, v))
}

// IsAdminNEQ applies the NEQ predicate on the "is_admin" field.
func IsAdminNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsAdmin, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsActive, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasRequested applies the HasEdge predicate on the "requested" edge.
func HasRequested() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RequestedTable, RequestedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRequestedWith applies the HasEdge predicate on the "requested" edge with a given conditions (other predicates).
func HasRequestedWith(preds ...predicate.Order) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newRequestedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReceived applies the HasEdge predicate on the "received" edge.
func HasReceived() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ReceivedTable, ReceivedPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceivedWith applies the HasEdge predicate on the "received" edge with a given conditions (other predicates).
func HasReceivedWith(preds ...predicate.Order) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newReceivedStep()
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
