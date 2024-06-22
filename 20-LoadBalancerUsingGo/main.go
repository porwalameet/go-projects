package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const (
	port = "8088"
)

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

type backendServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

func (b *backendServer) Address() string {
	return b.addr
}

func (b *backendServer) IsAlive() bool {
	return true
}

func (b *backendServer) Serve(w http.ResponseWriter, r *http.Request) {
	b.proxy.ServeHTTP(w, r)
}

func newBackendServer(addr string) *backendServer {
	serverURL, err := url.Parse(addr)
	handleErr(err)

	return &backendServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	targetSever := lb.getNextAvailableServer()
	fmt.Printf("Forwarding request to address %q\n", targetSever.Address())
	targetSever.Serve(w, r)
}

func main() {
	servers := []Server{
		newBackendServer("https://www.google.com"),
		newBackendServer("http://example.com"),
		newBackendServer("https://www.github.com"),
	}

	lb := newLoadBalancer(port, servers)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.ServeProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving request at localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+lb.port, nil))
}
