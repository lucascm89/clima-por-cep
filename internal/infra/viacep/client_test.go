package viacep

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"clima-por-cep/internal/usecase"
)

func TestClientReturnsLocationWhenZipcodeExists(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/01001000/json/" {
			t.Fatalf("expected path /01001000/json/, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"cep": "01001-000",
			"localidade": "São Paulo",
			"uf": "SP"
		}`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL)

	location, err := client.GetLocationByZipcode(context.Background(), "01001000")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if location.City != "São Paulo" {
		t.Fatalf("expected city São Paulo, got %s", location.City)
	}

	if location.State != "SP" {
		t.Fatalf("expected state SP, got %s", location.State)
	}
}

func TestClientReturnsZipcodeNotFoundWhenViaCEPReturnsErroTrue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"erro": true
		}`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL)

	location, err := client.GetLocationByZipcode(context.Background(), "99999999")

	if location != nil {
		t.Fatalf("expected location to be nil")
	}

	if !errors.Is(err, usecase.ErrZipcodeNotFound) {
		t.Fatalf("expected ErrZipcodeNotFound, got %v", err)
	}
}

func TestClientReturnsZipcodeNotFoundWhenLocalidadeIsEmpty(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"cep": "01001-000",
			"localidade": "",
			"uf": ""
		}`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL)

	location, err := client.GetLocationByZipcode(context.Background(), "01001000")

	if location != nil {
		t.Fatalf("expected location to be nil")
	}

	if !errors.Is(err, usecase.ErrZipcodeNotFound) {
		t.Fatalf("expected ErrZipcodeNotFound, got %v", err)
	}
}

func TestClientReturnsErrorWhenViaCEPReturnsNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL)

	location, err := client.GetLocationByZipcode(context.Background(), "01001000")

	if location != nil {
		t.Fatalf("expected location to be nil")
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

	client := NewClientWithBaseURL(server.URL)

	location, err := client.GetLocationByZipcode(context.Background(), "01001000")

	if location != nil {
		t.Fatalf("expected location to be nil")
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestClientReturnsZipcodeNotFoundWhenViaCEPReturnsErroAsString(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"erro": "true"
		}`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL)

	location, err := client.GetLocationByZipcode(context.Background(), "99999999")

	if location != nil {
		t.Fatalf("expected location to be nil")
	}

	if !errors.Is(err, usecase.ErrZipcodeNotFound) {
		t.Fatalf("expected ErrZipcodeNotFound, got %v", err)
	}
}
