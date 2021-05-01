package main

import "fmt"

func main() {
	str := "Hello world!"
	var s1 string = str[4:7]
	var s2 string = str[:5]
	var s3 string = str[6:]
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)

	str1 := "Hello 🌞"
	s1 = str1[4:7]
	s2 = str1[:5]
	s3 = str1[6:]
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)

	var a rune = 'x'
	fmt.Printf("%v\t%T\n", a, a)

	// type conversion
	var b rune = 'x'
	var s string = string(b)
	fmt.Println(s)
	var c byte = 'd'
	var s4 string = string(c)
	fmt.Println(s4)

	// string to slice of byte / slice of rune, vice versa
	var str2 string = "Hello, 😊"
	fmt.Println(str2)
	var bs []byte = []byte(str2)
	var rs []rune = []rune(str2)
	fmt.Println("slice of byte:", bs)
	fmt.Println("slice of rune:", rs)

	// Own -> for range over string to print unicode point
	var someString string = "Hello world"
	fmt.Println("Unicode point for string:", someString)
	for _, s := range someString {
		fmt.Printf("%v\n", s)
	}
}
