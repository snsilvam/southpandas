package repository

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

// Implementing design patron Repository
type WorkerOfClientRepository interface {
	CloseWorkerOfClient()
	InsertWorkerOfClient(ctx context.Context, workerOfClient *models.WorkerOfClient) error
	ListWorkerOfClient(ctx context.Context) ([]*models.WorkerOfClient, error)
}

// Abstraction of db
var repositoryWorkerOfClient WorkerOfClientRepository

// Constructor
func SetRepositoryWorkerOfClient(r WorkerOfClientRepository) {
	repositoryWorkerOfClient = r
}

// Implement the WorkerOfClientRepository interface, func close
func CloseWorkerOfClient() {
	repositoryWorkerOfClient.CloseWorkerOfClient()
}

// Implement the WorkerOfClientRepository interface, func insert workerOfClient
func InsertWorkerOfClient(ctx context.Context, workerOfClient *models.WorkerOfClient) error {
	return repositoryWorkerOfClient.InsertWorkerOfClient(ctx, workerOfClient)
}

// Implement the WorkerOfClientRepository interface, func list workerOfClient
func ListWorkerOfClient(ctx context.Context) ([]*models.WorkerOfClient, error) {
	return repositoryWorkerOfClient.ListWorkerOfClient(ctx)
}
