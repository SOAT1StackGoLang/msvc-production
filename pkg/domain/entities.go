package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type OrderStatus string

const (
	ORDER_STATUS_UNSET           OrderStatus = ""
	ORDER_STATUS_OPEN                        = "Aberto"
	ORDER_STATUS_WAITING_PAYMENT             = "Aguardando Pagamento"
	ORDER_STATUS_RECEIVED                    = "Recebido"
	ORDER_STATUS_PREPARING                   = "Em Preparação"
	ORDER_STATUS_DONE                        = "Pronto"
	ORDER_STATUS_FINISHED                    = "Finalizado"
	ORDER_STATUS_CANCELED                    = "Cancelado"
)

type Order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PaymentID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Price     decimal.Decimal
	Status    OrderStatus
	Products  []Product
}

type Product struct {
	ID          uuid.UUID
	CategoryID  uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Name        string
	Description string
	Price       decimal.Decimal
}
