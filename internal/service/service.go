package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

type productionSvc struct {
	cacheSvc CacheService
	pubSvc   Publisher
}

//go:generate mockgen -destination=../mocks/service.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/internal/service ProductionService
type ProductionService interface {
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status OrderStatus) (*Order, error)
}

func NewProductionService(cacheSvc CacheService, svc Publisher) ProductionService {
	return &productionSvc{
		cacheSvc: cacheSvc,
		pubSvc:   svc,
	}
}

func (p *productionSvc) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status OrderStatus) (*Order, error) {
	if status == ORDER_STATUS_UNSET {
		return nil, errors.New("invalid status")
	}
	order := &Order{
		ID:        orderID,
		UpdatedAt: time.Now(),
		Status:    status,
	}

	err := p.cacheSvc.SaveOrderStatusChanged(ctx, *order)
	if err != nil {
		return nil, err
	}

	err = p.pubSvc.PublishOrderStatusChanged(ctx, OrderStatusChannel, *order)

	return order, err
}
