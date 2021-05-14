package main

import "fmt"

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 1; i <= 10; i++ {
		fmt.Println(i, ":", twoBase(i), threeBase(i))
	}
}
