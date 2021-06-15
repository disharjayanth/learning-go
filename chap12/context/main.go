package main

import (
	"context"
	"fmt"
	"time"
)

func logic(ctx context.Context, info string) (string, error) {
	// do some interesting stuff here
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Second):
		fmt.Println("1 second op done!")
		return info, nil
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res, err := logic(ctx, "a string!")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("logic func returned string:", res)
}
