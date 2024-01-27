package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	kitlog "github.com/go-kit/log"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	logger     kitlog.Logger
}

func NewClient(baseURL string, httpClient *http.Client, logger kitlog.Logger) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		logger:     logger,
	}
}

func (c *Client) UpdateOrder(request UpdateOrderRequest) (UpdateOrderResponse, error) {
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

	resp, err := c.httpClient.Do(req)
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
