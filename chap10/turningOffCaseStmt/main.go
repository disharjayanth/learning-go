package main

import "fmt"

func readInAndIn2(in <-chan int, in2 <-chan int, done <-chan struct{}) {
	for {
		select {
		case v, ok := <-in:
			if !ok {
				in = nil // this case will no longer run because read from nil channel never return a value.
				continue
			}
			// process v that was read from in
			fmt.Println(v)
		case v, ok := <-in2:
			if !ok {
				in2 = nil // this case will no longer run because read from nil channel never return a value.
				continue
			}
			// process v that was read from in2
			fmt.Println(v)
		case <-done:
			return
		}
	}
}

func main() {
	in := make(chan int)
	in2 := make(chan int)
	done := make(chan struct{})
	readInAndIn2(in, in2, done)
}
