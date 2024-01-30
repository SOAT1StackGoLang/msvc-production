package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	kitlog "github.com/go-kit/log"
	"net/http"
)

type client struct {
	baseURL string
	logger  kitlog.Logger
}

type ProductionAPI interface {
	UpdateOrder(request UpdateOrderRequest) (UpdateOrderResponse, error)
}

//go:generate mockgen -destination=../mocks/api_mocks.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/pkg/api ProductionAPI
func NewClient(baseURL string, logger kitlog.Logger) ProductionAPI {
	return &client{
		baseURL: baseURL,
		logger:  logger,
	}
}

func (c *client) UpdateOrder(request UpdateOrderRequest) (UpdateOrderResponse, error) {
	url := fmt.Sprintf("%s/production", c.baseURL)

	payload, err := json.Marshal(request)
	if err != nil {
		return UpdateOrderResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return UpdateOrderResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return UpdateOrderResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UpdateOrderResponse{}, errors.New("unexpected status code")
	}

	var responseBody UpdateOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return UpdateOrderResponse{}, err
	}

	return responseBody, nil
}
