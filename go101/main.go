package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handlers "github.com/porwalameet/go-projects/go101/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	// create the handlers
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	// crate the new ServeMux and register the handlers
	sm := http.NewServeMux()
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  5 * time.Second,
		Handler:      sm,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
	l.Println("Recived terminate, graceful shutdown", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

	//http.ListenAndServe(":9090", sm)
}
