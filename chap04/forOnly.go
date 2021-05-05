package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i++
		time.Sleep(500 * time.Millisecond)
	}
}
