package usecase

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
