package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
		fmt.Println(v)
	}
	fmt.Println("evenVals:", evenVals)

	evenNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range evenNums {
		if v%2 == 0 {
			fmt.Println(v)
		} else {
			continue
		}
	}
}
