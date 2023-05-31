// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"shira-chan-dev/ent/order"
	"shira-chan-dev/ent/user"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// OrderEdge is the edge representation of Order.
type OrderEdge struct {
	Node   *Order `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// OrderConnection is the connection containing edges to Order.
type OrderConnection struct {
	Edges      []*OrderEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *OrderConnection) build(nodes []*Order, pager *orderPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Order
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Order {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Order {
			return nodes[i]
		}
	}
	c.Edges = make([]*OrderEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &OrderEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// OrderPaginateOption enables pagination customization.
type OrderPaginateOption func(*orderPager) error

// WithOrderOrder configures pagination ordering.
func WithOrderOrder(order *OrderOrder) OrderPaginateOption {
	if order == nil {
		order = DefaultOrderOrder
	}
	o := *order
	return func(pager *orderPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOrderOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOrderFilter configures pagination filter.
func WithOrderFilter(filter func(*OrderQuery) (*OrderQuery, error)) OrderPaginateOption {
	return func(pager *orderPager) error {
		if filter == nil {
			return errors.New("OrderQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type orderPager struct {
	reverse bool
	order   *OrderOrder
	filter  func(*OrderQuery) (*OrderQuery, error)
}

func newOrderPager(opts []OrderPaginateOption, reverse bool) (*orderPager, error) {
	pager := &orderPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOrderOrder
	}
	return pager, nil
}

func (p *orderPager) applyFilter(query *OrderQuery) (*OrderQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *orderPager) toCursor(o *Order) Cursor {
	return p.order.Field.toCursor(o)
}

func (p *orderPager) applyCursors(query *OrderQuery, after, before *Cursor) (*OrderQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultOrderOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *orderPager) applyOrder(query *OrderQuery) *OrderQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultOrderOrder.Field {
		query = query.Order(DefaultOrderOrder.Field.toTerm(direction.OrderTermOption()))
	}
	switch p.order.Field.column {
	case OrderOrderFieldReceiverCount.column:
	default:
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(p.order.Field.column)
		}
	}
	return query
}

func (p *orderPager) orderExpr(query *OrderQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	switch p.order.Field.column {
	case OrderOrderFieldReceiverCount.column:
		query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	default:
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(p.order.Field.column)
		}
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultOrderOrder.Field {
			b.Comma().Ident(DefaultOrderOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Order.
func (o *OrderQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OrderPaginateOption,
) (*OrderConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOrderPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if o, err = pager.applyFilter(o); err != nil {
		return nil, err
	}
	conn := &OrderConnection{Edges: []*OrderEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := o.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if o, err = pager.applyCursors(o, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		o.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := o.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	o = pager.applyOrder(o)
	nodes, err := o.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// OrderOrderFieldTitle orders Order by title.
	OrderOrderFieldTitle = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.Title, nil
		},
		column: order.FieldTitle,
		toTerm: order.ByTitle,
		toCursor: func(o *Order) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.Title,
			}
		},
	}
	// OrderOrderFieldType orders Order by type.
	OrderOrderFieldType = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.Type, nil
		},
		column: order.FieldType,
		toTerm: order.ByType,
		toCursor: func(o *Order) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.Type,
			}
		},
	}
	// OrderOrderFieldIsClosed orders Order by is_closed.
	OrderOrderFieldIsClosed = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.IsClosed, nil
		},
		column: order.FieldIsClosed,
		toTerm: order.ByIsClosed,
		toCursor: func(o *Order) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.IsClosed,
			}
		},
	}
	// OrderOrderFieldIsFinished orders Order by is_finished.
	OrderOrderFieldIsFinished = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.IsFinished, nil
		},
		column: order.FieldIsFinished,
		toTerm: order.ByIsFinished,
		toCursor: func(o *Order) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.IsFinished,
			}
		},
	}
	// OrderOrderFieldEvaluation orders Order by evaluation.
	OrderOrderFieldEvaluation = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.Evaluation, nil
		},
		column: order.FieldEvaluation,
		toTerm: order.ByEvaluation,
		toCursor: func(o *Order) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.Evaluation,
			}
		},
	}
	// OrderOrderFieldHopeAt orders Order by hope_at.
	OrderOrderFieldHopeAt = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.HopeAt, nil
		},
		column: order.FieldHopeAt,
		toTerm: order.ByHopeAt,
		toCursor: func(o *Order) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.HopeAt,
			}
		},
	}
	// OrderOrderFieldCreatedAt orders Order by created_at.
	OrderOrderFieldCreatedAt = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.CreatedAt, nil
		},
		column: order.FieldCreatedAt,
		toTerm: order.ByCreatedAt,
		toCursor: func(o *Order) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.CreatedAt,
			}
		},
	}
	// OrderOrderFieldUpdatedAt orders Order by updated_at.
	OrderOrderFieldUpdatedAt = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.UpdatedAt, nil
		},
		column: order.FieldUpdatedAt,
		toTerm: order.ByUpdatedAt,
		toCursor: func(o *Order) Cursor {
			return Cursor{
				ID:    o.ID,
				Value: o.UpdatedAt,
			}
		},
	}
	// OrderOrderFieldReceiverCount orders by RECEIVER_COUNT.
	OrderOrderFieldReceiverCount = &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.Value("receiver_count")
		},
		column: "receiver_count",
		toTerm: func(opts ...sql.OrderTermOption) order.OrderOption {
			return order.ByReceiverCount(
				append(opts, sql.OrderSelectAs("receiver_count"))...,
			)
		},
		toCursor: func(o *Order) Cursor {
			cv, _ := o.Value("receiver_count")
			return Cursor{
				ID:    o.ID,
				Value: cv,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f OrderOrderField) String() string {
	var str string
	switch f.column {
	case OrderOrderFieldTitle.column:
		str = "TITLE"
	case OrderOrderFieldType.column:
		str = "TYPE"
	case OrderOrderFieldIsClosed.column:
		str = "IS_CLOSED"
	case OrderOrderFieldIsFinished.column:
		str = "IS_FINISHED"
	case OrderOrderFieldEvaluation.column:
		str = "EVALUATION"
	case OrderOrderFieldHopeAt.column:
		str = "HOPE_AT"
	case OrderOrderFieldCreatedAt.column:
		str = "CREAT_AT"
	case OrderOrderFieldUpdatedAt.column:
		str = "UPDATED_AT"
	case OrderOrderFieldReceiverCount.column:
		str = "RECEIVER_COUNT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f OrderOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *OrderOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("OrderOrderField %T must be a string", v)
	}
	switch str {
	case "TITLE":
		*f = *OrderOrderFieldTitle
	case "TYPE":
		*f = *OrderOrderFieldType
	case "IS_CLOSED":
		*f = *OrderOrderFieldIsClosed
	case "IS_FINISHED":
		*f = *OrderOrderFieldIsFinished
	case "EVALUATION":
		*f = *OrderOrderFieldEvaluation
	case "HOPE_AT":
		*f = *OrderOrderFieldHopeAt
	case "CREAT_AT":
		*f = *OrderOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *OrderOrderFieldUpdatedAt
	case "RECEIVER_COUNT":
		*f = *OrderOrderFieldReceiverCount
	default:
		return fmt.Errorf("%s is not a valid OrderOrderField", str)
	}
	return nil
}

// OrderOrderField defines the ordering field of Order.
type OrderOrderField struct {
	// Value extracts the ordering value from the given Order.
	Value    func(*Order) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) order.OrderOption
	toCursor func(*Order) Cursor
}

// OrderOrder defines the ordering of Order.
type OrderOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *OrderOrderField `json:"field"`
}

// DefaultOrderOrder is the default ordering of Order.
var DefaultOrderOrder = &OrderOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &OrderOrderField{
		Value: func(o *Order) (ent.Value, error) {
			return o.ID, nil
		},
		column: order.FieldID,
		toTerm: order.ByID,
		toCursor: func(o *Order) Cursor {
			return Cursor{ID: o.ID}
		},
	},
}

// ToEdge converts Order into OrderEdge.
func (o *Order) ToEdge(order *OrderOrder) *OrderEdge {
	if order == nil {
		order = DefaultOrderOrder
	}
	return &OrderEdge{
		Node:   o,
		Cursor: order.Field.toCursor(o),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   *UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultUserOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	switch p.order.Field.column {
	case UserOrderFieldRequestedCount.column, UserOrderFieldReceivedCount.column:
	default:
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(p.order.Field.column)
		}
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	switch p.order.Field.column {
	case UserOrderFieldRequestedCount.column, UserOrderFieldReceivedCount.column:
		query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	default:
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(p.order.Field.column)
		}
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := u.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserOrderFieldUname orders User by uname.
	UserOrderFieldUname = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.Uname, nil
		},
		column: user.FieldUname,
		toTerm: user.ByUname,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Uname,
			}
		},
	}
	// UserOrderFieldIsAdmin orders User by is_admin.
	UserOrderFieldIsAdmin = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.IsAdmin, nil
		},
		column: user.FieldIsAdmin,
		toTerm: user.ByIsAdmin,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.IsAdmin,
			}
		},
	}
	// UserOrderFieldIsActive orders User by is_active.
	UserOrderFieldIsActive = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.IsActive, nil
		},
		column: user.FieldIsActive,
		toTerm: user.ByIsActive,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.IsActive,
			}
		},
	}
	// UserOrderFieldCreatedAt orders User by created_at.
	UserOrderFieldCreatedAt = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.CreatedAt, nil
		},
		column: user.FieldCreatedAt,
		toTerm: user.ByCreatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.CreatedAt,
			}
		},
	}
	// UserOrderFieldUpdatedAt orders User by updated_at.
	UserOrderFieldUpdatedAt = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.UpdatedAt, nil
		},
		column: user.FieldUpdatedAt,
		toTerm: user.ByUpdatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.UpdatedAt,
			}
		},
	}
	// UserOrderFieldRequestedCount orders by REQUESTED_COUNT.
	UserOrderFieldRequestedCount = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.Value("requested_count")
		},
		column: "requested_count",
		toTerm: func(opts ...sql.OrderTermOption) user.OrderOption {
			return user.ByRequestedCount(
				append(opts, sql.OrderSelectAs("requested_count"))...,
			)
		},
		toCursor: func(u *User) Cursor {
			cv, _ := u.Value("requested_count")
			return Cursor{
				ID:    u.ID,
				Value: cv,
			}
		},
	}
	// UserOrderFieldReceivedCount orders by RECEIVED_COUNT.
	UserOrderFieldReceivedCount = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.Value("received_count")
		},
		column: "received_count",
		toTerm: func(opts ...sql.OrderTermOption) user.OrderOption {
			return user.ByReceivedCount(
				append(opts, sql.OrderSelectAs("received_count"))...,
			)
		},
		toCursor: func(u *User) Cursor {
			cv, _ := u.Value("received_count")
			return Cursor{
				ID:    u.ID,
				Value: cv,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.column {
	case UserOrderFieldUname.column:
		str = "UNAME"
	case UserOrderFieldIsAdmin.column:
		str = "IS_ADMIN"
	case UserOrderFieldIsActive.column:
		str = "STATE"
	case UserOrderFieldCreatedAt.column:
		str = "CREAT_AT"
	case UserOrderFieldUpdatedAt.column:
		str = "UPDATED_AT"
	case UserOrderFieldRequestedCount.column:
		str = "REQUESTED_COUNT"
	case UserOrderFieldReceivedCount.column:
		str = "RECEIVED_COUNT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "UNAME":
		*f = *UserOrderFieldUname
	case "IS_ADMIN":
		*f = *UserOrderFieldIsAdmin
	case "STATE":
		*f = *UserOrderFieldIsActive
	case "CREAT_AT":
		*f = *UserOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *UserOrderFieldUpdatedAt
	case "REQUESTED_COUNT":
		*f = *UserOrderFieldRequestedCount
	case "RECEIVED_COUNT":
		*f = *UserOrderFieldReceivedCount
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
