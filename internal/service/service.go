package service

import (
	"context"
	"errors"
	"github.com/SOAT1StackGoLang/msvc-production/pkg/messages"
	"github.com/google/uuid"
	"time"
)

type productionSvc struct {
	cacheSvc CacheService
	pubSvc   Publisher
}

func (p *productionSvc) SubscribeToIncomingOrders() {
	go p.cacheSvc.SubscribeToIncomingOrders()
}

//go:generate mockgen -destination=../mocks/service.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/internal/service ProductionService
type ProductionService interface {
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status OrderStatus) (*Order, error)
	SubscribeToIncomingOrders()
}

func NewProductionService(cacheSvc CacheService, svc Publisher) ProductionService {
	pS := &productionSvc{
		cacheSvc: cacheSvc,
		pubSvc:   svc,
	}

	go pS.SubscribeToIncomingOrders()

	return pS
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

	cached, err := p.cacheSvc.GetOrder(ctx, orderID)
	if err == nil && cached.Status == ORDER_STATUS_CANCELED {
		return nil, errors.New("order already canceled")
	}

	err = p.cacheSvc.SaveOrderStatusChanged(ctx, *order)
	if err != nil {
		return nil, err
	}

	out := order.ToProductionStatusChangedMessage()

	err = p.pubSvc.PublishOrderStatusChanged(ctx, messages.ProductionStatusChannel, out)

	return order, err
}
