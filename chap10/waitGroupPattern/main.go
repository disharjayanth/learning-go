package main

import (
	"fmt"
	"sync"
)

func processor(n int) int {
	return n * 2
}

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}

	return result
}

func main() {
	const conc int = 10
	in := make(chan int, conc)
	go func() {
		for i := 0; i < conc; i++ {
			in <- i
		}
		close(in)
	}()
	fmt.Println(processAndGather(in, processor, conc))
}
