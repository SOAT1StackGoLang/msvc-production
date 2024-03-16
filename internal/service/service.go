package service

import (
	"context"
	"errors"
	"github.com/SOAT1StackGoLang/msvc-production/pkg"
	"github.com/google/uuid"
	"time"
)

type productionSvc struct {
	cacheSvc CacheService
	pubSvc   Publisher
}

func (p *productionSvc) SubscribeToIncomingOrders() {
	//TODO implement me
	panic("implement me")
}

//go:generate mockgen -destination=../mocks/service.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/internal/service ProductionService
type ProductionService interface {
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status OrderStatus) (*Order, error)
	SubscribeToIncomingOrders()
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

	out := order.ToProductionStatusChangedMessage()

	err = p.pubSvc.PublishOrderStatusChanged(ctx, pkg.OrderStatusChannel, out)

	return order, err
}
