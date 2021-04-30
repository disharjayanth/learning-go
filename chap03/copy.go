package main

import "fmt"

func main() {
	// #1 copy func
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	fmt.Println("Before copying from x:", y)
	num := copy(y, x)
	fmt.Println("num:", num, "\ny after copying from x:", y)

	// #2 copy from 4 element to 2 elem slice
	a := []int{1, 2, 3, 4}
	b := make([]int, 2)
	num = copy(b, a)
	fmt.Println("num:", num, "b:", b)

	// #3 copying from between elems in slice
	c := []int{10, 20, 30, 40, 50, 60, 70, 80}
	d := make([]int, 10)
	num = copy(d, c[1:6])
	fmt.Println("d:", d, "len:", len(d), "cap:", cap(d))

	// #4 overlapping slice
	e := []int{100, 200, 300, 400}
	num = copy(e[:3], e[1:])
	fmt.Println("e:", e, num)

	// #5
	f := []int{1, 2, 3, 4}
	g := [4]int{5, 6, 7, 8}
	h := make([]int, 2)
	copy(h, g[:])
	fmt.Println("h:", h)
	copy(g[:], f)
	fmt.Println("g:", g)
	fmt.Printf("f:%T\t g:%T\n", f, g)
}
