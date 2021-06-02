package main

import "fmt"

type foo struct {
	x int
	S string
}

func (f foo) Hello() string {
	return "Hello"
}

func (f foo) goodBye() string {
	return "good bye!"
}

type Bar = foo

func main() {
	bar := Bar{
		x: 10,
		S: "Hi",
	}

	fmt.Println("bar.Hello():", bar.Hello())
	fmt.Println("bar.goodBye():", bar.goodBye())

	var foo foo = bar
	fmt.Println("foo.Hello():", foo.Hello())
	fmt.Println("foo.goodBye():", foo.goodBye())
}
