package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	logger "github.com/SOAT1StackGoLang/msvc-payments/pkg/middleware"
	kitlog "github.com/go-kit/log"
	"net/http"
	"strings"
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
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}
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

	logger.Info(fmt.Sprintf("production url: %s", url))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
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
		return UpdateOrderResponse{}, errors.New(fmt.Sprintf("%s %v", "unexpected status code", resp.StatusCode))
	}

	var responseBody UpdateOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return UpdateOrderResponse{}, err
	}

	return responseBody, nil
}
