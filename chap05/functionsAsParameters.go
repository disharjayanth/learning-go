package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		Person{"Pat", "Patterson", 37},
		Person{"Tracy", "Bobbert", 23},
		Person{"Fred", "Fredson", 18},
	}
	fmt.Println("Unsorted By LastName:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println("Sorted By LastName:", people)

	fmt.Println("Accessing people after sorting:", people)

	fmt.Println("Sorting using age:")

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted By Age:", people)
}
