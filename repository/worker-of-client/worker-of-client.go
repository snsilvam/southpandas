package repository

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

// Implementing design patron Repository
type WorkerOfClientRepository interface {
	Close()
	InsertWorkerOfClient(ctx context.Context, workerOfClient *models.WorkerOfClient) error
	ListWorkersOfClient(ctx context.Context) ([]*models.WorkerOfClient, error)
}

// Abstraction of db
var repositoryWorkerOfClient WorkerOfClientRepository

// Constructor
func SetRepositoryWorkerOfClient(r WorkerOfClientRepository) {
	repositoryWorkerOfClient = r
}

// Implement the WorkerOfClientRepository interface, func close
func Close() {
	repositoryWorkerOfClient.Close()
}

// Implement the WorkerOfClientRepository interface, func insert workerOfClient
func InsertWorkerOfClient(ctx context.Context, workerOfClient *models.WorkerOfClient) error {
	return repositoryWorkerOfClient.InsertWorkerOfClient(ctx, workerOfClient)
}

// Implement the WorkerOfClientRepository interface, func list workerOfClient
func ListWorkersOfClient(ctx context.Context) ([]*models.WorkerOfClient, error) {
	return repositoryWorkerOfClient.ListWorkersOfClient(ctx)
}
