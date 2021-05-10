package main

import "fmt"

type myFuncOps struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	MyFunc(myFuncOps{
		FirstName: "Patel",
		Age:       50,
	})

	MyFunc(myFuncOps{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}

func MyFunc(opts myFuncOps) error {
	fmt.Println(opts.FirstName, opts.LastName, opts.Age)
	return nil
}
