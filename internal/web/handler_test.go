package web

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"clima-por-cep/internal/entity"
	"clima-por-cep/internal/usecase"
)

type fakeWeatherUseCase struct {
	output          *entity.WeatherResponse
	err             error
	receivedZipcode string
}

func (f *fakeWeatherUseCase) Execute(
	ctx context.Context,
	zipcode string,
) (*entity.WeatherResponse, error) {
	f.receivedZipcode = zipcode
	return f.output, f.err
}

func TestHandlerReturnsWeatherWhenZipcodeIsValid(t *testing.T) {
	fakeUseCase := &fakeWeatherUseCase{
		output: &entity.WeatherResponse{
			TempC: 28.5,
			TempF: 83.3,
			TempK: 301.65,
		},
	}

	handler := NewHandler(fakeUseCase)

	req := httptest.NewRequest(http.MethodGet, "/weather/01001000", nil)
	res := httptest.NewRecorder()

	handler.GetWeatherByZipcode(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.Code)
	}

	if fakeUseCase.receivedZipcode != "01001000" {
		t.Fatalf("expected zipcode 01001000, got %s", fakeUseCase.receivedZipcode)
	}

	var response entity.WeatherResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("expected valid json, got error %v", err)
	}

	if response.TempC != 28.5 {
		t.Fatalf("expected temp_C 28.5, got %.2f", response.TempC)
	}

	if response.TempF != 83.3 {
		t.Fatalf("expected temp_F 83.3, got %.2f", response.TempF)
	}

	if response.TempK != 301.65 {
		t.Fatalf("expected temp_K 301.65, got %.2f", response.TempK)
	}
}

func TestHandlerReturns422WhenZipcodeIsInvalid(t *testing.T) {
	fakeUseCase := &fakeWeatherUseCase{
		err: usecase.ErrInvalidZipcode,
	}

	handler := NewHandler(fakeUseCase)

	req := httptest.NewRequest(http.MethodGet, "/weather/01001-000", nil)
	res := httptest.NewRecorder()

	handler.GetWeatherByZipcode(res, req)

	if res.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected status 422, got %d", res.Code)
	}

	expectedBody := "invalid zipcode"
	if res.Body.String() != expectedBody {
		t.Fatalf("expected body %q, got %q", expectedBody, res.Body.String())
	}
}

func TestHandlerReturns404WhenZipcodeIsNotFound(t *testing.T) {
	fakeUseCase := &fakeWeatherUseCase{
		err: usecase.ErrZipcodeNotFound,
	}

	handler := NewHandler(fakeUseCase)

	req := httptest.NewRequest(http.MethodGet, "/weather/99999999", nil)
	res := httptest.NewRecorder()

	handler.GetWeatherByZipcode(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", res.Code)
	}

	expectedBody := "can not find zipcode"
	if res.Body.String() != expectedBody {
		t.Fatalf("expected body %q, got %q", expectedBody, res.Body.String())
	}
}

func TestHandlerReturns500WhenUseCaseReturnsUnexpectedError(t *testing.T) {
	fakeUseCase := &fakeWeatherUseCase{
		err: errors.New("unexpected error"),
	}

	handler := NewHandler(fakeUseCase)

	req := httptest.NewRequest(http.MethodGet, "/weather/01001000", nil)
	res := httptest.NewRecorder()

	handler.GetWeatherByZipcode(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", res.Code)
	}

	expectedBody := "internal server error"
	if res.Body.String() != expectedBody {
		t.Fatalf("expected body %q, got %q", expectedBody, res.Body.String())
	}
}

func TestHandlerReturns405WhenMethodIsNotGet(t *testing.T) {
	fakeUseCase := &fakeWeatherUseCase{}

	handler := NewHandler(fakeUseCase)

	req := httptest.NewRequest(http.MethodPost, "/weather/01001000", nil)
	res := httptest.NewRecorder()

	handler.GetWeatherByZipcode(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", res.Code)
	}
}
