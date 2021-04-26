package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(a, "length:", len(a), "capacity:", cap(a))
	var b []int
	b = []int{20, 40, 60, 80}
	fmt.Println(b, "length:", len(b), "capacity:", cap(b))
	var c = []int{2, 3, 4, 5}
	c = append(c, 1)
	fmt.Println(c, "length:", len(c), "capacity:", cap(c))
}
