package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func foo(words ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("this is the middleware func")

			for _, w := range words {
				fmt.Println(w)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is the handler func")
	fmt.Printf("w: %v\n", w)
	fmt.Printf("r: %v\n", r)
}

func main() {
	r := mux.NewRouter()

	m := alice.New(foo("hello", "world"))
	r.Handle("/", m.ThenFunc(handlerFunc))

	_ = http.ListenAndServe(":8080", r)
}
