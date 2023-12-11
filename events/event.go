package events

import (
	"context"
	"cqrs_project/models"
)

type EventStore interface {
	Close()
	// PublishedCreatedFeed publicar nuevos feed
	PublishedCreatedFeed(ctx context.Context, feed *models.Feed) error
	//SubscribeCreatedFeed para que se suscriba cuando un nuevo feed ha sido creado
	SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error)
	//OnCreatedFeed es un callback que reaccionara cuando un nuevo mensaje ha sido creado
	OnCreatedFeed(f func(CreatedFeedMessage)) error
}

var eventStore EventStore

func Close() {
	eventStore.Close()
}

func PublishedCreatedFeed(ctx context.Context, feed *models.Feed) error {
	return eventStore.PublishedCreatedFeed(ctx, feed)
}

func SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error) {
	return eventStore.SubscribeCreatedFeed(ctx)
}

func OnCreatedFeed(f func(CreatedFeedMessage)) error {
	return eventStore.OnCreatedFeed(f)
}

func SetEventStore(store EventStore) {
	eventStore = store
}
