// Package math contains math functions.
package math

import "errors"

// ErrDivideByZero is very much self explanatory.
var ErrDivideByZero = errors.New("cannot divide by zero")

// Add adds two floats.
func Add(a float64, b float64) float64 {
	return a + b
}

// Subtract subtracts two floats.
func Subtract(a float64, b float64) float64 {
	return a - b
}

// Multiply multiplies two floats.
func Multiply(a float64, b float64) float64 {
	return a * b
}

// Divide divides two floats.
func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return a / b, nil
}
