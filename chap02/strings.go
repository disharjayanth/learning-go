package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	str := `Hello world!`
	a := "a"
	fmt.Println(str)
	fmt.Println(a)
	fmt.Printf("%T\n%T\n", str, a)
	isTrue := str == a
	fmt.Printf("%t\t%T\n", isTrue, isTrue)

	p1 := person{
		name: "Jamie",
		age:  28,
	}
	fmt.Printf("%+v\n", p1)

	someString := "Hello " + "world!!"
	fmt.Println(someString)
}
