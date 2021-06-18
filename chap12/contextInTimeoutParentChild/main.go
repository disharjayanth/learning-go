package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// Smaller context timer wins!
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	child, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()
	start := time.Now()
	select {
	case <-parent.Done():
		fmt.Println("Parent done!")
	case <-child.Done():
		fmt.Println("Child done!")
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
