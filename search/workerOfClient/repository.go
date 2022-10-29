package search

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

type SearchRepository interface {
	Close()
	IndexWorkerOfClient(ctx context.Context, WorkerOfClient models.WorkerOfClient) error
	SearchWorkerOfClient(ctx context.Context, query string) ([]models.WorkerOfClient, error)
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}
func Close() {
	repo.Close()
}
func IndexWorkerOfClient(ctx context.Context, WorkerOfClient models.WorkerOfClient) error {
	return repo.IndexWorkerOfClient(ctx, WorkerOfClient)
}
func SearchWorkerOfClient(ctx context.Context, query string) ([]models.WorkerOfClient, error) {
	return repo.SearchWorkerOfClient(ctx, query)
}
