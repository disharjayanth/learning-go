package main

import "os"

func doPanic(msg string) {
	panic(msg)
}

func main() {
	// doPanic("Some panic")
	doPanic(os.Args[0])
}
