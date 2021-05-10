package main

import (
	"errors"
	"fmt"
)

func main() {
	dividend, remainder, error := divAndRemainder(4, 0)
	fmt.Println("Dividend:", dividend, "Remainder:", remainder, "Error:", error)
}

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by 0!")
	}
	return numerator / denominator, numerator % denominator, nil
}
