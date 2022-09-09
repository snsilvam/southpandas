package repository

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

type UserClientRepository interface {
	InsertUserClient(ctx context.Context, userClient *models.UserClient) error
	ListUsersClients(ctx context.Context) ([]*models.UserClient, error)
}

var implementationUserClient UserClientRepository

func SetClientRepository(repository UserClientRepository) {
	implementationUserClient = repository
}

//Implement the UserRepositoryClient interface, func insert userClient
func InsertUserClient(ctx context.Context, userClient *models.UserClient) error {
	return implementationUserClient.InsertUserClient(ctx, userClient)
}

//Implement the UserRepositoryClient interface, func list usersClient
func ListUsersClients(ctx context.Context) ([]*models.UserClient, error) {
	return implementationUserClient.ListUsersClients(ctx)
}
