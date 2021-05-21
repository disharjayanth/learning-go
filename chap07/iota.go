package main

import "fmt"

type counter int

const (
	a counter = iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(a, b, c, d, e)
}
