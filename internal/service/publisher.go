package service

import (
	"context"
	"encoding/json"
	"github.com/SOAT1StackGoLang/msvc-payments/pkg/datastore"
)

var OrderStatusChannel = "order_status_changed"

type pubSvc struct {
	pubSvc datastore.RedisStore
}

//go:generate mockgen -destination=../mocks/pubsub_mocks.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/internal/service Publisher
func (p *pubSvc) PublishOrderStatusChanged(ctx context.Context, channel string, order Order) error {
	marshalled, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return p.pubSvc.Publish(ctx, channel, marshalled)
}

type (
	Publisher interface {
		PublishOrderStatusChanged(ctx context.Context, channel string, order Order) error
	}
)

func NewPublisher(client datastore.RedisStore) Publisher {
	return &pubSvc{pubSvc: client}
}
