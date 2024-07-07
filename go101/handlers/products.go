package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/porwalameet/go-projects/go101/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		p.getProducts(rw, r)
// 		return
// 	case http.MethodPost:
// 		p.addProduct(rw, r)
// 		return
// 	case http.MethodPut:
// 		p.l.Println("Handle PUT method")
// 		re := regexp.MustCompile(`/([0-9]+)`)

// 		//p:= r.URI.Path
// 		group := re.FindAllStringSubmatch(r.URL.Path, -1)
// 		if len(group) != 1 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		if len(group[0]) != 2 {
// 			http.Error(rw, "Invalid Group", http.StatusBadRequest)
// 			return
// 		}
// 		idString := group[0][1]
// 		id, _ := strconv.Atoi(idString)
// 		p.updateProduct(id, rw, r)
// 		return
// 	default:
// 		rw.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }

func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	p.l.Println("Handle GET method")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST method")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT method")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(rw, "Failed to convert id to int", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "unable to unmarshal JSON", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)

	})

}
