package domain

import (
	"github.com/google/uuid"
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
	UpdatedAt time.Time
	Status    OrderStatus
}
