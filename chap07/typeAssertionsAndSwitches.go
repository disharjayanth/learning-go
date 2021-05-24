package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	fmt.Printf("%T\t%v\n", i, i)

	i2 := i.(MyInt)
	fmt.Printf("%v\t%T\n", i2, i2)

	// Code panics!
	// i3 := i.(string)
	// fmt.Printf("%v\t%T\n", i3, i3)

	// comma ok idiom
	i4, ok := i.(int)
	if !ok {
		fmt.Println("Error underlying type....")
		// return will exit main func
	}
	fmt.Println(i4)
}
