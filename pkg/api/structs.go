package api

type (
	UpdateOrderRequest struct {
		OrderID string `json:"order_id" description:"Código de identificação do pedido"`
		Status  string `json:"status" description:"Status para qual deseja mudar o pedido" enum:"Recebido|Preparacao|Pronto|Finalizado|Cancelado"`
	}
	UpdateOrderResponse struct {
		OrderID string `json:"order_id,omitempty" description:"Código de identificação do pedido"`
		Status  string `json:"status,omitempty" description:"Status para qual deseja mudar o pedido" enum:"Recebido|Preparacao|Pronto|Finalizado|Cancelado"`
	}

	ErrorResponse struct {
		Description string `json:"description,omitempty"`
		Code        int    `json:"code,omitempty"`
	}
)
