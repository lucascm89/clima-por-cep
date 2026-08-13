package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"clima-por-cep/internal/usecase"
)

const defaultBaseURL = "https://viacep.com.br/ws"

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type Response struct {
	CEP        string      `json:"cep"`
	Localidade string      `json:"localidade"`
	UF         string      `json:"uf"`
	Erro       interface{} `json:"erro"`
}

func NewClient() *Client {
	return &Client{
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func NewClientWithBaseURL(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) GetLocationByZipcode(
	ctx context.Context,
	zipcode string,
) (*usecase.Location, error) {
	url := fmt.Sprintf("%s/%s/json/", c.baseURL, zipcode)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("viacep returned status code %d", resp.StatusCode)
	}

	var viaCEPResponse Response

	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResponse); err != nil {
		return nil, err
	}

	if isViaCEPError(viaCEPResponse.Erro) {
		return nil, usecase.ErrZipcodeNotFound
	}

	if viaCEPResponse.Localidade == "" {
		return nil, usecase.ErrZipcodeNotFound
	}

	return &usecase.Location{
		City:  viaCEPResponse.Localidade,
		State: viaCEPResponse.UF,
	}, nil
}

func isViaCEPError(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v

	case string:
		return v == "true"

	default:
		return false
	}
}
