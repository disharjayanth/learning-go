package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i <= 100; i++ {
		n := rand.Intn(10)
		if n == 0 {
			fmt.Println("random number is too low....", n)
		} else if n > 5 {
			fmt.Println("random number is too big....", n)
		} else {
			fmt.Println("random number is good....", n)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
