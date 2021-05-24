package main

import (
	"fmt"
	"io"
	"os"
)

type MyInt int

func main() {
	var i interface{}
	var i2 MyInt = 10
	i = i2
	doThings(i)

	var i3 int = 10
	i = i3
	doThings(i)

	var i4 string = "Hello"
	i = i4
	doThings(i)

	i = nil
	doThings(i)

	var i5 bool = true
	i = i5
	doThings(i)

	var i6 rune = 100
	i = i6
	doThings(i)

	someStruct := struct {
		FirstName string
		LastName  string
	}{
		FirstName: "John",
		LastName:  "D",
	}
	i = someStruct
	doThings(i)

	file, err := os.Open("interfaceWithSwitchType.go")
	defer file.Close()
	if err != nil {
		fmt.Println("Error:", err)
	}
	doThings(file)

	var i7 os.File
	i = i7
	doThings(i)
}

func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		// i is nill, type of j is interface{}
		fmt.Println("Type of j is nil:", j)
		fmt.Printf("%T\n", j)
	case int:
		// j is type int
		fmt.Printf("%T\n", j)
	case MyInt:
		// j is type MyInt
		fmt.Printf("%T\n", j)
	case io.Reader:
		// j is type of io.Reader
		fmt.Printf("io.Reader type: %T\n", j)
	case string:
		// j is type of string
		fmt.Printf("%T\n", j)
	case bool, rune:
		// i is either bool or rune, so j is of type interface{}
		fmt.Printf("%T\n", j)
	default:
		// no idea what i is, j is of type interface{}
		fmt.Printf("Don't know the type: %T\n", j)
	}
}
