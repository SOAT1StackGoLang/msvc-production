package service

import (
	"context"
	"github.com/SOAT1StackGoLang/msvc-production/pkg/domain"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type productionSvc struct {
	cacheClient redis.Client
}

type ProductionService interface {
	UpdateOrderStatus(ctx context.Context, userID, orderID uuid.UUID, status domain.OrderStatus) (*domain.Order, error)
}

func NewProductionService() ProductionService {
	return &productionSvc{}
}

func (p *productionSvc) WithCache(client redis.Client) ProductionService {
	p.cacheClient = client
	return p
}

func (p *productionSvc) UpdateOrderStatus(ctx context.Context, userID, orderID uuid.UUID, status domain.OrderStatus) (*domain.Order, error) {
	return nil, nil
}
