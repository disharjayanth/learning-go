package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, c := range sample {
			fmt.Println(i, c, string(c))
		}
	}
}
