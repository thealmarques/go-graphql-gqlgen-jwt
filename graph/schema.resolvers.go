package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql-jwt/graph/generated"
	"go-graphql-jwt/graph/model"
	"go-graphql-jwt/graph/services"
	"time"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	now := int(time.Now().Unix())
	user := &model.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := services.DB.Exec("INSERT INTO `users` (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	user.ID = lastID
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	user1 := &model.User{
		CreatedAt: 111,
		Name:      "Andre",
		Email:     "andre@gmail.com",
		ID:        1,
	}
	users = append(users, user1)
	return users, nil
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
