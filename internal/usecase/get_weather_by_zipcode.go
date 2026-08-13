package usecase

import (
	"context"
	"errors"

	"clima-por-cep/internal/entity"
)

var (
	ErrInvalidZipcode  = errors.New("invalid zipcode")
	ErrZipcodeNotFound = errors.New("can not find zipcode")
)

type Location struct {
	City  string
	State string
}

type ZipcodeGateway interface {
	GetLocationByZipcode(ctx context.Context, zipcode string) (*Location, error)
}

type WeatherGateway interface {
	GetCurrentTemperatureCelsius(ctx context.Context, city string) (float64, error)
}

type GetWeatherByZipcodeUseCase struct {
	zipcodeGateway ZipcodeGateway
	weatherGateway WeatherGateway
}

func NewGetWeatherByZipcodeUseCase(
	zipcodeGateway ZipcodeGateway,
	weatherGateway WeatherGateway,
) *GetWeatherByZipcodeUseCase {
	return &GetWeatherByZipcodeUseCase{
		zipcodeGateway: zipcodeGateway,
		weatherGateway: weatherGateway,
	}
}

func (u *GetWeatherByZipcodeUseCase) Execute(
	ctx context.Context,
	zipcode string,
) (*entity.WeatherResponse, error) {
	if !IsValidZipcode(zipcode) {
		return nil, ErrInvalidZipcode
	}

	location, err := u.zipcodeGateway.GetLocationByZipcode(ctx, zipcode)
	if err != nil {
		return nil, err
	}

	if location == nil || location.City == "" {
		return nil, ErrZipcodeNotFound
	}

	tempC, err := u.weatherGateway.GetCurrentTemperatureCelsius(ctx, location.City)
	if err != nil {
		return nil, err
	}

	response := entity.NewWeatherResponse(tempC)

	return &response, nil
}
