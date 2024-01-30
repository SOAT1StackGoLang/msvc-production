package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/SOAT1StackGoLang/msvc-production/internal/endpoint"
	"github.com/SOAT1StackGoLang/msvc-production/internal/service"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrBadRequest = errors.New("parametros incorretos")
)

type errorer interface {
	error() error
}

func NewHttpHandler(pS service.ProductionService, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()
	sE := endpoint.MakeServerEndpoints(pS)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// POST /producao/ altera o status do pedido

	r.Methods(http.MethodPost).Path("/production").Handler(httptransport.NewServer(
		sE.UpdateOrderStatusEndpoint,
		decodeUpdateOrderRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeUpdateOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.UpdateOrderRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return endpoint.ErrorResponse{
			Description: "invalid request",
			Code:        http.StatusBadRequest,
		}, ErrBadRequest
	}

	return endpoint.UpdateOrderRequest{
		OrderID: req.OrderID,
		Status:  req.Status,
	}, nil
}
