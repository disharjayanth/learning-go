package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("Inside if statement before declaration of x:=5:", x)
		x := 5
		fmt.Println("x value after declaration x:=5:", x)
	}
	fmt.Println("x value outside if statment:", x)

	fmt.Println("----------------------------------------------------------")

	a := 10
	if a > 5 {
		a, b := 5, 20
		fmt.Println(a, b)
	}
	fmt.Println(a)

	fmt.Println("*********************************************")

	s := 10
	fmt.Println(s)
	fmt = "oops"
	fmt.Println(fmt)
}
