package main

import (
	"errors"
	"fmt"
)

func main() {
	dividend, remainder, err := dividendAndRemainder(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Dividend:", dividend, "Remainder:", remainder)
	}

	dividend, remainder, err = dividendAndRemainder(0, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Dividend:", dividend, "Remainder:", remainder)
	}
}

func dividendAndRemainder(numerator int, denominator int) (dividend int, remainder int, err error) {
	if numerator == 0 {
		err = errors.New("Cannot divide by 0.")
		return
	}
	dividend = numerator / denominator
	remainder = numerator % denominator
	return
}
