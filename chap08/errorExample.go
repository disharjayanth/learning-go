package main

import (
	"errors"
	"fmt"
	"os"
)

func calcRemainderAndMod(numerator, denominator int) (float64, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator cannot be zero")
	}
	return float64(numerator) / float64(denominator), numerator % denominator, nil
}

func main() {
	var numerator, denominator int
	var err error
	fmt.Println("Enter numerator:")
	fmt.Scanf("%d\n", &numerator)
	fmt.Println("Enter denominator:")
	fmt.Scanf("%d\n", &denominator)
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Remainder:", remainder, "Mod:", mod)
}
