package main

import "fmt"

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is name and age is %d.", p.Name, p.Age)
}

type Pet struct {
	Name string
}

func (p Pet) String() string {
	return fmt.Sprintf("%s is pet's name.", p.Name)
}

func callStringer(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	p1 := Person{
		Name: "John",
		Age:  28,
	}

	pet1 := Pet{
		Name: "Tom",
	}

	callStringer(p1)
	callStringer(pet1)
}
