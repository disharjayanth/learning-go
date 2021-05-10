package main

import "fmt"

func main() {
	fmt.Println(addToBase(1))
	fmt.Println(addToBase(2, 1, 2, 3, 4))
	a := []int{4, 3}
	fmt.Println(addToBase(2, a...))
	fmt.Println(addToBase(3, []int{1, 2, 3, 4}...))
}

func addToBase(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base*v)
	}
	return out
}
