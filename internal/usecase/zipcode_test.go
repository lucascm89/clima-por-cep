package usecase

import "testing"

func TestIsValidZipcodeReturnsTrueWhenZipcodeHasEightDigits(t *testing.T) {
	zipcode := "01001000"

	isValid := IsValidZipcode(zipcode)

	if !isValid {
		t.Fatalf("expected zipcode %s to be valid", zipcode)
	}
}

func TestIsValidZipcodeReturnsFalseWhenZipcodeHasLessThanEightDigits(t *testing.T) {
	zipcode := "1234567"

	isValid := IsValidZipcode(zipcode)

	if isValid {
		t.Fatalf("expected zipcode %s to be invalid", zipcode)
	}
}

func TestIsValidZipcodeReturnsFalseWhenZipcodeHasMoreThanEightDigits(t *testing.T) {
	zipcode := "123456789"

	isValid := IsValidZipcode(zipcode)

	if isValid {
		t.Fatalf("expected zipcode %s to be invalid", zipcode)
	}
}

func TestIsValidZipcodeReturnsFalseWhenZipcodeHasLetters(t *testing.T) {
	zipcode := "abc01000"

	isValid := IsValidZipcode(zipcode)

	if isValid {
		t.Fatalf("expected zipcode %s to be invalid", zipcode)
	}
}

func TestIsValidZipcodeReturnsFalseWhenZipcodeHasSpecialCharacters(t *testing.T) {
	zipcode := "01001-000"

	isValid := IsValidZipcode(zipcode)

	if isValid {
		t.Fatalf("expected zipcode %s to be invalid", zipcode)
	}
}

func TestIsValidZipcodeReturnsFalseWhenZipcodeIsEmpty(t *testing.T) {
	zipcode := ""

	isValid := IsValidZipcode(zipcode)

	if isValid {
		t.Fatalf("expected empty zipcode to be invalid")
	}
}
