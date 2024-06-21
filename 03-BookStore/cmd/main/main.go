package main

import (
	"log"
	"net/http"

	"github.com/porwalameet/go-projects/03-BookStore/pkg/routes"

	"github.com/gorilla/mux"
)

const (
	port = ":8088"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(port, r))
}
