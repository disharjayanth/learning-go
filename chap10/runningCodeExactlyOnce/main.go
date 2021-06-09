package main

import (
	"fmt"
	"sync"
)

var parser SlowComplicatedParser
var once sync.Once

type SomeWork struct {
	info string
}

func (s *SomeWork) Parse(str string) string {
	s.info = str
	return s.info
}

type SlowComplicatedParser interface {
	Parse(string) string
}

func Parse(dataToParse string) string {
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
	return &SomeWork{
		info: "Hi",
	}
}

func main() {
	fmt.Println(Parse("Hello World!"))
}
