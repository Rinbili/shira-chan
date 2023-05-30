// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"shira-chan-dev/ent/order"
	"shira-chan-dev/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (oc *OrderCreate) SetTitle(s string) *OrderCreate {
	oc.mutation.SetTitle(s)
	return oc
}

// SetContent sets the "content" field.
func (oc *OrderCreate) SetContent(s string) *OrderCreate {
	oc.mutation.SetContent(s)
	return oc
}

// SetContact sets the "contact" field.
func (oc *OrderCreate) SetContact(s string) *OrderCreate {
	oc.mutation.SetContact(s)
	return oc
}

// SetType sets the "type" field.
func (oc *OrderCreate) SetType(o order.Type) *OrderCreate {
	oc.mutation.SetType(o)
	return oc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (oc *OrderCreate) SetNillableType(o *order.Type) *OrderCreate {
	if o != nil {
		oc.SetType(*o)
	}
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(o order.Status) *OrderCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrderCreate) SetNillableStatus(o *order.Status) *OrderCreate {
	if o != nil {
		oc.SetStatus(*o)
	}
	return oc
}

// SetEvaluation sets the "evaluation" field.
func (oc *OrderCreate) SetEvaluation(f float64) *OrderCreate {
	oc.mutation.SetEvaluation(f)
	return oc
}

// SetNillableEvaluation sets the "evaluation" field if the given value is not nil.
func (oc *OrderCreate) SetNillableEvaluation(f *float64) *OrderCreate {
	if f != nil {
		oc.SetEvaluation(*f)
	}
	return oc
}

// SetHopeAt sets the "hope_at" field.
func (oc *OrderCreate) SetHopeAt(t time.Time) *OrderCreate {
	oc.mutation.SetHopeAt(t)
	return oc
}

// SetNillableHopeAt sets the "hope_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableHopeAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetHopeAt(*t)
	}
	return oc
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetRequesterID sets the "requester" edge to the User entity by ID.
func (oc *OrderCreate) SetRequesterID(id int) *OrderCreate {
	oc.mutation.SetRequesterID(id)
	return oc
}

// SetNillableRequesterID sets the "requester" edge to the User entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableRequesterID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetRequesterID(*id)
	}
	return oc
}

// SetRequester sets the "requester" edge to the User entity.
func (oc *OrderCreate) SetRequester(u *User) *OrderCreate {
	return oc.SetRequesterID(u.ID)
}

// AddReceifeIDs adds the "receives" edge to the User entity by IDs.
func (oc *OrderCreate) AddReceifeIDs(ids ...int) *OrderCreate {
	oc.mutation.AddReceifeIDs(ids...)
	return oc
}

// AddReceives adds the "receives" edges to the User entity.
func (oc *OrderCreate) AddReceives(u ...*User) *OrderCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return oc.AddReceifeIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.GetType(); !ok {
		v := order.DefaultType
		oc.mutation.SetType(v)
	}
	if _, ok := oc.mutation.Status(); !ok {
		v := order.DefaultStatus
		oc.mutation.SetStatus(v)
	}
	if _, ok := oc.mutation.HopeAt(); !ok {
		v := order.DefaultHopeAt
		oc.mutation.SetHopeAt(v)
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := order.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := order.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Order.title"`)}
	}
	if v, ok := oc.mutation.Title(); ok {
		if err := order.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Order.title": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Order.content"`)}
	}
	if v, ok := oc.mutation.Content(); ok {
		if err := order.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Order.content": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Contact(); !ok {
		return &ValidationError{Name: "contact", err: errors.New(`ent: missing required field "Order.contact"`)}
	}
	if v, ok := oc.mutation.Contact(); ok {
		if err := order.ContactValidator(v); err != nil {
			return &ValidationError{Name: "contact", err: fmt.Errorf(`ent: validator failed for field "Order.contact": %w`, err)}
		}
	}
	if _, ok := oc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Order.type"`)}
	}
	if v, ok := oc.mutation.GetType(); ok {
		if err := order.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Order.type": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	if _, ok := oc.mutation.HopeAt(); !ok {
		return &ValidationError{Name: "hope_at", err: errors.New(`ent: missing required field "Order.hope_at"`)}
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Order.updated_at"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.Title(); ok {
		_spec.SetField(order.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := oc.mutation.Content(); ok {
		_spec.SetField(order.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := oc.mutation.Contact(); ok {
		_spec.SetField(order.FieldContact, field.TypeString, value)
		_node.Contact = value
	}
	if value, ok := oc.mutation.GetType(); ok {
		_spec.SetField(order.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.Evaluation(); ok {
		_spec.SetField(order.FieldEvaluation, field.TypeFloat64, value)
		_node.Evaluation = &value
	}
	if value, ok := oc.mutation.HopeAt(); ok {
		_spec.SetField(order.FieldHopeAt, field.TypeTime, value)
		_node.HopeAt = value
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := oc.mutation.RequesterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.RequesterTable,
			Columns: []string{order.RequesterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_requests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ReceivesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.ReceivesTable,
			Columns: order.ReceivesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
