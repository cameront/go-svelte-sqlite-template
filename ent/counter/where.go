// Code generated by ent, DO NOT EDIT.

package counter

import (
	"entgo.io/ent/dialect/sql"
	"github.com/cameront/go-svelte-sqlite-template/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Counter {
	return predicate.Counter(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Counter {
	return predicate.Counter(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Counter {
	return predicate.Counter(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Counter {
	return predicate.Counter(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Counter {
	return predicate.Counter(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Counter {
	return predicate.Counter(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Counter {
	return predicate.Counter(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Counter {
	return predicate.Counter(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Counter {
	return predicate.Counter(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Counter {
	return predicate.Counter(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Counter {
	return predicate.Counter(sql.FieldContainsFold(FieldID, id))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int64) predicate.Counter {
	return predicate.Counter(sql.FieldEQ(FieldValue, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int64) predicate.Counter {
	return predicate.Counter(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int64) predicate.Counter {
	return predicate.Counter(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int64) predicate.Counter {
	return predicate.Counter(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int64) predicate.Counter {
	return predicate.Counter(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int64) predicate.Counter {
	return predicate.Counter(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int64) predicate.Counter {
	return predicate.Counter(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int64) predicate.Counter {
	return predicate.Counter(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int64) predicate.Counter {
	return predicate.Counter(sql.FieldLTE(FieldValue, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Counter) predicate.Counter {
	return predicate.Counter(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Counter) predicate.Counter {
	return predicate.Counter(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Counter) predicate.Counter {
	return predicate.Counter(sql.NotPredicates(p))
}