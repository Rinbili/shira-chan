// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "Orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Size: 100},
		{Name: "content", Type: field.TypeString, Size: 15000},
		{Name: "contact", Type: field.TypeString, Size: 20},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"software", "hardware", "unknown", "other"}, Default: "other"},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"requested", "received", "finished", "closed"}, Default: "requested"},
		{Name: "evaluation", Type: field.TypeFloat64, Nullable: true},
		{Name: "hope_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "user_requested", Type: field.TypeInt, Nullable: true},
	}
	// OrdersTable holds the schema information for the "Orders" table.
	OrdersTable = &schema.Table{
		Name:       "Orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Orders_Users_requested",
				Columns:    []*schema.Column{OrdersColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "Users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uname", Type: field.TypeString, Size: 30},
		{Name: "passwd", Type: field.TypeString, Size: 20},
		{Name: "phone", Type: field.TypeString, Unique: true, Size: 15},
		{Name: "wechat", Type: field.TypeString, Size: 30, Default: ""},
		{Name: "level", Type: field.TypeEnum, Enums: []string{"root", "admin", "president", "minister", "member", "user", "banned"}, Default: "user"},
		{Name: "dept", Type: field.TypeEnum, Enums: []string{"none"}, Default: "none"},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"on", "off"}, Default: "on"},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UsersTable holds the schema information for the "Users" table.
	UsersTable = &schema.Table{
		Name:       "Users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// OrderReceiverColumns holds the columns for the "order_receiver" table.
	OrderReceiverColumns = []*schema.Column{
		{Name: "order_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// OrderReceiverTable holds the schema information for the "order_receiver" table.
	OrderReceiverTable = &schema.Table{
		Name:       "order_receiver",
		Columns:    OrderReceiverColumns,
		PrimaryKey: []*schema.Column{OrderReceiverColumns[0], OrderReceiverColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_receiver_order_id",
				Columns:    []*schema.Column{OrderReceiverColumns[0]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "order_receiver_user_id",
				Columns:    []*schema.Column{OrderReceiverColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
		UsersTable,
		OrderReceiverTable,
	}
)

func init() {
	OrdersTable.ForeignKeys[0].RefTable = UsersTable
	OrdersTable.Annotation = &entsql.Annotation{
		Table: "Orders",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Table: "Users",
	}
	OrderReceiverTable.ForeignKeys[0].RefTable = OrdersTable
	OrderReceiverTable.ForeignKeys[1].RefTable = UsersTable
}
