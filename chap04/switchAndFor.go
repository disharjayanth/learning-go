package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		switch n := rand.Intn(10); {
		case n == 0:
			fmt.Println("That's too low:", n)
		case n > 5:
			fmt.Println("That's too big:", n)
		default:
			fmt.Println("That's a perfect number:", n)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
