package main

import (
	"fmt"
)

func process(val int) int {
	return val
}

func runThingConcurrently(in <-chan int, out chan<- int) {
	for val := range in {
		result := process(val)
		out <- result
	}
	close(out)
}

func sendToInChannel(in chan<- int) {
	for _, val := range []int{10, 20, 30, 40} {
		in <- val
	}
	close(in)
}

func readFromOutChannel(out <-chan int, done chan<- bool) {
	for val := range out {
		fmt.Println(val)
	}
	close(done)
}

func main() {
	done := make(chan bool)
	in := make(chan int)
	out := make(chan int)
	go sendToInChannel(in)
	go runThingConcurrently(in, out)
	go readFromOutChannel(out, done)
	fmt.Println("waiting")
	<-done
}
