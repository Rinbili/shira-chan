// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"shira-chan-dev/ent/order"
	"shira-chan-dev/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 标题
	Title string `json:"title,omitempty"`
	// 内容
	Content string `json:"content,omitempty"`
	// 联系方式
	Contact string `json:"contact,omitempty"`
	// 故障类别
	Type string `json:"type,omitempty"`
	// 是否被关闭
	IsClosed bool `json:"is_closed,omitempty"`
	// 是否完成
	IsFinished bool `json:"is_finished,omitempty"`
	// 评分
	Evaluation *float64 `json:"evaluation,omitempty"`
	// 期望时间
	HopeAt time.Time `json:"hope_at,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges          OrderEdges `json:"edges"`
	user_requested *int
	selectValues   sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// 需求者
	Requester *User `json:"requester,omitempty"`
	// 接单者
	Receiver []*User `json:"receiver,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedReceiver map[string][]*User
}

// RequesterOrErr returns the Requester value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) RequesterOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Requester == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Requester, nil
	}
	return nil, &NotLoadedError{edge: "requester"}
}

// ReceiverOrErr returns the Receiver value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) ReceiverOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Receiver, nil
	}
	return nil, &NotLoadedError{edge: "receiver"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldIsClosed, order.FieldIsFinished:
			values[i] = new(sql.NullBool)
		case order.FieldEvaluation:
			values[i] = new(sql.NullFloat64)
		case order.FieldID:
			values[i] = new(sql.NullInt64)
		case order.FieldTitle, order.FieldContent, order.FieldContact, order.FieldType:
			values[i] = new(sql.NullString)
		case order.FieldHopeAt, order.FieldCreatedAt, order.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case order.ForeignKeys[0]: // user_requested
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				o.Title = value.String
			}
		case order.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				o.Content = value.String
			}
		case order.FieldContact:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact", values[i])
			} else if value.Valid {
				o.Contact = value.String
			}
		case order.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				o.Type = value.String
			}
		case order.FieldIsClosed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_closed", values[i])
			} else if value.Valid {
				o.IsClosed = value.Bool
			}
		case order.FieldIsFinished:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_finished", values[i])
			} else if value.Valid {
				o.IsFinished = value.Bool
			}
		case order.FieldEvaluation:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field evaluation", values[i])
			} else if value.Valid {
				o.Evaluation = new(float64)
				*o.Evaluation = value.Float64
			}
		case order.FieldHopeAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field hope_at", values[i])
			} else if value.Valid {
				o.HopeAt = value.Time
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_requested", value)
			} else if value.Valid {
				o.user_requested = new(int)
				*o.user_requested = int(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryRequester queries the "requester" edge of the Order entity.
func (o *Order) QueryRequester() *UserQuery {
	return NewOrderClient(o.config).QueryRequester(o)
}

// QueryReceiver queries the "receiver" edge of the Order entity.
func (o *Order) QueryReceiver() *UserQuery {
	return NewOrderClient(o.config).QueryReceiver(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("title=")
	builder.WriteString(o.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(o.Content)
	builder.WriteString(", ")
	builder.WriteString("contact=")
	builder.WriteString(o.Contact)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(o.Type)
	builder.WriteString(", ")
	builder.WriteString("is_closed=")
	builder.WriteString(fmt.Sprintf("%v", o.IsClosed))
	builder.WriteString(", ")
	builder.WriteString("is_finished=")
	builder.WriteString(fmt.Sprintf("%v", o.IsFinished))
	builder.WriteString(", ")
	if v := o.Evaluation; v != nil {
		builder.WriteString("evaluation=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("hope_at=")
	builder.WriteString(o.HopeAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedReceiver returns the Receiver named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Order) NamedReceiver(name string) ([]*User, error) {
	if o.Edges.namedReceiver == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedReceiver[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Order) appendNamedReceiver(name string, edges ...*User) {
	if o.Edges.namedReceiver == nil {
		o.Edges.namedReceiver = make(map[string][]*User)
	}
	if len(edges) == 0 {
		o.Edges.namedReceiver[name] = []*User{}
	} else {
		o.Edges.namedReceiver[name] = append(o.Edges.namedReceiver[name], edges...)
	}
}

// Orders is a parsable slice of Order.
type Orders []*Order
