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
}
