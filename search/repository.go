package search

import (
	"context"
	"cqrs_project/models"
)

type SearchRepository interface {
	Close() error
	IndexFeed(ctx context.Context, feed models.Feed) error
	SearchFeed(ctx context.Context, query string) ([]models.Feed, error)
}

var repo SearchRepository

func SetSearchRepositroy(r SearchRepository) {
	repo = r
}
func Close() {
	repo.Close()
}

func IndexFeed(ctx context.Context, feed models.Feed) error {
	return repo.IndexFeed(ctx, feed)
}
func SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	return repo.SearchFeed(ctx, query)
}

//go get github.com/elastic/go-elasticsearch/v7
