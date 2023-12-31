// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend-demo/ent/payment"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
}

// SetEventID sets the "event_id" field.
func (pc *PaymentCreate) SetEventID(i int) *PaymentCreate {
	pc.mutation.SetEventID(i)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PaymentCreate) SetAmount(i int) *PaymentCreate {
	pc.mutation.SetAmount(i)
	return pc
}

// SetPaidBy sets the "paid_by" field.
func (pc *PaymentCreate) SetPaidBy(i int) *PaymentCreate {
	pc.mutation.SetPaidBy(i)
	return pc
}

// SetPaidTo sets the "paid_to" field.
func (pc *PaymentCreate) SetPaidTo(i int) *PaymentCreate {
	pc.mutation.SetPaidTo(i)
	return pc
}

// SetPaidAt sets the "paid_at" field.
func (pc *PaymentCreate) SetPaidAt(t time.Time) *PaymentCreate {
	pc.mutation.SetPaidAt(t)
	return pc
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillablePaidAt(t *time.Time) *PaymentCreate {
	if t != nil {
		pc.SetPaidAt(*t)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PaymentCreate) SetCreatedAt(t time.Time) *PaymentCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableCreatedAt(t *time.Time) *PaymentCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PaymentCreate) SetUpdatedAt(t time.Time) *PaymentCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableUpdatedAt(t *time.Time) *PaymentCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PaymentCreate) SetDeletedAt(t time.Time) *PaymentCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableDeletedAt(t *time.Time) *PaymentCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PaymentCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PaymentCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PaymentCreate) defaults() {
	if _, ok := pc.mutation.PaidAt(); !ok {
		v := payment.DefaultPaidAt()
		pc.mutation.SetPaidAt(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := payment.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := payment.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PaymentCreate) check() error {
	if _, ok := pc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event_id", err: errors.New(`ent: missing required field "Payment.event_id"`)}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Payment.amount"`)}
	}
	if _, ok := pc.mutation.PaidBy(); !ok {
		return &ValidationError{Name: "paid_by", err: errors.New(`ent: missing required field "Payment.paid_by"`)}
	}
	if _, ok := pc.mutation.PaidTo(); !ok {
		return &ValidationError{Name: "paid_to", err: errors.New(`ent: missing required field "Payment.paid_to"`)}
	}
	if _, ok := pc.mutation.PaidAt(); !ok {
		return &ValidationError{Name: "paid_at", err: errors.New(`ent: missing required field "Payment.paid_at"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Payment.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Payment.updated_at"`)}
	}
	return nil
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		_node = &Payment{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(payment.Table, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.EventID(); ok {
		_spec.SetField(payment.FieldEventID, field.TypeInt, value)
		_node.EventID = &value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.PaidBy(); ok {
		_spec.SetField(payment.FieldPaidBy, field.TypeInt, value)
		_node.PaidBy = value
	}
	if value, ok := pc.mutation.PaidTo(); ok {
		_spec.SetField(payment.FieldPaidTo, field.TypeInt, value)
		_node.PaidTo = value
	}
	if value, ok := pc.mutation.PaidAt(); ok {
		_spec.SetField(payment.FieldPaidAt, field.TypeTime, value)
		_node.PaidAt = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(payment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(payment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(payment.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// PaymentCreateBulk is the builder for creating many Payment entities in bulk.
type PaymentCreateBulk struct {
	config
	err      error
	builders []*PaymentCreate
}

// Save creates the Payment entities in the database.
func (pcb *PaymentCreateBulk) Save(ctx context.Context) ([]*Payment, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Payment, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PaymentCreateBulk) SaveX(ctx context.Context) []*Payment {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PaymentCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
