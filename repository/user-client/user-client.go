package repository

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

// Implementing design patron Repository
type UserClientRepository interface {
	Close()
	InsertUserClient(ctx context.Context, userClient *models.UserClient) error
	ListUsersClients(ctx context.Context) ([]*models.UserClient, error)
}

// Abstraction of db
var repositoryUserClient UserClientRepository

// Constructor
func SetClientRepository(repository UserClientRepository) {
	repositoryUserClient = repository
}

// Implement the UserRepository interface, func close
func Close() {
	repositoryUserClient.Close()
}

// Implement the UserRepositoryClient interface, func insert userClient
func InsertUserClient(ctx context.Context, userClient *models.UserClient) error {
	return repositoryUserClient.InsertUserClient(ctx, userClient)
}

// Implement the UserRepositoryClient interface, func list usersClient
func ListUsersClients(ctx context.Context) ([]*models.UserClient, error) {
	return repositoryUserClient.ListUsersClients(ctx)
}
