package service

import (
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
