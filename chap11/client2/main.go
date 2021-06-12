package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var posts []Post
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-My-Client", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(res.Body).Decode(&posts)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Posts:", posts)

	for i, v := range posts {
		if i == 5 {
			return
		}
		fmt.Printf("%d)%+v\n", i, v)
	}
}
