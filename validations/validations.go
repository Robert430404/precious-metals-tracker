package validations

import (
	"strconv"
	"errors"
)

func ValidatePrice(value string) error {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("Invalid price provided, please enter just the price")
	}

	return nil
}

func ValidateTotal(value string) error {
	_, err := strconv.ParseInt(value, 16, 32)
	if err != nil {
		return errors.New("Invalid total provided, please enter just the total number of units")
	}

	return nil
}

func ValidateString(value string) error {
	if len(value) == 0 {
		return errors.New("Please enter a valid purchase source. (ex. JM Bullion)")
	}

	return nil
}