package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		if n := rand.Intn(10); n == 0 {
			fmt.Println("random number is too small....")
		} else if n > 5 {
			fmt.Println("random number is big....")
		} else {
			fmt.Println("random number is perfect....")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
