package repository

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

// Implementing design patron Repository
type UserSouthpandasRepository interface {
	CloseUserSouthpandas()
	InsertUserSouthpandas(ctx context.Context, userSouthpandas *models.UserSouthPandas) error
	ListUsersSouthpandas(ctx context.Context) ([]*models.UserSouthPandas, error)
}

// Abstraction of db
var repositoryUserSouthpandas UserSouthpandasRepository

// Constructor
func SetRepositoryUserSouthpandas(r UserSouthpandasRepository) {
	repositoryUserSouthpandas = r
}

// Implement the UserSouthpandasRepository interface, func close
func CloseUserSouthpandas() {
	repositoryUserSouthpandas.CloseUserSouthpandas()
}

// Implement the UserSouthpandasRepository interface, func insert userSouthpandas
func InsertUserSouthpandas(ctx context.Context, userSouthpandas *models.UserSouthPandas) error {
	return repositoryUserSouthpandas.InsertUserSouthpandas(ctx, userSouthpandas)
}

// Implement the UserSouthpandasRepository interface, func list userSouthpandas
func ListUsersSouthpandas(ctx context.Context) ([]*models.UserSouthPandas, error) {
	return repositoryUserSouthpandas.ListUsersSouthpandas(ctx)
}
