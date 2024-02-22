package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

/*
	LoadBalacer defines interface for the load balancer

*/

type LoadBalancer interface {
	ServeHttp(w http.ResponseWriter, r *http.Request)
	GetNextAvailableServer() *Server
}

type Server struct {
	URL         string
	Alive       bool
	Weight      int
	Connections int
	mutex       sync.Mutex
}

type ReverseProxy struct {
	backendURL string
	proxy      *httputil.ReverseProxy
}

func NewReverseProxy(backendURL string) *ReverseProxy {

	backend, _ := url.Parse(backendURL)

	return &ReverseProxy{
		backendURL: backendURL,
		proxy:      httputil.NewSingleHostReverseProxy(backend),
	}
}

func (rp *ReverseProxy) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Forwarding request to %s : %s\n", rp.backendURL, r.URL.Path)
	rp.proxy.ServeHTTP(w, r)
}
