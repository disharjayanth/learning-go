package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var duration time.Duration
	fmt.Println("2 second to run operation")
	fmt.Println("Enter second to timeout:")
	fmt.Scanf("%v", &duration)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("Duration exceeded 2 second....")
	case <-time.After(duration * time.Second):
		fmt.Println("Duration to complete:", duration*time.Second)
	}
}
