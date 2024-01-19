package transport

import (
	"context"
	"github.com/SOAT1StackGoLang/msvc-production/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"net/http"
)

type (
	Endpoints struct {
		UpdateOrderStatusEndpoint endpoint.Endpoint
	}

	updateOrderRequest struct {
		OrderID string `json:"order_id" description:"Código de identificação do pedido"`
		UserID  string `json:"user_id" description:"Código de descrição do usuário requerente"`
		Status  string `json:"status" description:"Status para qual deseja mudar o pedido" enum:"Recebido|Preparacao|Pronto|Finalizado|Cancelado"`
	}
	updateOrderResponse struct {
		OrderID string        `json:"order_id,omitempty" description:"Código de identificação do pedido"`
		Status  string        `json:"status,omitempty" description:"Status para qual deseja mudar o pedido" enum:"Recebido|Preparacao|Pronto|Finalizado|Cancelado"`
		Error   errorResponse `json:"error,omitempty" description:"Detalhamento do erro"`
	}

	errorResponse struct {
		Description string `json:"description"`
		Code        int    `json:"code"`
	}
)

func MakeServerEndpoints(svc service.ProductionService) Endpoints {
	return Endpoints{
		UpdateOrderStatusEndpoint: makeUpdateOrderStatusEndpoint(svc),
	}
}

func makeUpdateOrderStatusEndpoint(svc service.ProductionService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(updateOrderRequest)

		_, err := svc.UpdateOrderStatus(ctx, uuid.MustParse(req.UserID), uuid.MustParse(req.OrderID), service.OrderStatus(req.Status))
		if err != nil {
			return updateOrderResponse{Error: errorResponse{
				Description: "Erro interno do servidor",
				Code:        http.StatusInternalServerError,
			}}, err
		}

		return updateOrderResponse{
			OrderID: req.OrderID,
			Status:  req.Status,
		}, nil
	}
}
