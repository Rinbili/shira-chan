// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUname holds the string denoting the uname field in the database.
	FieldUname = "uname"
	// FieldPasswd holds the string denoting the passwd field in the database.
	FieldPasswd = "passwd"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldIsAdmin holds the string denoting the is_admin field in the database.
	FieldIsAdmin = "is_admin"
	// FieldIsSecretary holds the string denoting the is_secretary field in the database.
	FieldIsSecretary = "is_secretary"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// EdgeRequested holds the string denoting the requested edge name in mutations.
	EdgeRequested = "requested"
	// EdgeReceived holds the string denoting the received edge name in mutations.
	EdgeReceived = "received"
	// Table holds the table name of the user in the database.
	Table = "Users"
	// RequestedTable is the table that holds the requested relation/edge.
	RequestedTable = "Orders"
	// RequestedInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	RequestedInverseTable = "Orders"
	// RequestedColumn is the table column denoting the requested relation/edge.
	RequestedColumn = "user_requested"
	// ReceivedTable is the table that holds the received relation/edge. The primary key declared below.
	ReceivedTable = "order_receiver"
	// ReceivedInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	ReceivedInverseTable = "Orders"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUname,
	FieldPasswd,
	FieldPhone,
	FieldIsAdmin,
	FieldIsSecretary,
	FieldIsActive,
}

var (
	// ReceivedPrimaryKey and ReceivedColumn2 are the table columns denoting the
	// primary key for the received relation (M2M).
	ReceivedPrimaryKey = []string{"order_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "shira-chan-dev/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() int64
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() int64
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int64
	// UnameValidator is a validator for the "uname" field. It is called by the builders before save.
	UnameValidator func(string) error
	// PasswdValidator is a validator for the "passwd" field. It is called by the builders before save.
	PasswdValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultIsAdmin holds the default value on creation for the "is_admin" field.
	DefaultIsAdmin bool
	// DefaultIsSecretary holds the default value on creation for the "is_secretary" field.
	DefaultIsSecretary bool
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUname orders the results by the uname field.
func ByUname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUname, opts...).ToFunc()
}

// ByPasswd orders the results by the passwd field.
func ByPasswd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswd, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByIsAdmin orders the results by the is_admin field.
func ByIsAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAdmin, opts...).ToFunc()
}

// ByIsSecretary orders the results by the is_secretary field.
func ByIsSecretary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsSecretary, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByRequestedCount orders the results by requested count.
func ByRequestedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRequestedStep(), opts...)
	}
}

// ByRequested orders the results by requested terms.
func ByRequested(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRequestedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReceivedCount orders the results by received count.
func ByReceivedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReceivedStep(), opts...)
	}
}

// ByReceived orders the results by received terms.
func ByReceived(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceivedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRequestedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RequestedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RequestedTable, RequestedColumn),
	)
}
func newReceivedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceivedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ReceivedTable, ReceivedPrimaryKey...),
	)
}
