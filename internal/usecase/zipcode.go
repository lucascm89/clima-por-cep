package usecase

import "unicode"

func IsValidZipcode(zipcode string) bool {
	if len(zipcode) != 8 {
		return false
	}

	for _, char := range zipcode {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}
