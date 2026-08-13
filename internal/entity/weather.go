package entity

import "math"

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewWeatherResponse(tempC float64) WeatherResponse {
	return WeatherResponse{
		TempC: roundToTwoDecimals(tempC),
		TempF: roundToTwoDecimals(CelsiusToFahrenheit(tempC)),
		TempK: roundToTwoDecimals(CelsiusToKelvin(tempC)),
	}
}

func CelsiusToFahrenheit(tempC float64) float64 {
	return tempC*1.8 + 32
}

func CelsiusToKelvin(tempC float64) float64 {
	return tempC + 273.15
}

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
