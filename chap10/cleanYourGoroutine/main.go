package main

import "fmt"

func countTo(max int, done chan<- bool) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= max; i++ {
			ch <- i
		}
		close(ch)
		close(done)
	}()
	return ch
}

func main() {
	done := make(chan bool)
	for val := range countTo(10, done) {
		fmt.Println(val)
	}
	<-done
}
