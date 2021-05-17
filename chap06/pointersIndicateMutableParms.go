// package main

// import "fmt"

// func main() {
// 	var f *int
// 	failedUpdate(f)
// 	fmt.Printf("%T\t%v\n", f, f)
// 	var i *int
// 	a := 100
// 	i = &a
// 	successUpdate(i)
// 	fmt.Println(*i)
// }

// func failedUpdate(f *int) {
// 	x := 10
// 	f = &x
// }

// func successUpdate(i *int) {
// 	*i = 200
// }

package main

import "fmt"

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func successUpdate(px *int) {
	*px = 20
}

func main() {
	x := 10
	failedUpdate(&x)
	fmt.Println("failedUpdate:", x)
	successUpdate(&x)
	fmt.Println("successUpdate:", x)
}
