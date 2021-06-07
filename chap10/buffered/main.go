package main

import "fmt"

func process(n int) int {
	return n * 2
}

func processChannel() []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func(i int) {
			results <- process(i)
		}(i)
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func main() {
	numSlice := processChannel()
	fmt.Println(numSlice)
}
