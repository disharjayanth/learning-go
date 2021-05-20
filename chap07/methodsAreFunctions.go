package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{
		start: 10,
	}
	fmt.Println(myAdder.AddTo(5))

	// method value
	f1 := myAdder.AddTo
	fmt.Println("Method values:", f1(10))

	// method expression
	f2 := Adder.AddTo
	fmt.Println("Method expression (created from type):", f2(myAdder, 30))
}
