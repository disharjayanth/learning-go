package main

import "fmt"

func main() {
	var f *int
	failedUpdate(f)
	fmt.Printf("%T\t%v\n", f, f)
	var i *int
	a := 100
	i = &a
	successUpdate(i)
	fmt.Println(*i)
}

func failedUpdate(f *int) {
	x := 10
	f = &x
}

func successUpdate(i *int) {
	*i = 200
}
