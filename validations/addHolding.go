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