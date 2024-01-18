package endpoint

import (
	"context"
	"errors"
	"github.com/SOAT1StackGoLang/msvc-production/pkg/domain"
	"github.com/SOAT1StackGoLang/msvc-production/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"net/http"
)

type (
	orderRequest struct {
		OrderID string `json:"order_id" description:"Código de identificação do pedido"`
		UserID  string `json:"user_id" description:"Código de descrição do usuário requerente"`
		Status  string `json:"status" description:"Status para qual deseja mudar o pedido" enum:"Recebido|Preparacao|Pronto|Finalizado|Cancelado"`
	}
	orderResponse struct {
		OrderID string        `json:"order_id,omitempty" description:"Código de identificação do pedido"`
		Status  string        `json:"status,omitempty" description:"Status para qual deseja mudar o pedido" enum:"Recebido|Preparacao|Pronto|Finalizado|Cancelado"`
		Error   errorResponse `json:"error,omitempty" description:"Detalhamento do erro"`
	}

	errorResponse struct {
		Description string `json:"description"`
		Code        int    `json:"code"`
	}
)

func makeStoreEndpoint(svc service.ProductionService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(orderRequest)
		if !ok {
			return orderResponse{Error: errorResponse{
				Description: "Dado fornecido incorreto",
				Code:        http.StatusBadRequest,
			}}, errors.New("novo erro")
		}
		orderID, err := uuid.Parse(req.OrderID)
		if err != nil {
			return orderResponse{Error: errorResponse{
				Description: "ID do Pedido incorreto",
				Code:        http.StatusBadRequest,
			}}, err
		}
		userID, err := uuid.Parse(req.UserID)
		if err != nil {
			return orderResponse{Error: errorResponse{
				Description: "ID do Usuário incorreto",
				Code:        http.StatusBadRequest,
			}}, err
		}

		_, err = svc.UpdateOrderStatus(ctx, userID, orderID, domain.OrderStatus(req.Status))
		if err != nil {
			return orderResponse{Error: errorResponse{
				Description: "Erro interno do servidor",
				Code:        http.StatusInternalServerError,
			}}, err
		}

		return orderResponse{
			OrderID: req.OrderID,
			Status:  req.Status,
		}, nil
	}
}
