package service

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type productionSvc struct {
	cacheSvc CacheService
}

//go:generate mockgen -destination=../mocks/service.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/pkg/service ProductionService
type ProductionService interface {
	UpdateOrderStatus(ctx context.Context, userID, orderID uuid.UUID, status OrderStatus) (*Order, error)
}

func NewProductionService(cacheSvc CacheService) ProductionService {
	return &productionSvc{}
}

func (p *productionSvc) UpdateOrderStatus(ctx context.Context, userID, orderID uuid.UUID, status OrderStatus) (*Order, error) {
	order := &Order{
		ID:        orderID,
		UserID:    userID,
		UpdatedAt: time.Now(),
		Status:    status,
	}

	err := p.cacheSvc.OrderStatusChanged(ctx, "alguma key? pubsub?", *order) // TODO persist in cache

	return order, err
}
