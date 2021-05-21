package main

import "fmt"

type Score int
type HighScore Score

type Person struct {
	FirstName string
	LastName  string
	Age       int
	isLicense bool
}
type Employee Person

func main() {
	var i int = 200
	var s Score = 400
	var hs HighScore = 800
	hs = HighScore(s)
	s = Score(i)
	fmt.Printf("%v: %T\n%v: %T\n", s, s, hs, hs)

	p1 := Person{
		FirstName: "S",
		LastName:  "B",
		Age:       20,
		isLicense: true,
	}

	fmt.Println(p1)

	e1 := Employee{
		FirstName: "R",
		LastName:  "T",
		Age:       22,
		isLicense: false,
	}
	fmt.Println(e1)
}
