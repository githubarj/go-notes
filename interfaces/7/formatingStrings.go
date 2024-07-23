package main

import (
	"errors"
	"fmt"
)

func getSmsErrorString (cost float64, receipient string) string {
	return fmt.Sprintf("SMS that cost %.2f to be sent to '%s' can no be sent", cost , receipient)

} 

type divideError struct {
	dividend float64
}

func (d divideError) Error () string {
	return fmt.Sprintf("can not divide %v by zero", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func divideNew(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by zero")
	}
	return x / y, nil
}

func validateStatus(status string) error {
	if len(status) == 0 {
		return errors.New("status cannot be empty")
	} 
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}
