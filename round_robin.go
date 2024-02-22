package main

import (
	"net/http"
)

type RoundRobin struct {
	servers []*Server
	next    int
}

func NewRoundRobin(servers []*Server) *RoundRobin {
	return &RoundRobin{
		servers: servers,
		next:    0,
	}
}

func (lb *RoundRobin) GetNextAvailableServer() *Server {

	numServers := len(lb.servers)

	start := lb.next

	for i := 0; i < numServers; i++ {
		serverIndex := (start + i) % numServers
		server := lb.servers[serverIndex]

		server.mutex.Lock()
		alive := server.Alive
		server.mutex.Unlock()

		if alive {
			lb.next = (serverIndex + 1) % numServers
			return server
		}

	}

	return nil

}

func (lb *RoundRobin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server := lb.GetNextAvailableServer()
	if server != nil {
		proxy := NewReverseProxy(server.URL)
		logger.Print("Server is: ", server)

		proxy.ServeHttp(w, r)

	} else {
		logger.Print("Server is: ", server)
	}
}
