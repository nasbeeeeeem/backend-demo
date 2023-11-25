// Code generated by ent, DO NOT EDIT.

package payment

import (
	"backend-demo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldID, id))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldEventID, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldAmount, v))
}

// PaidBy applies equality check predicate on the "paid_by" field. It's identical to PaidByEQ.
func PaidBy(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaidBy, v))
}

// PaidTo applies equality check predicate on the "paid_to" field. It's identical to PaidToEQ.
func PaidTo(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaidTo, v))
}

// PaidAt applies equality check predicate on the "paid_at" field. It's identical to PaidAtEQ.
func PaidAt(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaidAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldDeletedAt, v))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldEventID, vs...))
}

// EventIDGT applies the GT predicate on the "event_id" field.
func EventIDGT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldEventID, v))
}

// EventIDGTE applies the GTE predicate on the "event_id" field.
func EventIDGTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldEventID, v))
}

// EventIDLT applies the LT predicate on the "event_id" field.
func EventIDLT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldEventID, v))
}

// EventIDLTE applies the LTE predicate on the "event_id" field.
func EventIDLTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldEventID, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldAmount, v))
}

// PaidByEQ applies the EQ predicate on the "paid_by" field.
func PaidByEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaidBy, v))
}

// PaidByNEQ applies the NEQ predicate on the "paid_by" field.
func PaidByNEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldPaidBy, v))
}

// PaidByIn applies the In predicate on the "paid_by" field.
func PaidByIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldPaidBy, vs...))
}

// PaidByNotIn applies the NotIn predicate on the "paid_by" field.
func PaidByNotIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldPaidBy, vs...))
}

// PaidByGT applies the GT predicate on the "paid_by" field.
func PaidByGT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldPaidBy, v))
}

// PaidByGTE applies the GTE predicate on the "paid_by" field.
func PaidByGTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldPaidBy, v))
}

// PaidByLT applies the LT predicate on the "paid_by" field.
func PaidByLT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldPaidBy, v))
}

// PaidByLTE applies the LTE predicate on the "paid_by" field.
func PaidByLTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldPaidBy, v))
}

// PaidToEQ applies the EQ predicate on the "paid_to" field.
func PaidToEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaidTo, v))
}

// PaidToNEQ applies the NEQ predicate on the "paid_to" field.
func PaidToNEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldPaidTo, v))
}

// PaidToIn applies the In predicate on the "paid_to" field.
func PaidToIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldPaidTo, vs...))
}

// PaidToNotIn applies the NotIn predicate on the "paid_to" field.
func PaidToNotIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldPaidTo, vs...))
}

// PaidToGT applies the GT predicate on the "paid_to" field.
func PaidToGT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldPaidTo, v))
}

// PaidToGTE applies the GTE predicate on the "paid_to" field.
func PaidToGTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldPaidTo, v))
}

// PaidToLT applies the LT predicate on the "paid_to" field.
func PaidToLT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldPaidTo, v))
}

// PaidToLTE applies the LTE predicate on the "paid_to" field.
func PaidToLTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldPaidTo, v))
}

// PaidAtEQ applies the EQ predicate on the "paid_at" field.
func PaidAtEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaidAt, v))
}

// PaidAtNEQ applies the NEQ predicate on the "paid_at" field.
func PaidAtNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldPaidAt, v))
}

// PaidAtIn applies the In predicate on the "paid_at" field.
func PaidAtIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldPaidAt, vs...))
}

// PaidAtNotIn applies the NotIn predicate on the "paid_at" field.
func PaidAtNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldPaidAt, vs...))
}

// PaidAtGT applies the GT predicate on the "paid_at" field.
func PaidAtGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldPaidAt, v))
}

// PaidAtGTE applies the GTE predicate on the "paid_at" field.
func PaidAtGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldPaidAt, v))
}

// PaidAtLT applies the LT predicate on the "paid_at" field.
func PaidAtLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldPaidAt, v))
}

// PaidAtLTE applies the LTE predicate on the "paid_at" field.
func PaidAtLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldPaidAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Payment {
	return predicate.Payment(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Payment {
	return predicate.Payment(sql.FieldNotNull(FieldDeletedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.NotPredicates(p))
}
