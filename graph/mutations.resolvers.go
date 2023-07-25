package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"errors"
	"shira-chan-dev/app/utils"
	"shira-chan-dev/ent"
	"shira-chan-dev/ent/privacy"
	"shira-chan-dev/ent/user"

	"golang.org/x/crypto/bcrypt"
)

// Sign is the resolver for the sign field.
func (r *mutationResolver) Sign(ctx context.Context, input SignInput) (*Token, error) {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)
	// 通过手机号查找用户
	u, err := r.client.User.Query().
		Where(user.PhoneEQ(input.Phone)).
		// 该查找不经过privacy规则
		Only(ctx)
	if input.Uname != nil {
		// 注册
		if u == nil {
			u, err = utils.Client.User.Create().
				SetUname(*input.Uname).
				SetPhone(input.Phone).
				SetPasswd(input.Passwd).
				Save(ctx)
		} else {
			return nil, errors.New("user already exists")
		}
	} else {
		// 登录
		if u == nil {
			return nil, errors.New("user does not exists")
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(u.Passwd), []byte(input.Passwd))
			if err != nil {
				return nil, errors.New("password incorrect")
			}
		}
	}
	if u != nil {
		// 生成token
		if !u.IsActive {
			return nil, errors.New("banned")
		} else {
			var token string
			token, err = utils.GetToken(u.ID, u.IsAdmin)
			return &Token{Token: &token, IsAdmin: &u.IsAdmin, IsSecretary: &u.IsSecretary}, err
		}
	}
	return nil, errors.New("bad request")
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	return client.User.Create().SetInput(input).Save(ctx)
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input ent.CreateOrderInput) (*ent.Order, error) {
	client := ent.FromContext(ctx)
	return client.Order.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	return client.User.UpdateOneID(id).SetInput(input).Save(ctx)
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, id int, input ent.UpdateOrderInput) (*ent.Order, error) {
	client := ent.FromContext(ctx)
	return client.Order.UpdateOneID(id).SetInput(input).Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
