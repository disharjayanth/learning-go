package main

import "fmt"

func main() {
	str := someFunc()
	fmt.Println(*str)
}

func someFunc() *string {
	a := "Hello world"
	return &a
}
