package main

import (
	"errors"
	"fmt"
)

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func someError() error {
	err := errors.New("someError func")
	fmt.Printf("%+v", err)
	return nil
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}

	err := someError()
	if err != nil {
		fmt.Println(err)
	}
}
