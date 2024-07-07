package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l: l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello Requests")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops! Bad Request", http.StatusBadRequest)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "hello %s\n", d)

}
