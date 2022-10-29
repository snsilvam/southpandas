package search

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

type SearchRepository interface {
	Close()
	IndexUserClient(ctx context.Context, userClient models.UserClient) error
	SearchUserClient(ctx context.Context, query string) ([]models.UserClient, error)
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}
func Close() {
	repo.Close()
}
func IndexUserClient(ctx context.Context, userClient models.UserClient) error {
	return repo.IndexUserClient(ctx, userClient)
}
func SearchUserClient(ctx context.Context, query string) ([]models.UserClient, error) {
	return repo.SearchUserClient(ctx, query)
}
