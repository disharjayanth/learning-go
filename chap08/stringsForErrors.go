package main

import (
	"errors"
	"fmt"
	"os"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed.")
	}
	return i * 2, nil
}

func doubleEvenWithFmtErrorf(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("Number %d cannot be processed since it is not even", i)
	}
	return i * 2, nil
}

func main() {
	var evenNum int
	fmt.Println("Enter even number to get doubled:")
	fmt.Scanf("%d", &evenNum)
	doubledEven, err := doubleEven(evenNum)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Doubled even number:", doubledEven)

	doubledEven, err = doubleEvenWithFmtErrorf(evenNum)
	if err != nil {
		fmt.Println("Error from fmt.Errorf:", err)
		os.Exit(1)
	}
	fmt.Println("Doubled even number from doubleEvenWithFmtErrorf:", doubledEven)
}
