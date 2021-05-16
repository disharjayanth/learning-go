package main

import "fmt"

type person struct {
	firstName string
}

func main() {
	a := 10
	b := true
	fmt.Println("a:", a, "* pointer of a:", &a)
	fmt.Println("b:", b, "* pointer of b:", &b)

	x := "hello"
	pointerToX := &x
	fmt.Println("x:", x, "pointerToX:", *pointerToX)

	y := 10
	pointerToY := &y
	fmt.Println("y:", y, "pointerToY:", pointerToY, "*pointerToY:", *pointerToY)
	z := *pointerToY + 2
	fmt.Println("z:", z)

	var l *int
	fmt.Println("l:", l)
	fmt.Println("l == nil:", l == nil)
	// nil pointer deference
	// fmt.Println("l value:", *l)

	var m = new(int)
	fmt.Println("m:", m, "value of m:", *m)
	var someMap = new(map[int]string)
	fmt.Println("someMap:", someMap, "*someMap:", *someMap)
	someFloat := new(float64)
	fmt.Println("someFloat:", someFloat, "*someFloat:", *someFloat)
	somePerson := new(person)
	fmt.Println("somePerson:", somePerson, "*somePerson:", *somePerson)
	someString := new(string)
	fmt.Println("someString:", someString, "*someString:", *someString)

	type foo struct {
		firstName string
		lastName  string
	}

	r := &foo{}
	var s string
	t := &s

	fmt.Printf("Type of foo struct: %T\nType of s string: %T\n", r, t)

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := &person{
		FirstName:  "Pat",
		MiddleName: stringPointer("Perry"),
		LastName:   "Peterson",
	}

	fmt.Printf("The type of p person: %T and with value %v\n", p, p)

	var someStringPointer *string
	someString1 := "Hello world!"
	someStringPointer = &someString1
	fmt.Printf("Type of someStringPointer: %T and with value %v\n", someStringPointer, someString1)
}

func stringPointer(s string) *string {
	return &s
}
