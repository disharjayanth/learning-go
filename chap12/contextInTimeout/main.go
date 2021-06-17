package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var duration time.Duration
	fmt.Println("Enter the second:")
	fmt.Scanf("%v", &duration)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, duration*time.Second)
	defer cancel()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Duration exceeded 2 second.")
	case <-ctx.Done():
		fmt.Println("Contexted cancelled!")
	}
}
