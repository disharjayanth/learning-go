package main

import "fmt"

var (
	x    int
	y        = 20
	z    int = 30
	d, e     = 40, "Hello world"
	f, g string
)

func main() {
	// var x int = 10
	// fmt.Printf("%T\n", x)
	// var y = 20.2
	// fmt.Printf("%T\n", y)
	// var a, b int = 10, 20
	// fmt.Printf("%T\t a=%v\t | %T\t b=%v\n", a, a, b, b)
	// var c, d = 10, "Hello world"
	// fmt.Printf("%T\t c=%v\t %T\t d=%v\n", c, c, d, d)
	fmt.Println(x, y, z, d, e, f, g)
	h, i := 30, "Hi there!"
	fmt.Println(h, i)

	i, j := "Bye there", true
	fmt.Println(i, j)
}
