package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func loginCheck(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginInfo := r.FormValue("login")
		fmt.Println(loginInfo)
		if loginInfo == "1" {
			f(w, r)
		} else {
			w.Write([]byte("Invalid user!"))
		}
	}
}

func secretInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Secret Page"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/login", loginCheck(secretInfo))
	http.ListenAndServe(":3000", nil)
}
