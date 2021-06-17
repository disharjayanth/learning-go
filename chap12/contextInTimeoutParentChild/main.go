package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	child, cancel := context.WithTimeout(parent, 4*time.Second)
	defer cancel()
	start := time.Now()
	<-child.Done()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
