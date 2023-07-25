// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"shira-chan-dev/ent/order"
	"shira-chan-dev/ent/predicate"
	"shira-chan-dev/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(i int64) *UserUpdate {
	uu.mutation.ResetUpdatedAt()
	uu.mutation.SetUpdatedAt(i)
	return uu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (uu *UserUpdate) AddUpdatedAt(i int64) *UserUpdate {
	uu.mutation.AddUpdatedAt(i)
	return uu
}

// SetUname sets the "uname" field.
func (uu *UserUpdate) SetUname(s string) *UserUpdate {
	uu.mutation.SetUname(s)
	return uu
}

// SetPasswd sets the "passwd" field.
func (uu *UserUpdate) SetPasswd(s string) *UserUpdate {
	uu.mutation.SetPasswd(s)
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetIsAdmin sets the "is_admin" field.
func (uu *UserUpdate) SetIsAdmin(b bool) *UserUpdate {
	uu.mutation.SetIsAdmin(b)
	return uu
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsAdmin(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsAdmin(*b)
	}
	return uu
}

// SetIsSecretary sets the "is_secretary" field.
func (uu *UserUpdate) SetIsSecretary(b bool) *UserUpdate {
	uu.mutation.SetIsSecretary(b)
	return uu
}

// SetNillableIsSecretary sets the "is_secretary" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsSecretary(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsSecretary(*b)
	}
	return uu
}

// SetIsActive sets the "is_active" field.
func (uu *UserUpdate) SetIsActive(b bool) *UserUpdate {
	uu.mutation.SetIsActive(b)
	return uu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsActive(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsActive(*b)
	}
	return uu
}

// AddRequestedIDs adds the "requested" edge to the Order entity by IDs.
func (uu *UserUpdate) AddRequestedIDs(ids ...int) *UserUpdate {
	uu.mutation.AddRequestedIDs(ids...)
	return uu
}

// AddRequested adds the "requested" edges to the Order entity.
func (uu *UserUpdate) AddRequested(o ...*Order) *UserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.AddRequestedIDs(ids...)
}

// AddReceivedIDs adds the "received" edge to the Order entity by IDs.
func (uu *UserUpdate) AddReceivedIDs(ids ...int) *UserUpdate {
	uu.mutation.AddReceivedIDs(ids...)
	return uu
}

// AddReceived adds the "received" edges to the Order entity.
func (uu *UserUpdate) AddReceived(o ...*Order) *UserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.AddReceivedIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearRequested clears all "requested" edges to the Order entity.
func (uu *UserUpdate) ClearRequested() *UserUpdate {
	uu.mutation.ClearRequested()
	return uu
}

// RemoveRequestedIDs removes the "requested" edge to Order entities by IDs.
func (uu *UserUpdate) RemoveRequestedIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveRequestedIDs(ids...)
	return uu
}

// RemoveRequested removes "requested" edges to Order entities.
func (uu *UserUpdate) RemoveRequested(o ...*Order) *UserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.RemoveRequestedIDs(ids...)
}

// ClearReceived clears all "received" edges to the Order entity.
func (uu *UserUpdate) ClearReceived() *UserUpdate {
	uu.mutation.ClearReceived()
	return uu
}

// RemoveReceivedIDs removes the "received" edge to Order entities by IDs.
func (uu *UserUpdate) RemoveReceivedIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveReceivedIDs(ids...)
	return uu
}

