package search

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

type SearchRepository interface {
	Close()
	IndexUserSouthPandas(ctx context.Context, userSouthPandas models.UserSouthPandas) error
	SearchUserSouthPandas(ctx context.Context, query string) ([]models.UserSouthPandas, error)
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}
func Close() {
	repo.Close()
}
func IndexUserSouthPandas(ctx context.Context, userSouthPandas models.UserSouthPandas) error {
	return repo.IndexUserSouthPandas(ctx, userSouthPandas)
}
func SearchUserSouthPandas(ctx context.Context, query string) ([]models.UserSouthPandas, error) {
	return repo.SearchUserSouthPandas(ctx, query)
}
