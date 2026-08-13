package weatherapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const defaultBaseURL = "https://api.weatherapi.com/v1/current.json"

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type Response struct {
	Current Current `json:"current"`
}

type Current struct {
	TempC float64 `json:"temp_c"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func NewClientWithBaseURL(apiKey string, baseURL string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) GetCurrentTemperatureCelsius(
	ctx context.Context,
	city string,
) (float64, error) {
	if c.apiKey == "" {
		return 0, errors.New("weather api key is required")
	}

	if city == "" {
		return 0, errors.New("city is required")
	}

	requestURL, err := c.buildRequestURL(city)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("weatherapi returned status code %d: %s", resp.StatusCode, string(body))
	}

	var weatherAPIResponse Response

	if err := json.NewDecoder(resp.Body).Decode(&weatherAPIResponse); err != nil {
		return 0, err
	}

	return weatherAPIResponse.Current.TempC, nil
}

func (c *Client) buildRequestURL(city string) (string, error) {
	parsedURL, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	query := parsedURL.Query()
	query.Set("key", c.apiKey)
	query.Set("q", city)
	query.Set("aqi", "no")

	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}
