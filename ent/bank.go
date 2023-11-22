// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend-demo/ent/bank"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Bank is the model entity for the Bank schema.
type Bank struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BankQuery when eager-loading is set.
	Edges        BankEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BankEdges holds the relations/edges for other nodes in the graph.
type BankEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e BankEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bank) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bank.FieldID:
			values[i] = new(sql.NullInt64)
		case bank.FieldCode, bank.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bank fields.
func (b *Bank) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bank.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case bank.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				b.Code = value.String
			}
		case bank.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bank.
// This includes values selected through modifiers, order, etc.
func (b *Bank) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Bank entity.
func (b *Bank) QueryUsers() *UserQuery {
	return NewBankClient(b.config).QueryUsers(b)
}

// Update returns a builder for updating this Bank.
// Note that you need to call Bank.Unwrap() before calling this method if this Bank
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bank) Update() *BankUpdateOne {
	return NewBankClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bank entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bank) Unwrap() *Bank {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bank is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bank) String() string {
	var builder strings.Builder
	builder.WriteString("Bank(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("code=")
	builder.WriteString(b.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Banks is a parsable slice of Bank.
type Banks []*Bank
