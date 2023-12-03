// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend-demo/ent/bank"
	"backend-demo/ent/event"
	"backend-demo/ent/group"
	"backend-demo/ent/payment"
	"backend-demo/ent/schema"
	"backend-demo/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bankFields := schema.Bank{}.Fields()
	_ = bankFields
	// bankDescName is the schema descriptor for name field.
	bankDescName := bankFields[1].Descriptor()
	// bank.NameValidator is a validator for the "name" field. It is called by the builders before save.
	bank.NameValidator = bankDescName.Validators[0].(func(string) error)
	// bankDescID is the schema descriptor for id field.
	bankDescID := bankFields[0].Descriptor()
	// bank.IDValidator is a validator for the "id" field. It is called by the builders before save.
	bank.IDValidator = func() func(string) error {
		validators := bankDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescName is the schema descriptor for name field.
	eventDescName := eventFields[0].Descriptor()
	// event.NameValidator is a validator for the "name" field. It is called by the builders before save.
	event.NameValidator = eventDescName.Validators[0].(func(string) error)
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventFields[2].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescUpdatedAt is the schema descriptor for updated_at field.
	eventDescUpdatedAt := eventFields[3].Descriptor()
	// event.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	event.DefaultUpdatedAt = eventDescUpdatedAt.Default.(func() time.Time)
	// event.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	event.UpdateDefaultUpdatedAt = eventDescUpdatedAt.UpdateDefault.(func() time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescPaidAt is the schema descriptor for paid_at field.
	paymentDescPaidAt := paymentFields[4].Descriptor()
	// payment.DefaultPaidAt holds the default value on creation for the paid_at field.
	payment.DefaultPaidAt = paymentDescPaidAt.Default.(func() time.Time)
	// paymentDescCreatedAt is the schema descriptor for created_at field.
	paymentDescCreatedAt := paymentFields[5].Descriptor()
	// payment.DefaultCreatedAt holds the default value on creation for the created_at field.
	payment.DefaultCreatedAt = paymentDescCreatedAt.Default.(func() time.Time)
	// paymentDescUpdatedAt is the schema descriptor for updated_at field.
	paymentDescUpdatedAt := paymentFields[6].Descriptor()
	// payment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	payment.DefaultUpdatedAt = paymentDescUpdatedAt.Default.(func() time.Time)
	// payment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	payment.UpdateDefaultUpdatedAt = paymentDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
