package main

import "fmt"

func main() {
	const a = "Hello world!"
	const b string = "Hi there!"
	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", b, b)
}
