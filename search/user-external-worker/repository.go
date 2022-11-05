package search

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

type SearchRepository interface {
	Close()
	IndexUserExternalWorker(ctx context.Context, userExternalWorker models.UserExternalWorker) error
	SearchUserExternalWorker(ctx context.Context, query string) ([]models.UserExternalWorker, error)
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}
func Close() {
	repo.Close()
}
func IndexUserExternalWorker(ctx context.Context, userExternalWorker models.UserExternalWorker) error {
	return repo.IndexUserExternalWorker(ctx, userExternalWorker)
}
func SearchUserExternalWorker(ctx context.Context, query string) ([]models.UserExternalWorker, error) {
	return repo.SearchUserExternalWorker(ctx, query)
}
