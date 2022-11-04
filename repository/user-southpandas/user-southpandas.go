package repository

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

// Implementing design patron Repository
type UserSouthpandasRepository interface {
	Close()
	InsertUserSouthPandas(ctx context.Context, userSouthpandas *models.UserSouthPandas) error
	ListUsersSouthPandas(ctx context.Context) ([]*models.UserSouthPandas, error)
}

// Abstraction of db
var repositoryUserSouthpandas UserSouthpandasRepository

// Constructor
func SetRepositoryUserSouthpandas(r UserSouthpandasRepository) {
	repositoryUserSouthpandas = r
}

// Implement the UserSouthpandasRepository interface, func close
func Close() {
	repositoryUserSouthpandas.Close()
}

// Implement the UserSouthpandasRepository interface, func insert userSouthpandas
func InsertUserSouthPandas(ctx context.Context, userSouthpandas *models.UserSouthPandas) error {
	return repositoryUserSouthpandas.InsertUserSouthPandas(ctx, userSouthpandas)
}

// Implement the UserSouthpandasRepository interface, func list userSouthpandas
func ListUsersSouthPandas(ctx context.Context) ([]*models.UserSouthPandas, error) {
	return repositoryUserSouthpandas.ListUsersSouthPandas(ctx)
}
