// package calculator does simple calculations
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and output their addition
func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New("") takes a string arg and returns a value of type error thus allowing us to
		// make error value as per requirement , giving us full freedom .
		//You can use : fmt.Printf() or t.Errorf()
		return 0, errors.New("Divison by zero not allowed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative number is invalid input for sqrt function")
	}
	return math.Sqrt(a), nil
}
