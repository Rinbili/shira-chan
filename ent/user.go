// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"shira-chan-dev/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 用户名
	Uname string `json:"uname,omitempty"`
	// 密码
	Passwd string `json:"-"`
	// 手机号码
	Phone string `json:"phone,omitempty"`
	// 微信号
	Wechat string `json:"wechat,omitempty"`
	// 是否管理员
	IsAdmin bool `json:"is_admin,omitempty"`
	// 用户状态
	IsActive bool `json:"is_active,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// 需求
	Requested []*Order `json:"requested,omitempty"`
	// 接单
	Received []*Order `json:"received,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedRequested map[string][]*Order
	namedReceived  map[string][]*Order
}

// RequestedOrErr returns the Requested value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RequestedOrErr() ([]*Order, error) {
	if e.loadedTypes[0] {
		return e.Requested, nil
	}
	return nil, &NotLoadedError{edge: "requested"}
}

// ReceivedOrErr returns the Received value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReceivedOrErr() ([]*Order, error) {
	if e.loadedTypes[1] {
		return e.Received, nil
	}
	return nil, &NotLoadedError{edge: "received"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsAdmin, user.FieldIsActive:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUname, user.FieldPasswd, user.FieldPhone, user.FieldWechat:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uname", values[i])
			} else if value.Valid {
				u.Uname = value.String
			}
		case user.FieldPasswd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field passwd", values[i])
			} else if value.Valid {
				u.Passwd = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldWechat:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wechat", values[i])
			} else if value.Valid {
				u.Wechat = value.String
			}
		case user.FieldIsAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_admin", values[i])
			} else if value.Valid {
				u.IsAdmin = value.Bool
			}
		case user.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				u.IsActive = value.Bool
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryRequested queries the "requested" edge of the User entity.
func (u *User) QueryRequested() *OrderQuery {
	return NewUserClient(u.config).QueryRequested(u)
}

// QueryReceived queries the "received" edge of the User entity.
func (u *User) QueryReceived() *OrderQuery {
	return NewUserClient(u.config).QueryReceived(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("uname=")
	builder.WriteString(u.Uname)
	builder.WriteString(", ")
	builder.WriteString("passwd=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("wechat=")
	builder.WriteString(u.Wechat)
	builder.WriteString(", ")
	builder.WriteString("is_admin=")
	builder.WriteString(fmt.Sprintf("%v", u.IsAdmin))
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", u.IsActive))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedRequested returns the Requested named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedRequested(name string) ([]*Order, error) {
	if u.Edges.namedRequested == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedRequested[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedRequested(name string, edges ...*Order) {
	if u.Edges.namedRequested == nil {
		u.Edges.namedRequested = make(map[string][]*Order)
	}
	if len(edges) == 0 {
		u.Edges.namedRequested[name] = []*Order{}
	} else {
		u.Edges.namedRequested[name] = append(u.Edges.namedRequested[name], edges...)
	}
}

// NamedReceived returns the Received named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedReceived(name string) ([]*Order, error) {
	if u.Edges.namedReceived == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedReceived[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedReceived(name string, edges ...*Order) {
	if u.Edges.namedReceived == nil {
		u.Edges.namedReceived = make(map[string][]*Order)
	}
	if len(edges) == 0 {
		u.Edges.namedReceived[name] = []*Order{}
	} else {
		u.Edges.namedReceived[name] = append(u.Edges.namedReceived[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
