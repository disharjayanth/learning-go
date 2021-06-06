package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println("from go routine", v, v2)
		done <- true
	}()
	v := 2
	var v2 int
	for {
		select {
		case ch2 <- v:
			fmt.Println(v, "is sent to ch2")
		case v2 = <-ch1:
			fmt.Println("from main goroutine:", v, v2)
		case <-done:
			fmt.Println("returned from for loop")
			return
		}
	}
}
