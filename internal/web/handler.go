package web

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"clima-por-cep/internal/entity"
	"clima-por-cep/internal/usecase"
)

type WeatherUseCase interface {
	Execute(ctx context.Context, zipcode string) (*entity.WeatherResponse, error)
}

type Handler struct {
	weatherUseCase WeatherUseCase
}

func NewHandler(weatherUseCase WeatherUseCase) *Handler {
	return &Handler{
		weatherUseCase: weatherUseCase,
	}
}

func (h *Handler) GetWeatherByZipcode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	zipcode := strings.TrimPrefix(r.URL.Path, "/weather/")

	output, err := h.weatherUseCase.Execute(r.Context(), zipcode)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(output)
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, usecase.ErrInvalidZipcode):
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(usecase.ErrInvalidZipcode.Error()))

	case errors.Is(err, usecase.ErrZipcodeNotFound):
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(usecase.ErrZipcodeNotFound.Error()))

	default:
		log.Printf("unexpected error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
	}
}
