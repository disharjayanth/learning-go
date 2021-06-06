package main

import "fmt"

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= max; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("goroutine exited....")
	}()
	return ch
}

func main() {
	for val := range countTo(10) {
		fmt.Println(val)
	}
}
