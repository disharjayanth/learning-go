package main

import "fmt"

type nums struct {
	num1 int
	num2 int
}

func (n nums) Add(x int, y int) int {
	return x + y
}

func (n nums) Sub(x int, y int) int {
	return x - y
}

func (n nums) Div(x int, y int) int {
	return x / y
}

func (n nums) Multi(x int, y int) int {
	return x * y
}

type numOperations interface {
	Add(x int, y int) int
	Sub(x int, y int) int
	Div(x int, y int) int
	Multi(x int, y int) int
}

type MathExpress struct {
	operations numOperations
}

func main() {
	n := nums{
		num1: 2,
		num2: 2,
	}

	m := MathExpress{
		operations: n,
	}

	fmt.Println(m.operations.Add(2, 2))
	fmt.Println(m.operations.Sub(2, 2))
	fmt.Println(m.operations.Multi(2, 2))
	fmt.Println(m.operations.Div(2, 2))
}
