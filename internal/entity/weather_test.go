package entity

import (
	"math"
	"testing"
)

func assertFloatEquals(t *testing.T, expected float64, actual float64) {
	t.Helper()

	const tolerance = 0.001

	if math.Abs(expected-actual) > tolerance {
		t.Fatalf("expected %.2f, got %.2f", expected, actual)
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	result := CelsiusToFahrenheit(28.5)
	expected := 83.3

	assertFloatEquals(t, expected, result)
}

func TestCelsiusToKelvin(t *testing.T) {
	result := CelsiusToKelvin(28.5)
	expected := 301.65

	assertFloatEquals(t, expected, result)
}

func TestNewWeatherResponse(t *testing.T) {
	response := NewWeatherResponse(28.5)

	assertFloatEquals(t, 28.5, response.TempC)
	assertFloatEquals(t, 83.3, response.TempF)
	assertFloatEquals(t, 301.65, response.TempK)
}

func TestNewWeatherResponseRoundsValues(t *testing.T) {
	response := NewWeatherResponse(25.555)

	assertFloatEquals(t, 25.56, response.TempC)
	assertFloatEquals(t, 78.0, response.TempF)
	assertFloatEquals(t, 298.71, response.TempK)
}
