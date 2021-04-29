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

	// Declaring slice
	var a []int
	fmt.Println("a slice:", a)
	if a == nil {
		fmt.Println("a slice is nil....")
	}

	b := []int{}
	fmt.Println("b slice:", b)
	if b == nil {
		fmt.Println("b slice is nil....")
	}

	c := make([]int, 4)
	fmt.Println("c slice:", c)
	if c == nil {
		fmt.Println("c is nil....")
	}

	d := make([]int, 0, 10)
	d = append(d, 0)
	fmt.Println("d slice:", d, "length of d slice:", len(d), "capacity of d slice:", cap(d))

	d = append(d, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("d slice:", d, "length of d slice:", len(d), "capacity of d slice:", cap(d))

	d = append(d, 11)
	fmt.Println("d slice:", d, "length of d slice:", len(d), "capacity of d slice:", cap(d))
}
