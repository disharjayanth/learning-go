package main

import (
	"errors"
	"fmt"
	"time"
)

func domeSomework() (int, error) {
	return 4, nil
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = domeSomework()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(1 * time.Second):
		return 0, errors.New("work timeout")
	}
}

func main() {
	fmt.Println(timeLimit())
}
