package main

import "fmt"

type person struct {
	age  int
	name string
}

// primitive type and struct doesnt change value on passing to function
// it is also pass by value
func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	p := person{}
	i := 2
	s := "hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	sl := []int{1, 2}
	modSlice(sl)
	fmt.Println(sl)
}
