package service

import (
	"context"
	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -destination=../mocks/cache.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/pkg/service CacheService
type cacheSvc struct {
	cacheClient redis.Client
}

func (c *cacheSvc) OrderStatusChanged(ctx context.Context, whereTo any, order Order) error {
	//TODO implement me
	panic("implement me")
}

type CacheService interface {
	OrderStatusChanged(ctx context.Context, whereTo any, order Order) error
}

func NewCacheService(client redis.Client) CacheService {
	return &cacheSvc{cacheClient: client}
}
