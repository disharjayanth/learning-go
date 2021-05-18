package main

import "fmt"

func modifyMap(m map[int]string) {
	m[1] = "A"
}

func modifyArray(arr [10]int) {
	arr[0] = 100
	fmt.Println("inside modifyArray func:", arr)
}

func modifySlice(slice []int) {
	slice[0] = 100
}

func modifyInterface(something interface{}) {
	something = "Bye!"
}

func main() {
	a := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	modifyMap(a)
	fmt.Println(a)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	modifyArray(arr)
	fmt.Println(arr)

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	modifySlice(slice)
	fmt.Println(slice)

	something := "Hello world"
	modifyInterface(something)
	fmt.Println(something)
}
