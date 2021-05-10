package main

import "fmt"

func main() {
	res := div(2, 2)
	fmt.Println(res)
}

// div func returns division of two int numbers
func div(numerator, denominator int) int {
	if numerator == 0 {
		return 0
	}
	return numerator / denominator
}
