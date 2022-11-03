package repository

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

// Implementing design patron Repository
type UserExternalWorkerRepository interface {
	Close()
	InsertUserExternalWorker(ctx context.Context, userExternalWorker *models.UserExternalWorker) error
	ListUsersExternalWorkers(ctx context.Context) ([]*models.UserExternalWorker, error)
}

// Abstraction of db
var repositoryUserExternal UserExternalWorkerRepository

// Constructor
func SetRepositoryUserExternalWorker(r UserExternalWorkerRepository) {
	repositoryUserExternal = r
}

// Implement the UserExternalWorkerRepository interface, func close
func Close() {
	repositoryUserExternal.Close()
}

// Implement the UserExternalWorkerRepository interface, func insert userExternalWorker
func InsertUserExternalWorker(ctx context.Context, userExternalWorker *models.UserExternalWorker) error {
	return repositoryUserExternal.InsertUserExternalWorker(ctx, userExternalWorker)
}

// Implement the UserExternalWorkerRepository interface, func list userExternalWorker
func ListUsersExternalWorkers(ctx context.Context) ([]*models.UserExternalWorker, error) {
	return repositoryUserExternal.ListUsersExternalWorkers(ctx)
}
