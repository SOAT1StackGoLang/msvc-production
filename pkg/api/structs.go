package api

type (
	UpdateOrderRequest struct {
		OrderID string      `json:"order_id" description:"Código de identificação do pedido"`
		Status  OrderStatus `json:"status" description:"Status para qual deseja mudar o pedido" enum:"Recebido|Preparacao|Pronto|Finalizado|Cancelado"`
	}
	UpdateOrderResponse struct {
		OrderID string      `json:"order_id,omitempty" description:"Código de identificação do pedido"`
		Status  OrderStatus `json:"status,omitempty" description:"Status para qual deseja mudar o pedido" enum:"Recebido|Preparacao|Pronto|Finalizado|Cancelado"`
	}

	ErrorResponse struct {
		Description string `json:"description,omitempty"`
		Code        int    `json:"code,omitempty"`
	}
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
