package weatherapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientReturnsCurrentTemperatureCelsius(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/current.json" {
			t.Fatalf("expected path /v1/current.json, got %s", r.URL.Path)
		}

		if r.URL.Query().Get("key") != "test-key" {
			t.Fatalf("expected key test-key, got %s", r.URL.Query().Get("key"))
		}

		if r.URL.Query().Get("q") != "São Paulo" {
			t.Fatalf("expected q São Paulo, got %s", r.URL.Query().Get("q"))
		}

		if r.URL.Query().Get("aqi") != "no" {
			t.Fatalf("expected aqi no, got %s", r.URL.Query().Get("aqi"))
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"current": {
				"temp_c": 28.5
			}
		}`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL("test-key", server.URL+"/v1/current.json")

	tempC, err := client.GetCurrentTemperatureCelsius(context.Background(), "São Paulo")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tempC != 28.5 {
		t.Fatalf("expected tempC 28.5, got %.2f", tempC)
	}
}

func TestClientReturnsErrorWhenAPIKeyIsMissing(t *testing.T) {
	client := NewClientWithBaseURL("", "http://example.com")

	tempC, err := client.GetCurrentTemperatureCelsius(context.Background(), "São Paulo")

	if tempC != 0 {
		t.Fatalf("expected tempC 0, got %.2f", tempC)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestClientReturnsErrorWhenCityIsMissing(t *testing.T) {
	client := NewClientWithBaseURL("test-key", "http://example.com")

	tempC, err := client.GetCurrentTemperatureCelsius(context.Background(), "")

	if tempC != 0 {
		t.Fatalf("expected tempC 0, got %.2f", tempC)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestClientReturnsErrorWhenWeatherAPIReturnsNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	client := NewClientWithBaseURL("test-key", server.URL+"/v1/current.json")

	tempC, err := client.GetCurrentTemperatureCelsius(context.Background(), "São Paulo")

	if tempC != 0 {
		t.Fatalf("expected tempC 0, got %.2f", tempC)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestClientReturnsErrorWhenResponseIsInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`invalid-json`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL("test-key", server.URL+"/v1/current.json")

	tempC, err := client.GetCurrentTemperatureCelsius(context.Background(), "São Paulo")

	if tempC != 0 {
		t.Fatalf("expected tempC 0, got %.2f", tempC)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}
