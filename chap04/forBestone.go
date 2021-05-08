package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, even := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-1 {
			break
		}
		fmt.Println("index:", i, "number:", even)
	}

	fmt.Println("***********************")

	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println("index:", i, "evenNum:", evenVals[i])
	}
}
