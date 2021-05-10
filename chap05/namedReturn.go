package main

import (
	"errors"
	"fmt"
)

func main() {
	res, remainder, err := divAndRemainder(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Dividend:", res, "Remainder:", remainder)
	}

	res, remainder, err = divAndRemainder(0, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Dividend:", res, "Remainder:", remainder)
	}
}

func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if numerator == 0 {
		err = errors.New("Cannot divide by 0.")
		return
	}
	return numerator / denominator, numerator % denominator, nil
}
