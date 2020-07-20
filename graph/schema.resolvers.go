package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"go-graphql-jwt/graph/generated"
	"go-graphql-jwt/graph/middlewares"
	"go-graphql-jwt/graph/model"
	"go-graphql-jwt/graph/services"
	"go-graphql-jwt/graph/utils"
	"time"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	now := int(time.Now().Unix())
	passwordHash, err := utils.HashPassword(input.Password)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  passwordHash,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = services.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.Token, error) {
	user, err := services.FindUserByEmail(email)

	if err != nil || user == nil {
		return nil, errors.New("User not found")
	}

	if !utils.ComparePassword(password, user.Password) {
		return nil, errors.New("Passwords doesn't match")
	}

	expiredAt := int(time.Now().Add(time.Hour * 1).Unix())
	obj := &model.Token{
		Token:     utils.GenerateJwt(user.ID, int64(expiredAt)),
		ExpiredAt: expiredAt,
	}

	return obj, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	userAuth := middlewares.GetAuthFromContext(ctx)

	if userAuth.UserID == 0 {
		return nil, errors.New("Access denied")
	}

	return services.FindUsers()
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (int, error) {
	return int(obj.ID), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
