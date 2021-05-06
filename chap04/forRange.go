package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	for i, v := range a {
		fmt.Println("index:", i, "value:", v)
	}

	b := []string{"Hi", "Ola", "Hello", "Yo"}
	for _, v := range b {
		fmt.Println("value:", v)
	}

	c := []float64{45.5, 24.3, 12.2, 55.4}
	for i := range c {
		fmt.Println(i)
	}

	persons := map[string]bool{
		"John":  true,
		"Sa":    false,
		"Sammy": true,
		"Joe":   false,
	}

	for k, v := range persons {
		fmt.Println("Names in key:", k, "license:", v)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	for i := 0; i < 10; i++ {
		fmt.Println("---------", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
