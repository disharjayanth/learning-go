package main

import "fmt"

type opFuncType func(int, int) int

func add(i int, j int) int {
	return i + j
}

func mul(i int, j int) int {
	return i * j
}

func div(i int, j int) int {
	return i / j
}

func main() {
	var opMap = map[string]opFuncType{
		"+": add,
		"*": mul,
		"/": div,
	}

	op, ok := opMap["+"]
	if !ok {
		fmt.Println("Operation not supported:", op)
	}
	result := op(2, 2)
	fmt.Println("Addition:", result)

	op, ok = opMap["*"]
	if !ok {
		fmt.Println("Operation not supported:", op)
	}
	result = op(2, 6)
	fmt.Println("Multiplication:", result)

	op, ok = opMap["/"]
	if !ok {
		fmt.Println("Operation not supported:", op)
	}
	result = op(2, 2)
	fmt.Println("Division:", result)

	op, ok = opMap["%"]
	if !ok {
		fmt.Println("Operation not supported:", op)
	} else {
		result = op(2, 2)
		fmt.Println("Modulus:", result)
	}
}
