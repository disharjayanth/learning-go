package main

import "fmt"

type Foo struct {
	Field1 string
	Field2 int
}

// !!!! Don't do this!
func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// Do this
func MakeFoo1() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

func main() {
	f := &Foo{}
	err := MakeFoo(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MakeFoo func:", *f)

	foo1, err := MakeFoo1()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MakeFunc1 func:", foo1)
}
