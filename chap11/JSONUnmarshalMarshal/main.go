package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const data string = `{
	"id": "12345" ,
	"date_ordered": "2020-05-01T13:01:02Z" ,
	"customer_id": "3" ,
	"items": [{"id": "xyz123", "name": "Thing 1"}, {"id": "abc789", "name": "Thing 2"}]
}`

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	Customer    string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

func main() {
	var o Order
	err := json.Unmarshal([]byte(data), &o)
	checkError(err)
	fmt.Println("Order type:", o)

	sb, err := json.Marshal(o)
	if err != nil {
		checkError(err)
	}
	fmt.Println("JSON:", string(sb))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
