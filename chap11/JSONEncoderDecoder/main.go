package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{
		Name: "Fred",
		Age:  40,
	}

	var person2 Person

	f, err := os.OpenFile("./temp.json", os.O_RDWR, os.ModeAppend)
	checkError(err)
	defer f.Close()
	err = json.NewEncoder(f).Encode(person)
	checkError(err)

	f, err = os.Open("temp.json")
	checkError(err)
	err = json.NewDecoder(f).Decode(&person2)
	checkError(err)
	fmt.Println("person2:", person2)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
