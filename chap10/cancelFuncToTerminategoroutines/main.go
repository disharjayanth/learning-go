package main

import "fmt"

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancel
}

func main() {
	var max int
	fmt.Println("Enter max number to count")
	fmt.Scanf("%d", &max)
	numChan, cancel := countTo(max)
	for num := range numChan {
		if num > 5 {
			break
		}
		fmt.Println(num)
	}
	cancel()
	fmt.Println("Exited.")
}
