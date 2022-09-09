package search

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

type SearchRepository interface {
	Close()
	IndexUser(ctx context.Context, user models.User) error
	SearchUser(ctx context.Context, query string) ([]models.User, error)
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}
func Close() {
	repo.Close()
}
func IndexUser(ctx context.Context, user models.User) error {
	return repo.IndexUser(ctx, user)
}
func SearchUser(ctx context.Context, query string) ([]models.User, error) {
	return repo.SearchUser(ctx, query)
}
