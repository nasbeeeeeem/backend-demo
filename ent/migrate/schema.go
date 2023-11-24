// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BanksColumns holds the columns for the "banks" table.
	BanksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 4},
		{Name: "name", Type: field.TypeString},
	}
	// BanksTable holds the schema information for the "banks" table.
	BanksTable = &schema.Table{
		Name:       "banks",
		Columns:    BanksColumns,
		PrimaryKey: []*schema.Column{BanksColumns[0]},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_users_events",
				Columns:    []*schema.Column{EventsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "photo_url", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "account_code", Type: field.TypeString, Nullable: true},
		{Name: "branch_code", Type: field.TypeString, Nullable: true},
		{Name: "bank_code", Type: field.TypeString, Size: 4},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_banks_users",
				Columns:    []*schema.Column{UsersColumns[9]},
				RefColumns: []*schema.Column{BanksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BanksTable,
		EventsTable,
		PaymentsTable,
		UsersTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = BanksTable
}
