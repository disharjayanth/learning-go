package main

import "fmt"

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrFoo = Sentinel("foo error")
	ErrBar = Sentinel("bar error")
)

func main() {
	fmt.Printf("ErrBar: %T\t error value: %v\n", ErrBar, ErrBar)
	fmt.Printf("ErrFoo: %T\t error value: %v\n", ErrFoo, ErrFoo)

	someError := error(ErrFoo)
	fmt.Printf("%T\t%v\n", someError, someError)
}
