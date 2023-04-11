// Code generated by ent, DO NOT EDIT.

package masteraccountcategory

import (
	"outgrow/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldLTE(FieldID, id))
}

// AccountTypeID applies equality check predicate on the "account_type_id" field. It's identical to AccountTypeIDEQ.
func AccountTypeID(v int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldAccountTypeID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldDescription, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// AccountTypeIDEQ applies the EQ predicate on the "account_type_id" field.
func AccountTypeIDEQ(v int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldAccountTypeID, v))
}

// AccountTypeIDNEQ applies the NEQ predicate on the "account_type_id" field.
func AccountTypeIDNEQ(v int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNEQ(FieldAccountTypeID, v))
}

// AccountTypeIDIn applies the In predicate on the "account_type_id" field.
func AccountTypeIDIn(vs ...int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldIn(FieldAccountTypeID, vs...))
}

// AccountTypeIDNotIn applies the NotIn predicate on the "account_type_id" field.
func AccountTypeIDNotIn(vs ...int) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNotIn(FieldAccountTypeID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldContainsFold(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(sql.FieldLTE(FieldCreatedAt, v))
}

// HasAccounts applies the HasEdge predicate on the "accounts" edge.
func HasAccounts() predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountsTable, AccountsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountsWith applies the HasEdge predicate on the "accounts" edge with a given conditions (other predicates).
func HasAccountsWith(preds ...predicate.MasterAccount) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountsTable, AccountsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccType applies the HasEdge predicate on the "accType" edge.
func HasAccType() predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccTypeTable, AccTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccTypeWith applies the HasEdge predicate on the "accType" edge with a given conditions (other predicates).
func HasAccTypeWith(preds ...predicate.MasterAccountType) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccTypeTable, AccTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MasterAccountCategory) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MasterAccountCategory) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(func(s *sql.Selector) {
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
func Not(p predicate.MasterAccountCategory) predicate.MasterAccountCategory {
	return predicate.MasterAccountCategory(func(s *sql.Selector) {
		p(s.Not())
	})
}