// RemoveReceived removes "received" edges to Order entities.
func (uu *UserUpdate) RemoveReceived(o ...*Order) *UserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.RemoveReceivedIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Uname(); ok {
		if err := user.UnameValidator(v); err != nil {
			return &ValidationError{Name: "uname", err: fmt.Errorf(`ent: validator failed for field "User.uname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Passwd(); ok {
		if err := user.PasswdValidator(v); err != nil {
			return &ValidationError{Name: "passwd", err: fmt.Errorf(`ent: validator failed for field "User.passwd": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.Uname(); ok {
		_spec.SetField(user.FieldUname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Passwd(); ok {
		_spec.SetField(user.FieldPasswd, field.TypeString, value)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uu.mutation.IsAdmin(); ok {
		_spec.SetField(user.FieldIsAdmin, field.TypeBool, value)
	}
	if value, ok := uu.mutation.IsSecretary(); ok {
		_spec.SetField(user.FieldIsSecretary, field.TypeBool, value)
	}
	if value, ok := uu.mutation.IsActive(); ok {
		_spec.SetField(user.FieldIsActive, field.TypeBool, value)
	}
	if uu.mutation.RequestedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RequestedTable,
			Columns: []string{user.RequestedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRequestedIDs(); len(nodes) > 0 && !uu.mutation.RequestedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RequestedTable,
			Columns: []string{user.RequestedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RequestedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RequestedTable,
			Columns: []string{user.RequestedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ReceivedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.ReceivedTable,
			Columns: user.ReceivedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		createE := &ReceiveCreate{config: uu.config, mutation: newReceiveMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedReceivedIDs(); len(nodes) > 0 && !uu.mutation.ReceivedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.ReceivedTable,
			Columns: user.ReceivedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReceiveCreate{config: uu.config, mutation: newReceiveMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ReceivedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.ReceivedTable,
			Columns: user.ReceivedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReceiveCreate{config: uu.config, mutation: newReceiveMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(i int64) *UserUpdateOne {
	uuo.mutation.ResetUpdatedAt()
	uuo.mutation.SetUpdatedAt(i)
	return uuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (uuo *UserUpdateOne) AddUpdatedAt(i int64) *UserUpdateOne {
	uuo.mutation.AddUpdatedAt(i)
	return uuo
}

// SetUname sets the "uname" field.
func (uuo *UserUpdateOne) SetUname(s string) *UserUpdateOne {
	uuo.mutation.SetUname(s)
	return uuo
}

// SetPasswd sets the "passwd" field.
func (uuo *UserUpdateOne) SetPasswd(s string) *UserUpdateOne {
	uuo.mutation.SetPasswd(s)
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetIsAdmin sets the "is_admin" field.
func (uuo *UserUpdateOne) SetIsAdmin(b bool) *UserUpdateOne {
	uuo.mutation.SetIsAdmin(b)
	return uuo
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsAdmin(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsAdmin(*b)
	}
	return uuo
}

// SetIsSecretary sets the "is_secretary" field.
func (uuo *UserUpdateOne) SetIsSecretary(b bool) *UserUpdateOne {
	uuo.mutation.SetIsSecretary(b)
	return uuo
}

// SetNillableIsSecretary sets the "is_secretary" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsSecretary(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsSecretary(*b)
	}
	return uuo
}

// SetIsActive sets the "is_active" field.
func (uuo *UserUpdateOne) SetIsActive(b bool) *UserUpdateOne {
	uuo.mutation.SetIsActive(b)
	return uuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsActive(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsActive(*b)
	}
	return uuo
}

// AddRequestedIDs adds the "requested" edge to the Order entity by IDs.
func (uuo *UserUpdateOne) AddRequestedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddRequestedIDs(ids...)
	return uuo
}

// AddRequested adds the "requested" edges to the Order entity.
func (uuo *UserUpdateOne) AddRequested(o ...*Order) *UserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.AddRequestedIDs(ids...)
}

// AddReceivedIDs adds the "received" edge to the Order entity by IDs.
func (uuo *UserUpdateOne) AddReceivedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddReceivedIDs(ids...)
	return uuo
}

// AddReceived adds the "received" edges to the Order entity.
func (uuo *UserUpdateOne) AddReceived(o ...*Order) *UserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.AddReceivedIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearRequested clears all "requested" edges to the Order entity.
func (uuo *UserUpdateOne) ClearRequested() *UserUpdateOne {
	uuo.mutation.ClearRequested()
	return uuo
}

// RemoveRequestedIDs removes the "requested" edge to Order entities by IDs.
func (uuo *UserUpdateOne) RemoveRequestedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveRequestedIDs(ids...)
	return uuo
}

// RemoveRequested removes "requested" edges to Order entities.
func (uuo *UserUpdateOne) RemoveRequested(o ...*Order) *UserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.RemoveRequestedIDs(ids...)
}

// ClearReceived clears all "received" edges to the Order entity.
func (uuo *UserUpdateOne) ClearReceived() *UserUpdateOne {
	uuo.mutation.ClearReceived()
	return uuo
}

// RemoveReceivedIDs removes the "received" edge to Order entities by IDs.
func (uuo *UserUpdateOne) RemoveReceivedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveReceivedIDs(ids...)
	return uuo
}

// RemoveReceived removes "received" edges to Order entities.
func (uuo *UserUpdateOne) RemoveReceived(o ...*Order) *UserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.RemoveReceivedIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Uname(); ok {
		if err := user.UnameValidator(v); err != nil {
			return &ValidationError{Name: "uname", err: fmt.Errorf(`ent: validator failed for field "User.uname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Passwd(); ok {
		if err := user.PasswdValidator(v); err != nil {
			return &ValidationError{Name: "passwd", err: fmt.Errorf(`ent: validator failed for field "User.passwd": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.Uname(); ok {
		_spec.SetField(user.FieldUname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Passwd(); ok {
		_spec.SetField(user.FieldPasswd, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IsAdmin(); ok {
		_spec.SetField(user.FieldIsAdmin, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.IsSecretary(); ok {
		_spec.SetField(user.FieldIsSecretary, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.IsActive(); ok {
		_spec.SetField(user.FieldIsActive, field.TypeBool, value)
	}
	if uuo.mutation.RequestedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RequestedTable,
			Columns: []string{user.RequestedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRequestedIDs(); len(nodes) > 0 && !uuo.mutation.RequestedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RequestedTable,
			Columns: []string{user.RequestedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RequestedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RequestedTable,
			Columns: []string{user.RequestedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ReceivedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.ReceivedTable,
			Columns: user.ReceivedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		createE := &ReceiveCreate{config: uuo.config, mutation: newReceiveMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedReceivedIDs(); len(nodes) > 0 && !uuo.mutation.ReceivedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.ReceivedTable,
			Columns: user.ReceivedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReceiveCreate{config: uuo.config, mutation: newReceiveMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ReceivedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.ReceivedTable,
			Columns: user.ReceivedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReceiveCreate{config: uuo.config, mutation: newReceiveMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
