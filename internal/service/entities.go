package service

import (
	"github.com/SOAT1StackGoLang/msvc-production/pkg/messages"
	"github.com/google/uuid"
	"time"
)

type OrderStatus string

const (
	ORDER_STATUS_UNSET     OrderStatus = ""
	ORDER_STATUS_RECEIVED              = "Recebido"
	ORDER_STATUS_PREPARING             = "Em Preparação"
	ORDER_STATUS_DONE                  = "Pronto"
	ORDER_STATUS_FINISHED              = "Finalizado"
	ORDER_STATUS_CANCELED              = "Cancelado"
)

type Order struct {
	ID        uuid.UUID   `json:"id"`
	UpdatedAt time.Time   `json:"updated_at"`
	Status    OrderStatus `json:"status"`
}

func (o *Order) ToProductionStatusChangedMessage() messages.ProductionStatusChangedMessage {
	return messages.ProductionStatusChangedMessage{
		OrderID:   o.ID.String(),
		Status:    string(o.Status),
		UpdatedAt: o.UpdatedAt.Format(time.RFC3339),
	}
}

func OrderFromProductionStatusChangedMessage(msg messages.ProductionStatusChangedMessage) (*Order, error) {
	id, err := uuid.Parse(msg.OrderID)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse(time.RFC3339, msg.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &Order{
		ID:        id,
		UpdatedAt: updatedAt,
		Status:    OrderStatus(msg.Status),
	}, nil
}
