package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

const (
	port = ":8088"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to WebServer : %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Thanks for accessing /hello URI")
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
