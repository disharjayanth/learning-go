package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		func(j int) {
			fmt.Println("Anonymous function calling itself:", j)
		}(i)
	}
}
