package main

import (
	"errors"
	"fmt"
)

func main() {
	div, rem, err := divAndRemainder(2, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(div, rem)

	div, _, _ = divAndRemainder(4, 2)
	fmt.Println(div)
}

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if numerator == 0 {
		return 0, 0, errors.New("Cannot divide by 0.")
	}
	return numerator / denominator, numerator % denominator, nil
}
