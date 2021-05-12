package main

import "fmt"

func main() {
	someFunc := func() {
		fmt.Println("This is an anonymous func which is assigned to a variable called someFunc.")
	}

	someFunc()
}
