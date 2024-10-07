package validations

import (
	"strconv"
	"errors"
)

func ValidatePrice(value string) error {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("Invalid price provided, please enter just the number")
	}

	return nil
}

func ValidateString(value string) error {
	if len(value) == 0 {
		return errors.New("Please enter a valid purchase source. (ex. JM Bullion)")
	}

	return nil
}