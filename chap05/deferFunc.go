// package main

// import (
// 	"fmt"
// )

// func deferFunc() {
// 	fmt.Println("Running defer func....")
// }

// func helloWorld() string {
// 	defer deferFunc()
// 	defer func() {
// 		fmt.Println("Self invoking defer func.....")
// 	}()
// 	return "Hello, world!"
// }

// func main() {
// 	fmt.Println(helloWorld())
// }

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file selected....")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 2048)

	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
