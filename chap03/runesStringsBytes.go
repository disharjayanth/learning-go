package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello world!"
	fmt.Printf("%T\n", str)
	fmt.Println("Uppercase:", strings.ToUpper(str))
}
