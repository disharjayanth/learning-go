package main

import (
	"errors"
	"fmt"
)

func main() {
	res, remainder, err := divAndRemainder(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Dividend:", res, "Remainder:", remainder)
}

func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if numerator == 0 {
		err = errors.New("Cannot divide by 0.")
		return
	}
	result = numerator / denominator
	remainder = numerator % denominator
	return
}
