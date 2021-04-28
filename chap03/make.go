package main

import "fmt"

func main() {
	x := make([]int, 10)
	fmt.Println("Length of x:", len(x), "Capacity of x:", cap(x))
	y := make([]int, 5, 10)
	fmt.Println("Length of y:", len(y), "Capacity of y:", cap(y))
	z := make([]int, 0, 10)
	fmt.Println("Length of z:", len(z), "Capacity of z:", cap(z))
	z = append(z, 10, 20, 30, 40)
	fmt.Println("z:", z)
}
