package main

import "net/http"

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

func main() {

}
