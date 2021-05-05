package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz:", i)
			} else {
				fmt.Println("Fizz:", i)
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz:", i)
		} else {
			fmt.Println("Nothing:", i)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
