package main

import (
	"fmt"
	"learningGo/chap07/dependencyInjection/di"
	"net/http"
)

func main() {
	l := di.LoggerAdapter(di.LogOutput)
	ds := di.NewSimpleDataStore()
	logic := di.NewSimpleLogic(l, ds)
	c := di.NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	fmt.Println("Listening at PORT: 8000")
	http.ListenAndServe(":8000", nil)
}
