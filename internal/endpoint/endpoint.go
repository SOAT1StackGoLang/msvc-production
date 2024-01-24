package endpoint

import (
	"context"
	"github.com/SOAT1StackGoLang/msvc-production/internal/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"net/http"
)

type (
	Endpoints struct {
		UpdateOrderStatusEndpoint endpoint.Endpoint
	}

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

func MakeServerEndpoints(svc service.ProductionService) Endpoints {
	return Endpoints{
		UpdateOrderStatusEndpoint: makeUpdateOrderStatusEndpoint(svc),
	}
}

func makeUpdateOrderStatusEndpoint(svc service.ProductionService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(UpdateOrderRequest)

		orderID, err := uuid.Parse(req.OrderID)
		if err != nil {
			return ErrorResponse{
				Description: "invalid order id",
				Code:        http.StatusBadRequest,
			}, err
		}
		_, err = svc.UpdateOrderStatus(ctx, orderID, service.OrderStatus(req.Status))
		if err != nil {
			return ErrorResponse{
				Description: "Erro interno do servidor",
				Code:        http.StatusInternalServerError,
			}, err
		}

		return UpdateOrderResponse{
			OrderID: req.OrderID,
			Status:  req.Status,
		}, nil
	}
}
