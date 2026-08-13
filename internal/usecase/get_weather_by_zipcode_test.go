package usecase

import (
	"context"
	"errors"
	"testing"
)

type fakeZipcodeGateway struct {
	location *Location
	err      error
}

func (f *fakeZipcodeGateway) GetLocationByZipcode(ctx context.Context, zipcode string) (*Location, error) {
	return f.location, f.err
}

type fakeWeatherGateway struct {
	tempC float64
	err   error
}

func (f *fakeWeatherGateway) GetCurrentTemperatureCelsius(ctx context.Context, city string) (float64, error) {
	return f.tempC, f.err
}

func TestGetWeatherByZipcodeUseCaseReturnsInvalidZipcodeWhenZipcodeIsInvalid(t *testing.T) {
	zipcodeGateway := &fakeZipcodeGateway{}
	weatherGateway := &fakeWeatherGateway{}

	useCase := NewGetWeatherByZipcodeUseCase(zipcodeGateway, weatherGateway)

	output, err := useCase.Execute(context.Background(), "01001-000")

	if output != nil {
		t.Fatalf("expected output to be nil")
	}

	if !errors.Is(err, ErrInvalidZipcode) {
		t.Fatalf("expected ErrInvalidZipcode, got %v", err)
	}
}

func TestGetWeatherByZipcodeUseCaseReturnsZipcodeNotFoundWhenLocationIsNil(t *testing.T) {
	zipcodeGateway := &fakeZipcodeGateway{
		location: nil,
	}
	weatherGateway := &fakeWeatherGateway{}

	useCase := NewGetWeatherByZipcodeUseCase(zipcodeGateway, weatherGateway)

	output, err := useCase.Execute(context.Background(), "01001000")

	if output != nil {
		t.Fatalf("expected output to be nil")
	}

	if !errors.Is(err, ErrZipcodeNotFound) {
		t.Fatalf("expected ErrZipcodeNotFound, got %v", err)
	}
}

func TestGetWeatherByZipcodeUseCaseReturnsZipcodeNotFoundWhenCityIsEmpty(t *testing.T) {
	zipcodeGateway := &fakeZipcodeGateway{
		location: &Location{
			City: "",
		},
	}
	weatherGateway := &fakeWeatherGateway{}

	useCase := NewGetWeatherByZipcodeUseCase(zipcodeGateway, weatherGateway)

	output, err := useCase.Execute(context.Background(), "01001000")

	if output != nil {
		t.Fatalf("expected output to be nil")
	}

	if !errors.Is(err, ErrZipcodeNotFound) {
		t.Fatalf("expected ErrZipcodeNotFound, got %v", err)
	}
}

func TestGetWeatherByZipcodeUseCaseReturnsWeatherWhenZipcodeIsValid(t *testing.T) {
	zipcodeGateway := &fakeZipcodeGateway{
		location: &Location{
			City:  "São Paulo",
			State: "SP",
		},
	}
	weatherGateway := &fakeWeatherGateway{
		tempC: 28.5,
	}

	useCase := NewGetWeatherByZipcodeUseCase(zipcodeGateway, weatherGateway)

	output, err := useCase.Execute(context.Background(), "01001000")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if output == nil {
		t.Fatalf("expected output not to be nil")
	}

	assertFloatEquals(t, 28.5, output.TempC)
	assertFloatEquals(t, 83.3, output.TempF)
	assertFloatEquals(t, 301.65, output.TempK)
}

func TestGetWeatherByZipcodeUseCaseReturnsErrorWhenZipcodeGatewayFails(t *testing.T) {
	expectedErr := errors.New("zipcode gateway error")

	zipcodeGateway := &fakeZipcodeGateway{
		err: expectedErr,
	}
	weatherGateway := &fakeWeatherGateway{}

	useCase := NewGetWeatherByZipcodeUseCase(zipcodeGateway, weatherGateway)

	output, err := useCase.Execute(context.Background(), "01001000")

	if output != nil {
		t.Fatalf("expected output to be nil")
	}

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected zipcode gateway error, got %v", err)
	}
}

func TestGetWeatherByZipcodeUseCaseReturnsErrorWhenWeatherGatewayFails(t *testing.T) {
	expectedErr := errors.New("weather gateway error")

	zipcodeGateway := &fakeZipcodeGateway{
		location: &Location{
			City:  "São Paulo",
			State: "SP",
		},
	}
	weatherGateway := &fakeWeatherGateway{
		err: expectedErr,
	}

	useCase := NewGetWeatherByZipcodeUseCase(zipcodeGateway, weatherGateway)

	output, err := useCase.Execute(context.Background(), "01001000")

	if output != nil {
		t.Fatalf("expected output to be nil")
	}

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected weather gateway error, got %v", err)
	}
}
