package repository

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

// Implementing design patron Repository
type UserRepository interface {
	CloseUser()
	InsertUser(ctx context.Context, user *models.User) error
	ListUsers(ctx context.Context) ([]*models.User, error)
}

// Abstraction of db
var repository UserRepository

// Constructor
func SetRepository(r UserRepository) {
	repository = r
}

// Implement the UserRepository interface, func close
func CloseUser() {
	repository.CloseUser()
}

// Implement the UserRepository interface, func insert user
func InsertUser(ctx context.Context, user *models.User) error {
	return repository.InsertUser(ctx, user)
}

// Implement the UserRepository interface, func list users
func ListUsers(ctx context.Context) ([]*models.User, error) {
	return repository.ListUsers(ctx)
}
