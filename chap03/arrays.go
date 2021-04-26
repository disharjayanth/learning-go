package main

import "fmt"

func main() {
	var x = [4]int{10, 20, 30, 40}
	fmt.Println(x)
	var y = [100]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)
	var z = [...]int{10, 20, 30, 40}
	fmt.Println(z)

	a := [4]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	fmt.Println("a==b", a == b)

	c := [2][3]int{}
	fmt.Println(c)
	fmt.Println("x:", len(x))
	fmt.Println("y:", len(y))
	fmt.Println("z:", len(z))
	fmt.Println("a:", len(a))
	fmt.Println("b:", len(b))
}
