package messages

type OrderSentMessage struct {
	OrderID string      `json:"order_id"`
	Status  OrderStatus `json:"status"`
}

type ProductionStatusChangedMessage struct {
	OrderID   string `json:"order_id"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

type OrderStatus string

const (
	ORDER_STATUS_UNSET     OrderStatus = ""
	ORDER_STATUS_RECEIVED              = "Recebido"
	ORDER_STATUS_PREPARING             = "Em Preparação"
	ORDER_STATUS_DONE                  = "Pronto"
	ORDER_STATUS_FINISHED              = "Finalizado"
	ORDER_STATUS_CANCELED              = "Cancelado"
)
