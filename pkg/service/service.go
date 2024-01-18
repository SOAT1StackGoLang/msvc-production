package service

import (
	"context"
	"github.com/SOAT1StackGoLang/msvc-production/pkg/domain"
	"github.com/google/uuid"
)

type productionSvc struct {
	cacheSvc CacheService
}

//go:generate mockgen -destination=../mocks/service.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/pkg/service ProductionService
type ProductionService interface {
	UpdateOrderStatus(ctx context.Context, userID, orderID uuid.UUID, status domain.OrderStatus) (*domain.Order, error)
}

func NewProductionService(cacheSvc CacheService) ProductionService {
	return &productionSvc{}
}

func (p *productionSvc) UpdateOrderStatus(ctx context.Context, userID, orderID uuid.UUID, status domain.OrderStatus) (*domain.Order, error) {
	return nil, nil
}
