package main

import (
	"fmt"
	"math/rand"
)

func searcher(str string, index int) []string {
	fmt.Println("Running searcher no.:", index)
	n := rand.Intn(10)
	someChar := string(str[n])
	fmt.Println("Ending searcher:", index)
	return []string{someChar}
}

func searchData(s string, searchers []func(string, int) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for index, searcher := range searchers {
		go func(searcher func(string, int) []string, index int) {
			fmt.Println("********Running goroutine:", index)
			select {
			case result <- searcher(s, index):
			case <-done:
				return
			}
		}(searcher, index)
	}
	r := <-result
	close(done)
	return r
}

func main() {
	fmt.Println(searchData("Hello World", []func(string, int) []string{searcher, searcher, searcher, searcher}))
}
