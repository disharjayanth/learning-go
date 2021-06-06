// package main

// import "fmt"

// func main() {
// 	a := []int{2, 4, 6, 8, 10}
// 	ch := make(chan int, len(a))
// 	for _, v := range a {
// 		go func() {
// 			ch <- v * 2
// 		}()
// 	}

// 	for i := 0; i < len(a); i++ {
// 		fmt.Println(i, <-ch)
// 	}
// }

// two solutions:
// 1) using shadowing of variable inside for-loop
// package main

// import "fmt"

// func main() {
// 	a := []int{2, 4, 6, 8, 10}
// 	ch := make(chan int)
// 	for _, v := range a {
// 		v := v
// 		go func() {
// 			ch <- v * 2
// 		}()
// 	}

// 	for i := 0; i < len(a); i++ {
// 		fmt.Println(i, <-ch)
// 	}
// }

// 2) sending value inside goroutine as parameter
package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	for _, v := range a {
		go func(v int) {
			ch <- v * 2
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, <-ch)
	}
}
