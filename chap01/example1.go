package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	wg.Add(11)
	for i := 0; i <= 10; i++ {
		go print(i)
	}
	wg.Wait()
	fmt.Println("End at:", time.Since(start))
}

func print(i int) {
	defer wg.Done()
	fmt.Println(i)
}
