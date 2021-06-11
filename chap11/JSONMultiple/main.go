package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

const data string = `{"name": "Fred", "age": 40}{"name": "Mary", "age": 21}{"name": "Pat", "age": 30}`

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var person Person
	dec := json.NewDecoder(strings.NewReader(data))

	var b bytes.Buffer
	enc := json.NewEncoder(&b)

	for dec.More() {
		err := dec.Decode(&person)
		if err != nil {
			fmt.Println("Error decoding:", err)
		} else {
			fmt.Println("Person:", person)
		}

		err = enc.Encode(person)
		if err != nil {
			fmt.Println("Error encoding:", err)
		}
	}

	fmt.Println(b.String())
}
