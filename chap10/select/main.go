package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Println(v)
			case v := <-ch2:
				fmt.Println(v)
			case v := <-ch3:
				fmt.Println(v)
			case <-done:
				fmt.Println("Done!")
				return
			}
			fmt.Println("For")
		}
	}()
	ch1 <- 2
	ch2 <- 4
	ch3 <- 8
	done <- true
	fmt.Println("Exited.")
}
