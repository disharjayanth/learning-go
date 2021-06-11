package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	defer ticker.Stop()
	go func() {
		time.Sleep(20 * time.Second)
		close(done)
	}()
	for {
		select {
		case <-done:
			return
		case tick := <-ticker.C:
			fmt.Println(tick)
		}
	}
}
