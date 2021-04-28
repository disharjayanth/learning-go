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
	d := append(a, 100, 200, 300, 400)
	fmt.Println(d)
	e := append(a, b...)
	fmt.Println(e)
	a = append(a, 60)
	fmt.Println(a)
	fmt.Println("Length of a:", len(a))
	fmt.Println("Capacity of a:", cap(a))
}
