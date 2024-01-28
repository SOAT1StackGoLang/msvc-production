package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SOAT1StackGoLang/msvc-payments/pkg/datastore"
	kitlog "github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"time"
)

const productionPrefix = "production_"

var ErrRecordNotFound = errors.New("no matching key found")

//go:generate mockgen -destination=../mocks/cache_mocks.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/internal/service CacheService
type cacheSvc struct {
	cacheClient datastore.RedisStore
	logger      kitlog.Logger
}

func (c *cacheSvc) GetOrder(ctx context.Context, orderID uuid.UUID) (*Order, error) {
	var out *Order
	order, err := c.cacheClient.Get(ctx, fmt.Sprintf("%s%s", productionPrefix, orderID))
	if err != nil && order == "" {
		return nil, ErrRecordNotFound
	}

	err = json.Unmarshal([]byte(order), out)
	return out, err
}

func (c *cacheSvc) SaveOrderStatusChanged(ctx context.Context, order Order) error {
	orderJson, err := json.Marshal(order)
	if err != nil {
		c.logger.Log("failed marshaling order")
		return err
	}
	err = c.cacheClient.Set(ctx, fmt.Sprintf("%s%s", productionPrefix, order.ID), orderJson, time.Hour*24)
	if err != nil {
		c.logger.Log("failed persisting into cache")
	}
	return err
}

type CacheService interface {
	SaveOrderStatusChanged(ctx context.Context, order Order) error
	GetOrder(ctx context.Context, orderID uuid.UUID) (*Order, error)
}

func NewCacheService(client datastore.RedisStore, log kitlog.Logger) CacheService {
	return &cacheSvc{cacheClient: client, logger: log}
}
