// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (o *Order) Requester(ctx context.Context) (*User, error) {
	result, err := o.Edges.RequesterOrErr()
	if IsNotLoaded(err) {
		result, err = o.QueryRequester().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (o *Order) Receiver(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *UserOrder,
) (*UserConnection, error) {
	opts := []UserPaginateOption{
		WithUserOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := o.Edges.totalCount[1][alias]
	if nodes, err := o.NamedReceiver(alias); err == nil || hasTotalCount {
		pager, err := newUserPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &UserConnection{Edges: []*UserEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return o.QueryReceiver().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Requested(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *OrderOrder,
) (*OrderConnection, error) {
	opts := []OrderPaginateOption{
		WithOrderOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[0][alias]
	if nodes, err := u.NamedRequested(alias); err == nil || hasTotalCount {
		pager, err := newOrderPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &OrderConnection{Edges: []*OrderEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryRequested().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Received(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *OrderOrder,
) (*OrderConnection, error) {
	opts := []OrderPaginateOption{
		WithOrderOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[1][alias]
	if nodes, err := u.NamedReceived(alias); err == nil || hasTotalCount {
		pager, err := newOrderPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &OrderConnection{Edges: []*OrderEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryReceived().Paginate(ctx, after, first, before, last, opts...)
}